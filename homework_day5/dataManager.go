package homework_day5
// You can edit this code!
// Click here and start typing.

import (
"encoding/json"
"fmt"
"io/ioutil"
"os"
"sync"
)

// User represents an entity with ID and Name fields.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// DataManager encapsulates CRUD operations on the in-memory storage.
type DataManager struct {
	mu      sync.Mutex
	Storage map[int]User
}

// NewDataManager creates a new DataManager instance.
func NewDataManager() *DataManager {
	return &DataManager{
		Storage: make(map[int]User),
	}
}

// LoadDataFromFile loads initial data from a JSON file.
func (dm *DataManager) LoadDataFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return err
	}

	for _, user := range users {
		dm.Storage[user.ID] = user
	}

	return nil
}

// SaveDataToFile saves the current data to a JSON file.
func (dm *DataManager) SaveDataToFile(filename string) error {
	//TODO #1 implement
	f, err := os.Create(filename)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	encoder := json.NewEncoder(f)
	err = encoder.Encode(dm.Storage)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// Create adds a new User to the in-memory storage.
func (dm *DataManager) Create(name string) error {
	//TODO #2 implement: generate IDs in an incremental manner
	dm.mu.Lock()
	defer dm.mu.Unlock()
	maxIndex := 0
	for i := range dm.Storage {
		if i > maxIndex {
			maxIndex = i
		}
	}
	dm.Storage[maxIndex+1] = User{ID: maxIndex + 1, Name: name}
	return nil
}

// Read retrieves and prints the details of a User based on its ID.
func (dm *DataManager) Read(id int) (User, error) {
	//TODO #3 implement
	user, found := dm.Storage[id]
	if found {
		fmt.Println("User ", id, " found: ", user)
		return user, nil
	} else {
		return User{}, fmt.Errorf("UserID %v not exists", user.ID)
	}
}

// Update updates the details of an existing User based on its ID.
func (dm *DataManager) Update(user User) error {
	//TODO #4 implement
	dm.mu.Lock()
	defer dm.mu.Unlock()
	_, found := dm.Storage[user.ID]
	if found {
		dm.Storage[user.ID] = user
		fmt.Println("Success update user", user)
		return nil
	} else {
		return fmt.Errorf("UserID %v not exists", user.ID)
	}
}

// Delete removes a User from the in-memory storage based on its ID.
func (dm *DataManager) Delete(id int) error {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	//TODO #5 implement: delete and save it to data_deleted.json
	targetUser, found := dm.Storage[id]
	deletedDm := DataManager{Storage: map[int]User{id: targetUser}}
	if found {
		// save to data_deleted.json
		err := deletedDm.SaveDataToFile("data_deleted.json")
		if err != nil {
			panic("fail to save data_deleted.json")
		}

		// delete user
		delete(dm.Storage, id)

	} else {
		panic("error when trying delete user")
	}
	return nil
}

/*
	Requirement:
		* Implement TODO

		*  Ensure safe concurrent access to the in-memory data
			+ https://gobyexample.com/mutexes

		*  Handle the case when a user does not exist or duplicate

		*  Handle panic: https://gobyexample.com/recover

*/

func main() {
	// Example usage:
	dataManager := NewDataManager()

	// Load initial data from a file
	if err := dataManager.LoadDataFromFile("sample.json"); err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	defer func() {
		if err := dataManager.SaveDataToFile("data.json"); err != nil {
			panic("dataManager.SaveDataToFile: " + err.Error())
		}
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Done()
		user, err := dataManager.Read(1)
		if err != nil {
			panic("dataManager.Read: " + err.Error())
		}
		fmt.Println(user)
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		user := User{ID: 2, Name: "This is new name"}
		if err := dataManager.Update(user); err != nil {
			panic("dataManager.Update: " + err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		if err := dataManager.Delete(2); err != nil {
			panic("dataManager.Delete: " + err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		if err := dataManager.Create("Hello 3"); err != nil {
			panic("dataManager.Create" + err.Error())
		}
	}()

	wg.Wait()
	fmt.Println("------- DONE ------")
}


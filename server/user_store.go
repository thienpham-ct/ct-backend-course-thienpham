package server

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
)

// TODO #1: implement in-memory user store

func NewUserStore() *UserStore {
	return &UserStore{data: make(map[string]UserInfo)}
}

type UserStore struct {
	mu   sync.Mutex
	data map[string]UserInfo
}

var userStore = NewUserStore()

func generateMD5Hash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (u *UserStore) Save(info UserInfo) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	key := generateMD5Hash(info.Username)
	u.data[key] = info
	fmt.Println("Success save user info ", info)
	return nil
}

func (u *UserStore) Get(username string) (UserInfo, error) {
	u.mu.Lock()
	defer u.mu.Unlock()
	targetUser, found := u.data[generateMD5Hash(username)]
	if found {
		fmt.Println("GET Success, found user", username)
		return targetUser, nil
	} else {
		fmt.Println("GET Fail, cannot find user", username)
		return UserInfo{}, nil
	}
}

type UserInfo struct {
	ID       int
	Username string
	Password string
	FullName string
	Address  string
}

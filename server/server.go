package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/api/public/register", register)
	http.HandleFunc("/api/public/login", login)
	http.HandleFunc("/api/private/self", self)

	http.HandleFunc("/api/public/log/register", LogWrapper(register))
	http.HandleFunc("/api/public/log/login", LogWrapper(login))
	http.HandleFunc("/api/private/log/self", LogWrapper(self))

	http.ListenAndServe(":8090", nil)
}

/*
		TODO #2:
		- implement the logic to register a new user (username, password, full_name, address)
	  	- Validate username (not empty and unique)
	  	- Validate password (length should at least 8)
*/
func register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := userStore.Save(UserInfo{}); err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

type RegisterRequest struct {
}

/*
		TODO #3:
		- implement the logic to login
		- validate the user's credentials (username, password)
	  	- Return JWT token to client
*/
func login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	user, err := userStore.Get("")
	if err != nil {
		return
	}

	token, err := GenerateToken("", 24*time.Second)
	if err != nil {
		return
	}

	resp := LoginResponse{Token: token}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
	return
}

type LoginRequest struct {
}

type LoginResponse struct {
	Token string
}

/*
TODO #4:
- implement the logic to get user info
- Extract the JWT token from the header
- Validate Token
- Return user info`
*/
func self(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	extractUserNameFn := func(authenticationHeader string) (string, error) {
		panic("implement me")
	}

	username, err := extractUserNameFn(authHeader)
	if err != nil {
		return
	}

	user, _ := userStore.Get(username)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

/*
TODO: extra wrapper
Print some logs to console
  - Path
  - Http Status code
  - Time start, Duration
*/
func LogWrapper(handler http.HandlerFunc) http.HandlerFunc {
	panic("TODO implement me")
}

/*
	TODO #1: implement in-memory user store
	TODO #2: implement register handler
	TODO #3: implement login handler
	TODO #4: implement self handler

	Extra: implement log handler
*/

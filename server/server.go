package server

import (
	"encoding/json"
	"fmt"
	"github.com/felixge/httpsnoop"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

func Server() {
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

	//Read request to req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := UserInfo{
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Address:  req.Address,
	}
	_, found := userStore.data[generateMD5Hash(u.Username)]

	switch {
	case u.Username == "":
		fmt.Println("Username is empty")
		return
	case found:
		fmt.Println("Username is not unique")
		return
	case len(u.Password) < 8:
		fmt.Println("Password is at least 8 char")
		return
	}

	if err := userStore.Save(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

/*
		TODO #3:
		- implement the logic to login
		- validate the user's credentials (username, password)
	  	- Return JWT token to client
*/
func login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" || len(req.Password) < 8 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Username/Password is invalid")
		return
	}
	user, err := userStore.Get(req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		fmt.Println("Username is not exists")
		return
	}

	if user.Password != req.Password {
		http.Error(w, err.Error(), http.StatusForbidden)
		fmt.Println("Password is incorrect")
		return
	}

	token, err := GenerateToken(user.Username, 24*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	resp := LoginResponse{Token: token}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}

	return
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
		type MyCustomClaims struct {
			Username string `json:"sub"`
			jwt.RegisteredClaims
		}

		token, err := jwt.ParseWithClaims(authenticationHeader, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("ct-secret-key"), nil
		})

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			fmt.Printf("Token valid - user %v", claims.Username)
			return claims.Username, nil
		} else {
			fmt.Println(err)
			return "", err
		}

	}

	username, err := extractUserNameFn(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	user, _ := userStore.Get(username)
	fmt.Println("Get user info,", user)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}

/*
TODO: extra wrapper
Print some logs to console
  - Path
  - Http Status code
  - Time start, Duration
*/
func LogWrapper(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, r)
		log.Printf("%s %s | Status %d | Duration %s\n", r.Method, r.URL.Path, m.Code, m.Duration)

	}
}

/*
	TODO #1: implement in-memory user store
	TODO #2: implement register handler
	TODO #3: implement login handler
	TODO #4: implement self handler
	Extra: implement log handler
*/

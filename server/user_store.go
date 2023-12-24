package main

import "sync"

// TODO #1: implement in-memory user store

var userStore = NewUserStore()

func NewUserStore() *UserStore {
	return &UserStore{data: make(map[string]UserInfo)}
}

type UserStore struct {
	mu   sync.Mutex
	data map[string]UserInfo
}

func (u *UserStore) Save(info UserInfo) error {
	panic("TODO: implement me")
}

func (*UserStore) Get(username string) (UserInfo, error) {
	panic("TODO: implement me")
}

type UserInfo struct {
	// TODO: implement me
}

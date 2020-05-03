package model

import (
	"sync"
)

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Followers []string `json:"followers"`
}

var UsersMux = &sync.Mutex{}

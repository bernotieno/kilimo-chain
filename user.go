package main

import (
	"encoding/json"
	"os"
	"sync"
)

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
	Company   string `json:"company,omitempty"`
	Password  string `json:"password"`
}

var (
	usersMutex sync.RWMutex
	usersFile  = "users.json"
)

func (u *User) Save() error {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	// Open the file in append mode, or create it if it doesn't exist
	file, err := os.OpenFile(usersFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the user as JSON
	encoder := json.NewEncoder(file)
	err = encoder.Encode(u)
	if err != nil {
		return err
	}

	return nil
}

func LoadUsers() ([]User, error) {
	usersMutex.RLock()
	defer usersMutex.RUnlock()

	file, err := os.OpenFile(usersFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []User
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var user User
		err := decoder.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

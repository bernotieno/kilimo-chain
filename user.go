// user.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Company   string `json:"company"`
	Password  string `json:"password"`
}

func (u *User) Save() error {
	users, err := LoadUsers()
	if err != nil {
		return err
	}
	users = append(users, *u)
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("users.json", data, 0644)
}

func LoadUsers() ([]User, error) {
	file, err := os.Open("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []User{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

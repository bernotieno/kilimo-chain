package asfuncss

import (
	"encoding/json"
	"net/http"
	"os"
)

func CheckUsernameExist(useremail string, w http.ResponseWriter) bool {
	// Open the JSON file containing the users' data
	file, err := os.Open("users.json")
	if err != nil {
		// If the file doesn't exist, return false (username doesn't exist)
		if os.IsNotExist(err) {
			return false
		}
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return false
	}
	defer file.Close()

	// Read the file content
	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return false
	}

	// Check if the username exists in the list of users
	for _, user := range users {
		if user.Email == useremail {
			return true
		}
	}

	return false
}

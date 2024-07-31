package asfuncss

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

type Error struct {
	Err string
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Open the JSON file containing user data
	file, err := os.Open("users.json")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read and unmarshal the JSON content
	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	useremail := r.Form.Get("loginemaail")
	// Check if the username exists and the password matches

	for _, user := range users {
		if user.Email == useremail {
			fmt.Println(user.Email)
			if CheckPassword(r, user.Password) {
				http.Redirect(w, r, "/about", http.StatusFound)
			} else {
				data := Error{Err: "Wrong Passwors or user is not registerd."}
				tmpl, _ := template.ParseFiles("signin.html")
				r.ParseForm()
				tmpl.Execute(w, data)
				return
			}
		}
	}
}

func CheckPassword(r *http.Request, storedPassword string) bool {
	// compare the hashed password
	providedPassword := r.Form.Get("loginpassword")
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}

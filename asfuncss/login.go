package asfuncss

import (
	"encoding/json"
	"fmt"
	"log"
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
				http.Redirect(w, r, "/Dashboard", http.StatusFound)
			} else {

				data := Error{"WRONG PASSWORD"}
				tmpl, _ := template.ParseFiles("Signin.html")
				err = tmpl.Execute(w, data)
				if err != nil {
					log.Fatalf("template parsing error: %v", err)
				}
				log.Printf("Executing template with data: %+v", data)
				return
			}
		}
	}
	data := Error{"Username not found"}
	tmpl, _ := template.ParseFiles("Signin.html")
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("template parsing error: %v", err)
	}
	log.Printf("Executing template with data: %+v", data)
}

func CheckPassword(r *http.Request, storedPassword string) bool {
	// compare the hashed password
	providedPassword := r.Form.Get("loginpassword")
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}

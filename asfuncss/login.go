package asfuncss

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

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
	useremail := r.Form.Get("loginemail")
	// Check if the username exists and the password matches
	response := Response{}

	for _, user := range users {
		if user.Email == useremail {
			if CheckPassword(r, user.Password) {
				response.Target = "html"
				response.HTML = "<div class='success'>YOU HAVE SUCCESSFULLY SIGNED UP</div>"
				fmt.Println("you are in")
			} else {
				response.Target = "#erro"
				response.HTML = "<div class='error'>WRONG PASSWORD OR USERNAME</div>"
				w.Write([]byte("WRONG PASSWORD OR USERNAME"))
				fmt.Println(user.Password)
				fmt.Println(r.Form.Get("loginpassword"))
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CheckPassword(r *http.Request, storedPassword string) bool {
	// compare the hashed password
	providedPassword := r.Form.Get("loginpassword")
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}

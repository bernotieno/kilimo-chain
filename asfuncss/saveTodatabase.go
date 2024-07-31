package asfuncss

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
)

type User struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Confirmpassword string `json:"confirmpassword"`
}

type PageData struct {
	ErrorMessage string
}

var Email string

func Reg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := User{
		Name:            r.FormValue("Name"),
		Email:           r.FormValue("Email"),
		Password:        r.FormValue("Password"),
		Confirmpassword: r.FormValue("Confirmpassword"),
	}

	if CheckUsernameExist(user.Email, w) {
		data := Error{"WRONG PASSWORD"}
		tmpl, _ := template.ParseFiles("Signup.html")
		err := tmpl.Execute(w, data)
		if err != nil {
			log.Fatalf("template parsing error: %v", err)
		}
		log.Printf("Executing template with data: %+v", data)
		return

	} else if user.Password != user.Confirmpassword {
		return
	} else {
		user.Password = Hashpassword(user.Password)
		user.Confirmpassword = user.Password
		Email = user.Email
		SaveDetails(user, w)
		http.Redirect(w, r, "/signin", http.StatusFound)
	}
}

// func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
// 	t, err := template.ParseFiles(tmpl)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	t.Execute(w, data)
// }

func AboutUs(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Dashboard.html"))
	tmpl.Execute(w, nil)
}

func SaveDetails(user User, w http.ResponseWriter) {
	databaseFile := "users.json"

	var users []User
	if _, err := os.Stat(databaseFile); err == nil {
		// If the file exists, read the existing users
		fileContent, err := os.ReadFile(databaseFile)
		if err != nil {
			http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
			return
		}
		json.Unmarshal(fileContent, &users)
	}

	// Append the new user to the users slice
	users = append(users, user)

	// Marshal the updated users slice to JSON
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	// Write the JSON data to the file
	err = os.WriteFile(databaseFile, data, 0o644)
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
}

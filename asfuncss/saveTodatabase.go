package asfuncss

import (
	"encoding/json"
	"net/http"
	"os"
)

type User struct {
	Firstname       string `json:"Firstname"`
	Secondname      string `json:"Secondname"`
	Email           string `json:"email"`
	Company         string `json:"company"`
	Password        string `json:"password"`
	Confirmpassword string `json:"confirmpassword"`
}

type Response struct {
    Target string `json:"target"`
    HTML   string `json:"html"`
}

func Reg(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    user := User{
        Firstname:       r.Form.Get("Firstname"),
        Secondname:      r.Form.Get("Secondname"),
        Email:           r.Form.Get("Email"),
        Company:         r.Form.Get("Company"),
        Password:        r.Form.Get("Password"),
        Confirmpassword: r.Form.Get("Confirmpassword"),
    }

    response := Response{}

    if CheckUsernameExist(user.Email, w) {
    
        response.Target = ".erro"
        response.HTML = "<div class='error'>Username already exists</div>"
    } else if user.Password != user.Confirmpassword {
    
        response.Target = "#erro"
        response.HTML = "<div class='error'>Password should be the same as the confirm password</div>"
    } else {
        user.Password = Hashpassword(user.Password)
        user.Confirmpassword = user.Password
        SaveDetails(user, w)
        response.Target = "html"
        response.HTML = "<div class='success'>YOU HAVE SUCCESSFULLY SIGNED UP</div>"
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
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

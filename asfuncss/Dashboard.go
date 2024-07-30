package asfuncss

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"
)

type Farmer struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	IdentityNo int    `json:"identity_no"`
	Contact    int    `json:"contact"`
	Email      string `json:"email"`
}

// type Transaction struct {
// 	TotalSales     int
// 	CompletedSales int
// 	InTransitSales int
// 	Profit         int
// }
// type TrendingItems struct {
// 	TrendingItems string
// }

// type PageDatas struct {
// 	farmerDetail    Farmer
// 	Spenditure      Transaction
// 	CurrentTrending TrendingItems
// }

// Function to handle the HTTP request and render the dashboard
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Open and read the JSON file
	file, err := os.Open("users.json")
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Decode JSON data into a slice of Farmer structs
	var farmers []Farmer
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&farmers); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	// Find the farmer with the specified email
	var foundFarmer *Farmer
	for _, farmer := range farmers {
		if farmer.Email == Email {
			foundFarmer = &farmer
			break
		}
	}

	// Check if the farmer was found
	if foundFarmer == nil {
		http.Error(w, "Farmer not found", http.StatusNotFound)
		return
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("Dashboard.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, foundFarmer)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

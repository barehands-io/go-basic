package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person structure
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func CheckUsers() {
	// Read the JSON file
	data, err := os.ReadFile("./uploads/users.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Log the JSON file content
	fmt.Println(string(data))

	// Define a variable to hold the parsed data as a slice of Person
	var peopleData []Person

	// Unmarshal the JSON data into the slice of Person
	err = json.Unmarshal(data, &peopleData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

}

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Person structure
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func CheckUsers() {
	// Open and read the JSON file
	file, err := os.Open("./uploads/users.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the file contents
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
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

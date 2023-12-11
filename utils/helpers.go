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
	fmt.Println("JSON File Content:")
	fmt.Println(string(data))

	// Define a variable to hold the parsed data
	var personData Person

	// Unmarshal the JSON data into the Go struct
	err = json.Unmarshal(data, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Access the data in the Go struct
	fmt.Println("Name:", personData.Name)
	fmt.Println("Age:", personData.Age)
	fmt.Println("Email:", personData.Email)
}

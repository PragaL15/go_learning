package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Corrected constant declaration
	const myURL = "http://localhost:8000/post"

	// Define a struct for the JSON data
	type User struct {
		Name   string   `json:"name"`
		Age    int      `json:"age"`
		Email  string   `json:"email"`
		Skills []string `json:"skills"`
	}

	// Perform HTTP GET request
	response, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Create an instance of the User struct
	var user User

	// Unmarshal the raw JSON data into the user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	// Access the unmarshaled data
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Email:", user.Email)
	fmt.Println("Skills:", user.Skills)
}

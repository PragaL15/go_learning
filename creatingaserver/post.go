package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	PerformCustomPostRequest()
}

func PerformCustomPostRequest() {
	// Define the URL where you want to make the POST request
	const myURL = "http://localhost:8000/post"

	// Create the JSON payload
	requestBody := `{
		"name": "Pragal",
		"email": "pragal@example.com"
	}`

	// Create a new HTTP request
	req, err := http.NewRequest("POST", myURL, bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		panic(err)
	}

	// Set headers for the request (such as Content-Type for JSON)
	req.Header.Set("Content-Type", "application/json")

	// Send the request using http.Client
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Read the response body
	content, _ := ioutil.ReadAll(response.Body)

	// Print the response status and content
	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", string(content))
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	PerformCustomPostRequest()
}

func PerformCustomPostRequest() {
	// Define the URL where you want to make the POST request
	const myURL = "http://localhost:8000/post"

	// Create form data
	formData := url.Values{}
	formData.Add("Name", "Pragal")
	formData.Add("Age", "19") // Age must be a string

	// Perform the POST request
	response, err := http.PostForm(myURL, formData)
	if err != nil {
		fmt.Println("Error in POST request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print response status and body
	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", string(body))
}

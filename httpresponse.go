package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://localhost:5173/"

func main() {
	
	res, err := http.Get(url)
	if err != nil {
		panic(err) }
	defer res.Body.Close() 

	databyte, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Converting the response bytes to a string
	content := string(databyte)
	fmt.Println(content) 
}

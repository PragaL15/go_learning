package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
Name string `json:"CourseName"`
Price int 
Platform string `json:"website"`
Password string `json:"-"`
Tags []string  `json:"tags,omitempty"`
}  

func main() {
	EncodeJson()
}
func EncodeJson(){
	courses := []course{
		{"ReactJS",300,"youtube.com","pragalya123",[]string{"web develop","DSA"}},
		{"NodeJS",400,"chrome.com","pragal@!#",[]string{"python","Java"}},
		{"NextJS",700,"brave.com","pragal123",nil},
	}
	finalJson , _ := json.MarshalIndent(courses,"","\t")
	fmt.Printf("%s\n", finalJson)

} 
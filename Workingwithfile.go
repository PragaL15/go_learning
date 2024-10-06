package main

import (
	"fmt"
	"io"
	"os"
)
func main(){
	content := "Hey there, I'm Pragalya from department of computer science and design"
	file ,err := os.Create("./About.txt")
	if err!=nil{
		panic(err)
	}
	len , err := io.WriteString(file, content)
	if err!=nil{
		panic (err)
	}
	fmt.Println(len)
	defer file.Close()
}
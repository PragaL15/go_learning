package main
import "fmt"
func main(){
	lang := make(map[int]string)
	lang[1] = "react"
	lang[2] = "java"
	lang[3] = "javascript"
	lang[4] = "python"
	fmt.Println(lang)//the o/p is map[1:react 2:java 3:javascript 4:python]
	fmt.Println(lang[1])//the print the o/p of individual value
	delete(lang , 2)
	fmt.Println(lang)//to delete the value of that specific key
}
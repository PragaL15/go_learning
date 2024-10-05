package main

import "fmt"

func main() {
    var course = []string{"react", "java", "ruby"}
    var index int = 2
    course = append(course[:index], course[index+1:]...)
    fmt.Println(course)
}

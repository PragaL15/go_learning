package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main(){
	PerformGetRequest()
}
func PerformGetRequest(){
	const myURL = "http://localhost:8000/get"
	response , err := http.Get(myURL)
	if(err != nil){
		panic(err)
}
defer response.Body.Close()
fmt.Println("status code:", response.StatusCode)
fmt.Println("content length:", response.ContentLength)

content , _ := ioutil.ReadAll(response.Body)
fmt.Println(string(content))
}
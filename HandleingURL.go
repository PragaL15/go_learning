package main
import ("fmt"
"net/url")
const Myurl string="http://localhost:5173/add" 
func main() {
//parsing the url
result,_ :=  url.Parse(Myurl)
fmt.Println(result.Scheme)
fmt.Println(result.Host)
fmt.Println(result.Path)

}
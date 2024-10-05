package main
import "fmt"
func main(){

	type User struct {
		Name string
		Email string
		Age int
}
Pragalya := User{"Pragal","pragal@gmail.com",19}
fmt.Printf("the detail of user %+v\n",Pragalya)



}
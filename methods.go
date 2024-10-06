package main
import (
	"fmt"
)
func main(){
	Pragal := User {"Pragalya",20}
	Pragal.GetAge()
	Pragal.ChangeName()//Manipulated the name 
	fmt.Println(Pragal)//but the name isn't changed
	// at the original place but it just worked at 
	//the copy of the struct and original remains same
}
	type User struct {
		Name string
		Age int 
	}
	func (u User) GetAge(){
		fmt.Println(u.Age)
	}
 func (u User) ChangeName(){
	u.Name = "Nagul"
	fmt.Println(u.Name)
 }
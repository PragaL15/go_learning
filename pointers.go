package main
import "fmt"
func main() {
	//output would be nil
//  var ptr *int;
//  fmt.Println(ptr)


myNum := 30
var ptr = &myNum
fmt.Println(ptr)//This will print the address of the myNum 
fmt.Println(*ptr)
*ptr = *ptr + 1
fmt.Println(myNum)//op will be 31 , as the operation is performed in it's original value

var myNum2  = myNum
myNum2 = myNum2+2 //similar stuff is haappening but the vaue is not changed in myNum as the operation is performed on the copied value
fmt.Println(myNum)
fmt.Println(myNum2)
}
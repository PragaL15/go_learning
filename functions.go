package main
import "fmt"
func main(){

result := adder(2,3,4,5,1)
fmt.Println(result)
}
func adder(values ...int) int {
	total:=0
	for _, val := range values {
   total += val
	}
	return total
}
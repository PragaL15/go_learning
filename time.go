package main
import("fmt"
"time"
)
func main() {
	fmt.Println("time formatting")
	presentTime:=time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05")) 
}
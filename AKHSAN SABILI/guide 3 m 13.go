package main
import "fmt"
	func main(){
	var x, y int
	var kond bool
	fmt.Scan(&x, &y)
	for kond == false {
		x-= y
		fmt.Println(x)
		kond = x <= 0
	}
	fmt.Println(x == 0)
}
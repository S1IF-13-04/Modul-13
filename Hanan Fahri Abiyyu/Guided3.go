package main
import "fmt"
func main() {
	var a, b int
	var stop bool
	fmt.Scan(&a, &b)
	for stop = false; !stop;{
		a -= b
		fmt.Println(a)
		stop = a <= 0
	}
	fmt.Println(a == 0)
}

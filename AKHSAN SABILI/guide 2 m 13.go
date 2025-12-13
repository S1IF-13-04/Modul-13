package main
import "fmt"
func main(){
	var n int
	var x bool
	for x == false{
		fmt.Scan(&n)
		x = n > 0
	}
	fmt.Println(n, "adalah bilangan bulat positif")
}
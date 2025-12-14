package main
import "fmt"
func main(){
	var n int
	for k := false; !k;{
		fmt.Scan(&n)
		k = n > 0
	}
	fmt.Println(n,"adalah bilbulpos")
}
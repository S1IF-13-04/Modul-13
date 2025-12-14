package main
import "fmt"

func main (){
	var x int64
	var k bool
	for k = false; !k;{
		fmt.Scan(&x)
		k = x > 0
	}
	fmt.Printf("%d adalah bilangan bulat positif", x)
}
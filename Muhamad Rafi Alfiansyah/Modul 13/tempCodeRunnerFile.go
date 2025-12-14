package main
import "fmt"
func main(){
	var n int
	var k bool
	fmt.Scan(&n)
	h := 0
	for k=false; !k;{
		n = n / 10
		h++
		k = n == 0
	}
	fmt.Print(h)
}
package main
import "fmt"
func main(){
	var x,y int
	var k bool
	fmt.Scan(&x,&y)
	for k = false;!k;{
		x = x-y
		fmt.Println(x)
		k = x <= 0
	}
		fmt.Print(x==0)
}
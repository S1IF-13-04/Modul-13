package main
import "fmt"

func main (){
	var x,y int64
	var k,a bool
	fmt.Scan(&x,&y)
	for k = false; !k;{
		x -= y
		fmt.Println(x)
		k = x<=0
	}
	a = x ==0
	fmt.Println(a)
}
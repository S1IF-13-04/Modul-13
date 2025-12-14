package main
import "fmt"

func main (){
	var x int64
	var y string
	var k bool
	fmt.Scan(&y,&x)
	for k = false; !k;{
		fmt.Println(y)
		x--
		k = x==0
	}
}
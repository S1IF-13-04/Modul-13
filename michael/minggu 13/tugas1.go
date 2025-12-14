package main
import "fmt"

func main(){
	var n,y int64
	fmt.Scan(&n)
	for n > 0 {
		y+=1
		n/=10
	}
	fmt.Printf("%d",y)
}
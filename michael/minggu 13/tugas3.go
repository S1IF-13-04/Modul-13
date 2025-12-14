package main
import "fmt"

func main(){
	var n,y,d,o int64
	fmt.Scan(&n)
	for n > y {
		fmt.Scan(&d)
		o+=1
		y+=d
		fmt.Printf("donatur%d menyumbang %d total %d",o,d,y)
	}
	fmt.Printf("Total: %d dari %d donatur",y,o)
}
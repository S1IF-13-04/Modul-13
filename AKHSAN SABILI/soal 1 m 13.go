package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	hasil := 0
	for kons := false; !kons;{
		hasil++
		n /= 10
		kons = n <= 0
	}
	fmt.Println(hasil)
}
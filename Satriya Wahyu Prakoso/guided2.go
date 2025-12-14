package main
import "fmt"
func main() {
	var n int
	for kondisi := false; !kondisi; {
		fmt.Scan(&n)
		kondisi = n > 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", n)
}
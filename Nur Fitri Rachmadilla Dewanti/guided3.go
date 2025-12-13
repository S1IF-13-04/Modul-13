package main
import "fmt"
func main() {
	var x, y int
	fmt.Scan(&x, &y)
	hasil := x

	for kondisi := false; !kondisi; {
		hasil = hasil - y
		fmt.Println(hasil)
		kondisi = hasil <= 0
	}
	fmt.Println(hasil == 0)
}
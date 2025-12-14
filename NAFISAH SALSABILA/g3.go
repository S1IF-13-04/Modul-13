package main
import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	for kondisi := false; !kondisi; {
		x -= y
		kondisi = x == y || x < 0
		fmt.Println(x == 0)
	}
}

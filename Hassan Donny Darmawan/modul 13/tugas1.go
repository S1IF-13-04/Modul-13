package main
import "fmt"

func main() {
	var a, t int
	fmt.Scan(&a)

	for kondisi := false; !kondisi; {
		t++
		a /= 10
		kondisi = a<=0
	}
	fmt.Print(t)
}

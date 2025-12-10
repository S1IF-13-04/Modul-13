package main
import "fmt"
func main() {
	var x int
	var kondisi bool
	fmt.Scan(&x)
	counter := 0
	for kondisi = false; !kondisi;{
		x /= 10
		counter++
		
		kondisi = x == 0
	}
	fmt.Println(counter)
}

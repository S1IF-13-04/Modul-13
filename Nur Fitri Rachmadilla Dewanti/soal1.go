package main
import "fmt"
func main() {
	var bilangan, digit int
	fmt.Scan(&bilangan)

	for kondisi := false; !kondisi; {
		bilangan = bilangan/10
		digit++
		if bilangan == 0{
		kondisi = true
		}
	}
	fmt.Println(digit)
}
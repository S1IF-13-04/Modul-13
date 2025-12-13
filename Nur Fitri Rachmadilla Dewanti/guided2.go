package main
import "fmt"
func main() {
	var bilangan int

	for kondisi := false; !kondisi; {
		fmt.Scan(&bilangan)
		kondisi = bilangan > 0
	}
	fmt.Println(bilangan, "adalah bilangan bulat positif")
}
package main 

import "fmt"

func main() {
	var x int
	var kondisi bool

	for kondisi = false; !kondisi; {
		fmt.Scan(&x)
		kondisi = x > 0
	}
	fmt.Print(x, " adalah bilangan bulat positif")
}
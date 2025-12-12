package main

import "fmt"

func main() {
	var x int
	for kondisi := false; !kondisi; {
		fmt.Scan(&x)
		kondisi = (x > 0)
	}
	fmt.Println(x, "adalah bilangan bulat positif")
}

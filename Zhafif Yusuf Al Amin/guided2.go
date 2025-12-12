package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for kondisi := true; kondisi; {
		fmt.Scan(&n)
		kondisi = n <= 0
	}
	fmt.Print(n, " adalah bilangan bulat positif")

}

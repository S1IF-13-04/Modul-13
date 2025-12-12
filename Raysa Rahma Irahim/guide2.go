package main

import "fmt"

func main() {
	var n int

	for kondisi := false; !kondisi; {
		fmt.Scan(&n)
		kondisi = (n > 0)
	}
	fmt.Println(n, "adalah bilangan bulat positif")
}

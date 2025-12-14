package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	for kondisi := false; !kondisi; {
		x -= y
		kondisi = x == 0 || x < 0
		fmt.Println(x)
	}
	fmt.Println(x == 0)
}

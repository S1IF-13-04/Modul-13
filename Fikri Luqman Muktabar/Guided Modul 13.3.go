package main

import "fmt"

func main() {
	var x, y int
	var kondisi bool
	fmt.Scan(&x, &y)

	for kondisi = false; !kondisi; {
		x = x - y
		fmt.Println(x)
		kondisi = x <= 0
	}
	
	fmt.Println(x == 0)
}
package main

import "fmt"

func main() {	
	var x, y int
	var kondisi bool

	fmt.Scan(&x, &y)

	for kondisi = true; kondisi; {
		pengurangan := x - y
		fmt.Println(pengurangan)
		x = pengurangan
		kondisi = x > 0
	}

	fmt.Print(x == 0)
}
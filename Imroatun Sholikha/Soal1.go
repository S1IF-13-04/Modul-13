package main

import "fmt"

func main() {
	var n, digit int
	var kondisi bool

	fmt.Scan(&n)

	for kondisi = false; !kondisi; {
		digit++
		n = n / 10
		kondisi = n == 0
	}
	fmt.Println(digit)
}

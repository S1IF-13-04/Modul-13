package main

import "fmt"

func main() {
	var n, digit int
	var selesai bool

	fmt.Scan(&n)
	digit = 0

	for selesai = false; !selesai; {
		n = n / 10
		digit++
		selesai = n == 0
	}

	fmt.Println(digit)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	digit := 0
	selesai := false
	for !selesai {
		n = n / 10
		digit++
		selesai = (n == 0)
	}
	fmt.Println(digit)
}

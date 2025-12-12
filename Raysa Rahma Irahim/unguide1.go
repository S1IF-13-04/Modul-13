package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	var selesai bool

	for selesai = false; !selesai; {
		n = n / 10
		count++
		selesai = (n == 0)
	}
	fmt.Println(count)
}

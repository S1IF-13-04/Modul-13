package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	i := 0
	for kondisi := false; !kondisi; {
		n = n / 10
		i++
		kondisi = n <= 0
	}
	fmt.Print(i)
}

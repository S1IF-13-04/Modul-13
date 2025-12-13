package main

import "fmt"

func main() {
	var x int
	var counter int

	fmt.Scan(&x)
	counter = 0

	for kondisi := false; !kondisi; {
		if x % 10 >= 1 {
			counter++
			x = x / 10
		}
		kondisi = x == 0
	}
	fmt.Print(counter)
}
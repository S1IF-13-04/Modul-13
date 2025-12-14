package main

import "fmt"

func main() {
	var n int
	var c bool
	for c = true; c; {
		fmt.Scan(&n)
		c = n <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", n)
}

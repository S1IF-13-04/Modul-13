package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	target := int(x) + 1

	for x < float64(target) {
		x = x + 0.1
		fmt.Printf("%.1f\n", x)
	}
}

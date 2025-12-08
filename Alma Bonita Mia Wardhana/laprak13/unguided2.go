package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	upper := float64(int(x)) + 1

	for v := x + 0.1; v <= upper; v += 0.1 {
		fmt.Printf("%.1f\n", v)
	}
}

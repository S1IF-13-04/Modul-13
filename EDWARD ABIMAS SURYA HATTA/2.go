package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)
	target := math.Ceil(x)
	for ok := true; ok; {
		x += 0.1
		fmt.Printf("%.1f\n", x)
		ok = x < target-0.00001
	}
}
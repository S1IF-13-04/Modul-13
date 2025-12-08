package main

import (
	"fmt"
	"math"
)

func main() {
	var dec float64
	fmt.Scan(&dec)
	num := dec
	fmt.Printf("Ceiling = %.2f\n", math.Ceil(num))

	for i := int(dec * 10); i <= int(math.Ceil(num)*10); i++ {
		fmt.Printf("%.1f\n", dec)
		dec += 0.1
	}
}

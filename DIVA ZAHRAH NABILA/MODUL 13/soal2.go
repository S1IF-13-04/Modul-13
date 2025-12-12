package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	target := int(x)
	if float64(target) < x {
		target++
	}
	selesai := false
	for !selesai {
		x += 0.1
		fmt.Printf("%.1f\n", x)
		selesai = (x >= float64(target))
	}
}

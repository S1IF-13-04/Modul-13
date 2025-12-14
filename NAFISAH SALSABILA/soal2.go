package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var kondisi bool = false

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&x)

	batas := math.Ceil(x)

	for kondisi = false; !kondisi; {
		x = x + 0.1
		if x > batas {
			kondisi = true
		} else {
			fmt.Printf("%.1f\n", x)
		}
	}
}

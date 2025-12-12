package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)

	batas := int(math.Ceil(n))
	awal := int(n*10 + 0.000001)
	target := batas * 10

	hasil := awal
	var selesai bool

	for selesai = false; !selesai; {
		hasil = hasil + 1

		if hasil%10 == 0 {
			fmt.Println(hasil / 10)
		} else {
			fmt.Printf("%.1f\n", float64(hasil)/10.0)
		}
		selesai = hasil >= target
	}
}

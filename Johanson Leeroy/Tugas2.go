package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	var bil, bulat float64
	var kondisi bool
	fmt.Scan(&bil)
	bulat = math.Ceil(bil)
	bulat -= bil
	bulat *= 10
	bulat = math.Ceil(bulat)
	var bulatINT = int(bulat)

	for kondisi = false; !kondisi; {
		bil += 0.1
		fmt.Printf("%.1f \n", bil)
		i++
		kondisi = i == bulatINT
	}
}

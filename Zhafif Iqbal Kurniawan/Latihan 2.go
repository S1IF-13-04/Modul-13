package main

import (
	"fmt"
)

func main() {
	var bil float64
	var bol bool
	fmt.Scan(&bil)
	x := int(bil * 10)
	hasil := (x/10 + 1) * 10

	for bol = false; !bol; {
		x++
		if x == hasil {
			fmt.Println(x / 10)
		} else {
			fmt.Printf("%.1f\n", float64(x)/10)
		}
		bol = x >= hasil
	}
}

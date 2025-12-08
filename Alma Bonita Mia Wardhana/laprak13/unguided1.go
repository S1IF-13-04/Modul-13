package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	digit := len(fmt.Sprintf("%d", n))

	fmt.Println("Jumlah digit:", digit)
}

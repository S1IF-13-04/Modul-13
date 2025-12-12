package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&x)

	// Pembulatan ke atas (ceil) secara manual
	batas := float64(int(x))
	if x > batas {
		batas = batas + 1
	}

	hasil := x

	for hasil < batas {
		hasil += 0.1
		fmt.Printf("%.1f\n", hasil)
	}
}

package main

import "fmt"

func main() {
	var bil, digit int
	var kondisi bool
	fmt.Scan(&bil)

	for kondisi = false; !kondisi; {
		bil /= 10
		digit++
		kondisi = bil == 0
	}
	fmt.Print(digit)
}

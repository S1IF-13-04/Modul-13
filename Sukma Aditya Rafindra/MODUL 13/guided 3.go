package main

import "fmt"

func main() {
	var x, y int
	var kondisi, cek bool
	fmt.Scan(&x, &y)

	for kondisi = false; !kondisi; {
		x = x - y
		fmt.Println(x)
		kondisi = x <= 0
	}
	if x == 0 {
		cek = true
	}
	fmt.Println(cek)
}

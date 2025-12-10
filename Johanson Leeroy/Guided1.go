package main

import "fmt"

func main() {
	var kata string
	var jumlah, counter int
	var kondisi bool

	fmt.Scan(&kata, &jumlah)

	for kondisi = false; !kondisi; {
		fmt.Println(kata)
		counter++
		kondisi = counter == jumlah
	}
}

package main

import "fmt"

func main() {
	var kata string
	var jumlah int
	fmt.Scan(&kata, &jumlah)

	i := 0

	for kondisi := false; !kondisi; {
		fmt.Println(kata)
		i++	
		kondisi = (i >= jumlah)
	}
}
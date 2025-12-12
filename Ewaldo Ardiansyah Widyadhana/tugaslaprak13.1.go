package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	jumlahDigit := 0
	temp := angka

	for temp > 0 {
		temp = temp / 10
		jumlahDigit++
	}

	fmt.Println("Jumlah digit:", jumlahDigit)
}

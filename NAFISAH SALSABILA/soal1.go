package main

import "fmt"

func main() {
	var n int
	var kondisi bool = false
	jumlah := 0

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	for kondisi = false; !kondisi; {
		jumlah++
		n = n / 10
		if n == 0 {
			kondisi = true
		}
	}

	fmt.Println("Jumlah digit:", jumlah)
}

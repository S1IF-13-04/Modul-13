package main

import "fmt"

func main() {
	var target, donasi int
	fmt.Scan(&target)

	total := 0
	jumlah := 0

	selesai := false
	for !selesai {
		fmt.Scan(&donasi)
		jumlah++
		total += donasi

		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlah, donasi, total)
		selesai = (total >= target)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, jumlah)
}

package main

import "fmt"

func main() {
	var target, donasi, total, jumlahDonatur int
	fmt.Scan(&target)
	for ok := true; ok; {
		fmt.Scan(&donasi)
		total += donasi
		jumlahDonatur++
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, total)
		ok = total < target
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, jumlahDonatur)
}
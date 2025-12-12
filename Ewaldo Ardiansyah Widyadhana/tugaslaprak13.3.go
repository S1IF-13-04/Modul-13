package main

import "fmt"

func main() {
	var target int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	total := 0
	donatur := 1

	for total < target {
		var donasi int
		fmt.Printf("Donatur %d menyumbang: ", donatur)
		fmt.Scan(&donasi)

		total += donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
			donatur, donasi, total)

		// jika sudah mencapai target → berhenti
		if total >= target {
			fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n",
				total, donatur)
			break
		}

		donatur++
	}
}

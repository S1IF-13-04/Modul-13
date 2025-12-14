package main

import "fmt"

func main() {
	var target, donasi int
	var kondisi bool = false
	total := 0
	donatur := 0

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	for kondisi = false; !kondisi; {
		donatur++
		fmt.Printf("Donatur %d menyumbang: ", donatur)
		fmt.Scan(&donasi)

		total = total + donasi
		fmt.Println("Total terkumpul:", total)

		if total >= target {
			kondisi = true
		}
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur\n", total, donatur)
}

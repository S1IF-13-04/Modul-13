package main

import "fmt"

func main() {
	var target, donasi, total, donatur int
	var selesai bool

	fmt.Scan(&target)

	total = 0
	donatur = 0

	for selesai = false; !selesai; {
		fmt.Scan(&donasi)

		donatur++
		total = total + donasi

		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donatur, donasi, total)

		selesai = total >= target
	}

	fmt.Printf(
		"Target tercapai! Total donasi: %d dari %d donatur.\n",
		total, donatur,
	)
}

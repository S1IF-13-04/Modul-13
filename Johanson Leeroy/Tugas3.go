package main

import "fmt"

func main() {
	var target, donasi, donatur, total int
	var kondisi bool
	donasi = 0
	fmt.Scan(&target)

	for kondisi = false; !kondisi; {
		fmt.Scan(&donasi)
		total += donasi
		donatur++
		fmt.Print("Donatur ", donatur, " menyumbang ", donasi, ". Total terkumpul :", total, "\n")
		kondisi = total >= target
	}
	fmt.Println("Targert tercapai! Total donasi:", total, "dari", donatur, "donatur")
}

package main

import "fmt"

func main() {
	var donasi, totalDonasi, orang, terkumpul int
	var kondisi bool

	fmt.Scan(&totalDonasi)

	for kondisi = false; !kondisi; {
		fmt.Scan(&donasi)
		orang++
		terkumpul = terkumpul + donasi 
		fmt.Println("Donatur ", orang, " memberikan donasi sebesar ", donasi, " total terkumpul ", terkumpul)

		kondisi = totalDonasi <= terkumpul
	}
	fmt.Println("Target Tercapai! Total donasi: ", terkumpul, "dari", orang, "donatur")
}
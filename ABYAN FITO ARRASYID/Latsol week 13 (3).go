package main

import "fmt"

func main() {
	var target, donasi int
	total := 0
	donatur := 0
	var data [100]int

	fmt.Scan(&target)

	for total < target {
		fmt.Scan(&donasi)
		data[donatur] = donasi
		donatur++
		total += donasi
	}

	jumlah := 0
	for i := 0; i < donatur; i++ {
		jumlah += data[i]
		fmt.Println("Donatur", i+1, ": Menyumbang", data[i], ". Total terkumpul:", jumlah,)
	}

	fmt.Println("Target tercapai! Total donasi:", total, "dari", donatur, "donatur.",)
}

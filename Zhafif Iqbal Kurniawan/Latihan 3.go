package main

import "fmt"

func main() {
	var donasi, target, hasil int
	fmt.Scan(&target)
	i := 1
	for bol := false; !bol; {
		fmt.Scan(&donasi)
		hasil = hasil + donasi
		fmt.Println("Donatur", i, ": Menyumbang", donasi, "hasil terkumpul :", hasil)

		i++
		bol = hasil >= target
	}

	fmt.Println("Target tercapai! hasil donasi :", hasil, "dari", i-1, "donatur")

}

package main

import "fmt"

func main() {
	var target, donasi, total, donatur int

	fmt.Scan(&target)

	for kondisi := false; !kondisi; {
		fmt.Scan(&donasi)
		total += donasi
		donatur++
		kondisi = total >= target
	}
	fmt.Println("Target Tercapai! Total donasi :", total, "dari", donatur, "donatur")
}

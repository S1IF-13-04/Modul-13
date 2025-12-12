package main

import "fmt"

func main() {
	var target int
	var sumbang int
	fmt.Scanln(&target)

	donatur := 0
	totalSumbang := 0
	for kondisi := false; !kondisi; {
		fmt.Scanln(&sumbang)
		totalSumbang += sumbang
		fmt.Println("Donatur ", donatur, " Menyumbang ", sumbang, ".", " Total terkumpul ", totalSumbang)
		kondisi = totalSumbang >= target
		donatur++
	}
	fmt.Print("Target tercapai! Total donasi: ", totalSumbang, " dari ", donatur, " donatur")
}

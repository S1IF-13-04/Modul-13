package main
import "fmt"
func main() {
	var donasi, targetdonasi, total int
	fmt.Scan(&targetdonasi)
	i := 1
	for kondisi := false; !kondisi;{
		fmt.Scan(&donasi)
		total = total + donasi
		fmt.Println("Donatur", i,": Menyumbang", donasi, "Total terkumpul :", total)

		i++
		kondisi = total >= targetdonasi
	}

	fmt.Println("Target tercapai! Total donasi :", total)

}

package main
import "fmt"

func main() {
	var target, donasi, total, jumlahDonatur int
	fmt.Scan(&target)
	total = 0
	jumlahDonatur = 0
	for total < target {
		fmt.Scan(&donasi)
		jumlahDonatur++
		total = total + donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, total)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, jumlahDonatur)
}
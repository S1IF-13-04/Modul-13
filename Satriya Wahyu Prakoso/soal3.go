package main
import "fmt"
func main() {
	var target, bayar, jumlah, donatur int
	jumlah = 0
	donatur = 0
	fmt.Scan(&target)
	for done := false; !done; {
		fmt.Scan(&bayar)
		donatur++
		jumlah = jumlah + bayar
		fmt.Printf("Donatur %d : Menyumbang %d. Total terkumpul : %d\n", donatur, bayar, jumlah)
		done = jumlah >= target
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.", jumlah, donatur)
}
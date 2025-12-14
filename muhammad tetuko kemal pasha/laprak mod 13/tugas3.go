package main
import "fmt"
func main() {
	var target int
	fmt.Scan(&target)
	total := 0
	donatur := 1
	for total < target {
		var donasi int
		fmt.Scan(&donasi)
		total += donasi
		fmt.Printf ("Donatur %d: menyumbang %d. Total terkumpul: %d\n", donatur, donasi, total)
		donatur++
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, donatur-1)
}

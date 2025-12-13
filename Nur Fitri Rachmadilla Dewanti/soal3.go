package main
import "fmt"
func main() {
	var target, donasi, total, donatur int
	fmt.Scan(&target)

	for kondisi := false; !kondisi; {
		fmt.Scan(&donasi)
		donatur++
		total = total + donasi

		fmt.Println("Donatur", donatur,": Menyumbang", donasi, 
		". Total terkumpul:", total,
		)

		if total >= target {
			kondisi = true
		}
	}
	fmt.Println("Target tercapai! Total donasi:", total, "dari",
		donatur, "donatur.",
	)
}
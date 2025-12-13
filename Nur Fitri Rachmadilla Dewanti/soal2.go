package main
import "fmt"
func main() {
	var bilangan float64
	var hasil int
	fmt.Scan(&bilangan)

	hasil = int(bilangan)
	if bilangan > float64(hasil) {
		hasil++
	}

	for kondisi := false; !kondisi; {
		bilangan = bilangan + 0.1
		fmt.Printf("%.1f\n", bilangan)

		if bilangan+0.000001 >= float64(hasil) {
			kondisi = true
		}
	}
}
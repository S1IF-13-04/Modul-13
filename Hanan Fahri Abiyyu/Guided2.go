package main
import "fmt"
func main() {
	var angka int
	var kondisi bool
	for kondisi = false; !kondisi; {
		fmt.Scan(&angka)
		
		kondisi = angka > 0
	}
	fmt.Println(angka, "Merupakan bilangan bulat positif")
}

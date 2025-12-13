package main
import "fmt"
func main(){
	var n, m int
	fmt.Scan(&n)
	hasil := 0
	hasil2 := 0
	for kons := false; !kons; {
		fmt.Scan(&m)
		hasil += 1
		hasil2 += m
		fmt.Println("Donatur", hasil, ": menyumbang", m,". Total terkumpul :", hasil2)
		kons = hasil2 >= n
	}
	fmt.Println("Target tercapai! Total donasi :", hasil2, "dari", hasil, "donatur")
}
package main
import "fmt"
func main() {
	var kata string
	var ulang, i int
	fmt.Scan(&kata, &ulang)
	for kondisi := false; !kondisi; {
		fmt.Println(kata)
		i++
		kondisi = i == ulang
	}
}
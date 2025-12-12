package main
import "fmt"
func main() {
	var kata string
	var ulang int
	fmt.Scan(&kata, &ulang)
	i := 0
	for kondisi := false; !kondisi; {
		fmt.Println(kata)
		i++
		kondisi = (i >= ulang)
	}
}

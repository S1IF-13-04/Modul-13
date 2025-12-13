package main
import "fmt"
func main() {
	var kata string
	var n int
	fmt.Scan(&kata, &n)
	
	i := 0
	for kondisi := false; !kondisi; {
		fmt.Println(kata)
		i++
		kondisi = (i >= n)
	}
}
package main
import (
	"fmt"
)
func main() {
	var bil float64
	var kondisi bool
	fmt.Scan(&bil)
	x := int(bil * 10)
	done := (x/10 + 1) * 10

	for kondisi = false; !kondisi;{
		x++
		if x == done{
			fmt.Println(x / 10)
		}else{
			fmt.Printf("%.1f\n", float64(x)/10)
		}
		kondisi = x >= done
	}
}

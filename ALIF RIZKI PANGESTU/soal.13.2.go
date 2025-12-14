package main
import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&x)

	target := math.Ceil(x)
	current := x + 0.1

	
	for {
	
		fmt.Printf("%.1f\n", current)

		if current >= target {
			break
		}

		current += 0.1
	}
}

package main
import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Scan(&a)

	target := math.Ceil(a)

	for {
		a += 0.1
		if a >= target {
			a = target
			fmt.Printf("%.1f\n", a)
			break
		}
		fmt.Printf("%.1f\n", a)
	}
}

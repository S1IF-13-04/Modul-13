package main
import "fmt"
func main() {
	var x float64
	fmt.Scan(&x)
	batasAtas := float64(int(x)) + 1
	for v := x + 0.1; v <= batasAtas; v += 0.1 {
		fmt.Printf("%.1f\n", v)
	}
}
package main
import "fmt"
func main() {
	var n float64
	fmt.Scan(&n)
	i := int(n*10 + 0.5)
	t := ((i / 10) + 1) * 10
	for k := false; !k; {
		fmt.Printf("%.1f\n", float64(i)/10)
		i++
		k = i == t
	}
	fmt.Printf("%.0f\n",float64(i)/10)
}

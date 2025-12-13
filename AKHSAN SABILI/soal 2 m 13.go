package main
import ("fmt"
		"math")
func main(){
	var n float64
	fmt.Scan(&n)
	bulat := math.Ceil(n)	
	for n < bulat{
		n += 0.1
		if n >= bulat{
			n = bulat 
			fmt.Printf("%.1f\n", n)
			break
		}
		fmt.Printf("%.1f\n", n)
	}
}
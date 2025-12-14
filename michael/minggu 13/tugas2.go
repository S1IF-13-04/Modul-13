package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	n := int(x * 10)

	for {
		n += 1
		fmt.Printf("%.1f\n", float64(n)/10)

		if n%10 == 0 {
			break
		}
	}
}

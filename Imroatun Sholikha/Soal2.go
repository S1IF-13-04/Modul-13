package main

import "fmt"

func main() {
	var x float64
	var batas, i int
	var kondisi bool

	fmt.Scan(&x)

	batas = int(x + 0.999999)
	i = int(x * 10)

	for kondisi = false; !kondisi; {
		i++
		if i == batas*10 {
			fmt.Println(batas)
			kondisi = true
		} else {
			fmt.Printf("%.1f\n", float64(i)/10)
		}
	}
}

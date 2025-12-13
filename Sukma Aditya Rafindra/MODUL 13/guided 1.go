package main

import "fmt"

func main() {
	var x string
	var y int
	var kondisi bool
	fmt.Scan(&x, &y)

	i := 0
	for kondisi = false; !kondisi; {
		fmt.Println(x)
		i++
		kondisi = (i >= y)
	}
}

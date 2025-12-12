package main

import "fmt"

func main() {
	var x int
	var bol bool
	fmt.Scan(&x)
	i := 0
	for bol = false; !bol; {
		x /= 10
		i++

		bol = x == 0
	}
	fmt.Println(i)
}

package main

import "fmt"

func main() {
	var len, num int
	fmt.Scan(&num)

	for num > 0 {
		len++
		num /= 10
	}
	fmt.Print(len)
}

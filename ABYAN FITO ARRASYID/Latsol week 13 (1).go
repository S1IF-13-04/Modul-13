package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println(1)
		return
	}

	count := 0
	for n != 0 {
		count++
		n /= 10
	}

	fmt.Println(count)
}

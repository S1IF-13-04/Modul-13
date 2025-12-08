package main

import "fmt"

func main() {
	var bilangan, digit int
	fmt.Scan(&bilangan)
	for ok := true; ok; {
		bilangan = bilangan / 10
		digit++
		ok = bilangan > 0
	}
	fmt.Println(digit)
}
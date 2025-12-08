package main

import "fmt"

func main() {
	var num int
	for posi := false; !posi; {
		fmt.Print("enter an int:")
		fmt.Scan(&num)
		posi = num > 0
	}
	fmt.Print(num, " is a positive int")
}

package main

import "fmt"

func main() {
	var word string
	var num int
	fmt.Scanln(&word, &num)
	for done := false; !done; {
		fmt.Println(word)
		num--
		done = num == 0
	}
}

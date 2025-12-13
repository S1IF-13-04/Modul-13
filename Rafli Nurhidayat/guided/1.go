package main

import "fmt"

func main() {
	var word string
	var repeat int
	var kondisi bool

	fmt.Scan(&word, &repeat)

	for kondisi = false; !kondisi; {
		fmt.Println(word)
		repeat -= 1
		kondisi = repeat <= 0
	}
}

package main

import "fmt"

func main() {
	var word string
	var repedtitions int
	fmt.Scan(&word, &repedtitions)
	counter := 0
	for done := false; !done; {
		fmt.Println(word)
		counter++
		done = (counter >= repedtitions)
	}
}

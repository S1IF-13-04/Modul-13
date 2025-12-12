package main

import "fmt"

func main() {
	var text string
	var n int
	fmt.Scan(&text, &n)
	i := 0
	for kondisi := false; !kondisi; {
		fmt.Println(text)
		i++
		kondisi = (i >= n)
	}
}

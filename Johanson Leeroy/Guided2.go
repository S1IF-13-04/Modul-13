package main

import "fmt"

func main() {
	var bil int
	var kondisi bool
	bil = 0
	for kondisi = false; !kondisi; {
		fmt.Scan(&bil)
		kondisi = bil > 0
	}
	fmt.Println(bil, "adalah bilangan bulat positif")
}

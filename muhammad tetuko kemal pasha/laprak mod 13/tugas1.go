package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	jumlahDigit := len(fmt.Sprintf("%d", bilangan))
	fmt.Println("Jumlah digit:", jumlahDigit)
}

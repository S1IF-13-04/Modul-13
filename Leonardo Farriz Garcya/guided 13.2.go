package main

import "fmt"

func main (){
	var n int
	var kondisi bool
	
	for kondisi = false; !kondisi;{
		fmt.Scan(&n)
		kondisi = n > 0
		
	}
	fmt.Println(n, "bilangan bulat positif")

}

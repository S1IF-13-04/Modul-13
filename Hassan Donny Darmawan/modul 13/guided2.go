package main
import "fmt"

func main (){
	var a int
	var kondisi bool
	
	for kondisi= false; !kondisi; {
		fmt.Scan(&a)
		kondisi= a>0
	}
	fmt.Println(a,"adalah bilangan positif")
}
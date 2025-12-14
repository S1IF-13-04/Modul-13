package main
import "fmt"

func main(){
	var bilangan int
	var kondisi bool
	for kondisi = false; !kondisi; {
		fmt.Scan(&bilangan)
		kondisi = bilangan > 0
	}
	fmt.Print(bilangan, " adalah bilangan bulat positif")
}
package main
import "fmt"

func main() {
	var bilangan int
	var jumlahDigit int
	fmt.Scan(&bilangan)
	jumlahDigit = 0
	for bilangan > 0 {
		jumlahDigit++
		bilangan = bilangan / 10
	}
	fmt.Println(jumlahDigit)
}

package main
import "fmt"

func main() {
	var bilangan float64
	var batas, nilai int
	fmt.Scan(&bilangan)
	nilai = int(bilangan * 10)
	batas = nilai / 10
	if nilai%10 != 0 {
		batas++
	}
	for nilai < batas*10 {
		nilai = nilai + 1
		if nilai%10 == 0 {
			fmt.Println(nilai / 10)
		} else {
			fmt.Printf("%d.%d\n", nilai/10, nilai%10)
		}
	}
}
package main
import "fmt"
func main(){
	var p int
	var q string
	fmt.Scan(&q, &p)
	for i := 1; i <= p; i++{
		fmt.Println(q)
	}
}
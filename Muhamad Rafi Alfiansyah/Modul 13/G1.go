package main
import "fmt"
func main(){
	var s string
	var n int
	fmt.Scan(&s,&n)
	i := 0
	for k := false; !k;{
		fmt.Println(s)
		i++
		k = i >= n
	}
}
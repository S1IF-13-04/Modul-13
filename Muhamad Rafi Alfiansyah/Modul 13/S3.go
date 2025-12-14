package main
import "fmt"
func main(){
	var n,d,t int
	k := false
	h := 0
	fmt.Scan(&n)
	for !k{
		fmt.Scan(&d)
		h ++
		t += d
		fmt.Printf("Donatur %d: Menyumbang %d. Total Terkumpul %d\n",h,d,t)
		k = t >= n
	}
	fmt.Printf("Target Tercapai! Total donasi: %d dari %d donatur",t,h)
}
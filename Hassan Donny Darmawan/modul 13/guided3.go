package main
import "fmt"

func main (){
	var a,b int
	fmt.Scan(&a,&b)

	for kondisi:= false;!kondisi;{
		a-=b
		kondisi = a<=0
		fmt.Println(a)
	}
	fmt.Print(a==0)
}
package main
import "fmt"

func main (){
	var target, total, donatur int
	fmt.Scan(&target)
	jumlah:=0
	for kondisi :=true; kondisi;{
		fmt.Scan(&donatur)
		jumlah++
		total += donatur 
		fmt.Println("Donatur",jumlah,": menyumbang",donatur,". Total terkumpul:", total)
		kondisi = total<target
	}
	fmt.Println("Target tercapai! Total donasi:", total, "dari",jumlah,"donatur")
}
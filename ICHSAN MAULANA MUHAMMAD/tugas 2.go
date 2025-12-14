package main

import "fmt"

func main() {
	var bil float64
	var nilai, batas int
	var selesai bool

	fmt.Scan(&bil)
	nilai = int(bil * 10)
	batas = (int(bil) + 1) * 10

	for selesai = false; !selesai; {
		nilai = nilai + 1
		selesai = nilai > batas
		if !selesai {
			fmt.Printf("%.1f\n", float64(nilai)/10)
		}
	}
}

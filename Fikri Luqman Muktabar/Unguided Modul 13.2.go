package main

import "fmt"

func main() {
    var bilangan float64
    fmt.Scan(&bilangan)

    pembulatan := int(bilangan)
    if float64(pembulatan) < bilangan {
        pembulatan++
    }

    hasil := bilangan

    for kondisi := true; kondisi; {
        hasil += 0.1

        hasil = float64(int(hasil*10+0.0001)) / 10

        if hasil == float64(int(hasil)) {
            fmt.Println(int(hasil))
        } else {
            fmt.Printf("%.1f\n", hasil)
        }

        kondisi = hasil < float64(pembulatan)
    }
}
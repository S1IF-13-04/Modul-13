package main

import "fmt"

func main() {
    var x float64
    var start, end int
    var kondisi bool

    fmt.Scan(&x)

    end = (int(x) + 1) * 10
    start = int(x*10) + 1

    for kondisi = false; !kondisi; {
        hasil := float64(start) / 10

        if hasil == float64(int(hasil)) {
            fmt.Printf("%d\n", int(hasil))
        } else {
            fmt.Printf("%.1f\n", hasil)
        }

        start++
        kondisi = start > end
    }
}

package main

import "fmt"

func main() {
    var bilangan int
    fmt.Scan(&bilangan)

    jumlahDigit := 0

    for bilangan > 0 {
        bilangan = bilangan / 10
        jumlahDigit++
    }

    fmt.Println(jumlahDigit)
}
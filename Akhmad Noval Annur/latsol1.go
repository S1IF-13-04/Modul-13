package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    count := 0
    for selesai := false; !selesai; {
        count++
        n = n / 10
        selesai = n == 0
    }

    fmt.Println(count)
}

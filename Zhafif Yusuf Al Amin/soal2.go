package main

import "fmt"

func main() {
    var n float64
    fmt.Scan(&n)

    target := float64(int(n + 0.9999999))
    i := 0.1

    for kondisi := false; !kondisi; {
        n += i
        if n >= target {
            fmt.Printf("%.0f\n", n) 
            kondisi = true
        } else {
            fmt.Printf("%.1f\n", n)
        }
    }
}

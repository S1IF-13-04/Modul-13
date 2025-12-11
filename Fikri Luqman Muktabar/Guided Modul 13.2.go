package main

import "fmt"

func main() {
    var bilangan int
	fmt.Scan(&bilangan)

    for bilangan <= 0 {
        fmt.Scan(&bilangan)
    }

    fmt.Printf("%d adalah bilangan bulat positif\n", bilangan)
}
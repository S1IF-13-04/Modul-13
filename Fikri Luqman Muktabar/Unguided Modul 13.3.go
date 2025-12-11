package main

import "fmt"

func main() {
    var target, donasi int
    fmt.Scan(&target)

    totalDonasi := 0
	jumlahDonatur := 0

    for kondisi := true; kondisi; {
        fmt.Scan(&donasi)

        jumlahDonatur++
        totalDonasi += donasi

        fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
            jumlahDonatur, donasi, totalDonasi)

        kondisi = totalDonasi < target
    }

    fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}
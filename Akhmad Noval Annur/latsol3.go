package main

import "fmt"

func main() {
    var target int
    var donasi int
    total := 0
    donatur := 0

    fmt.Scan(&target)

    for selesai := false; !selesai; {
        fmt.Scan(&donasi)

        donatur++
        total = total + donasi

        fmt.Printf(
            "Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
            donatur, donasi, total,
        )

        selesai = total >= target
    }

    fmt.Printf(
        "Target tercapai! Total donasi: %d dari %d donatur.\n",
        total, donatur,
    )
}

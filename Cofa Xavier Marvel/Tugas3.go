package main

import "fmt"

func main() {
	var Donors, Donated, TotCollected, Goal int

	fmt.Scan(&Goal)

	for ; Goal > 0; Goal -= Donated {
		fmt.Scan(&Donated)
		Donors++
		TotCollected += Donated
		fmt.Printf("Donor %d: Donated %d. Total collected: %d\n", Donors, Donated, TotCollected)
	}
	if Goal < 0 {
		fmt.Printf("\nTarget reached! Total donation: %d from %d donors.", TotCollected, Donors)
	}
}

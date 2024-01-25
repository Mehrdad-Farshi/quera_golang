package main

import (
	"fmt"
)

func main() {
	var tax float64 = 0.0
	var income int
	const floor1 float64 = 0.05
	const floor2 float64 = 0.10
	const floor3 float64 = 0.15
	const floor4 float64 = 0.20

	const limit1 int = 10000
	const limit2 int = 40000
	const limit3 int = 50000
	fmt.Print(" tell me how much is your income   : ")
	fmt.Scanln(&income)
	}

	tax1 = min(income, limit1) * rate1
    tax2 = max(0, min(income, limit2) - limit1) * rate2
    tax3 = max(0, min(income, limit3) - limit2) * rate3
    tax4 = max(0, income - limit3) * rate4

    total_tax = tax1 + tax2 + tax3 + tax4


	// if deposit < 10000 {
	// tax += floor1 * float64(deposit)
	// }
	// if 10000 < deposit && deposit < 40000 {
	// tax += floor2 * (float64(deposit - 10000))
	// }
	// if 40000 < deposit && deposit < 50000 {
	// tax += floor3 * (float64(deposit - 40000))
	// }
	// if deposit > 100000 {
	// tax += floor4 * (float64(deposit - 100000))
	// }
	fmt.Printf("you have %d and you have to pay %2.f for your tax\n", deposit, tax)

}

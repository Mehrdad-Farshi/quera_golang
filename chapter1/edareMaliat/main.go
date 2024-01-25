package main

import (
	"fmt"
)

func main() {
	var tax float64 = 0.0
	var deposit int
	const floor1 float64 = 0.05
	const floor2 float64 = 0.10
	const floor3 float64 = 0.15
	const floor4 float64 = 0.20
	fmt.Print(" tell me how much you have  : ")
	fmt.Scanln(&deposit)

	if deposit < 10000 {
		tax += floor1 * float64(deposit)
	}
	if 10000 < deposit && deposit < 40000 {
		tax += floor2 * (float64(deposit - 10000))
	}
	if 40000 < deposit && deposit < 50000 {
		tax += floor3 * (float64(deposit - 40000))
	}
	if deposit > 100000 {
		tax += floor4 * (float64(deposit - 100000))
	}
	fmt.Printf("you have %d and you have to pay %2.f for your tax\n", deposit, tax)

	// for i := 1; i <= deposit; i++ {
	// 	if i%mazrab == 0 {
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }
}

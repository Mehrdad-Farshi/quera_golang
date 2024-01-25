package main

import (
	"fmt"
	"strings"
)

func main() {
	var lastNumber, mazrab int
	// fmt.Print(" two numbers: (mazrab) and then  Type  lastNumber ( end of game) :")
	fmt.Scanln(&mazrab, &lastNumber)
	// fmt.Println("lastNumber is equal to = ", lastNumber, ", mazrab is equal to = ", mazrab)

	for i := 1; i <= lastNumber; i++ {
		if i%mazrab == 0 {
			fmt.Println(strings.Repeat("Hope ", i/mazrab))
		} else {
			fmt.Println(i)
		}
	}
}

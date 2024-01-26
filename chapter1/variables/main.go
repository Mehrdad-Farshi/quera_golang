package main

import "fmt"

func main() {
	var tic_tac_toe = [4][3]bool{{true, true, false},
		{false, false, false},
		{true, false, true},
		{true, false, true}}

	// fmt.Println(tic_tac_toe)

	for index, element := range tic_tac_toe {
		fmt.Printf("Element %v At Index %d\n", element, index)
	}
}

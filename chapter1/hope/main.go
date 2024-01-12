package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&number)
	fmt.Println("number:", number)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("number:", number)
}

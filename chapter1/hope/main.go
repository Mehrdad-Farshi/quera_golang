package main

import (
	"errors"
	"fmt"
)

func main() {
	var q, p int
	fmt.Println("Type two numbers: q ( end of game) and then p (mazrab) ")
	if q == 0 {
		return "error", errors.New("negetive number")
	}
	fmt.Scanln(&q, &p)
	fmt.Println("q is equal to = ", q, ", p is equal to = ", p)
}

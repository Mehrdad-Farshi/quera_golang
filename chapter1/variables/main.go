package main

import "fmt"

var (
	name      string = "mehrdad"
	age       int    = 26
	isStudent bool   = true
)

var (
	income  float64
	taxRate float64
)

const (
	floor1 int = 111
)

var (
	seasons [4]string
)

func main() {
	seasons[0] = "spring"
	seasons[1] = "summer"
	seasons[2] = "fall"
	seasons[3] = "winter"
	for i := 0; i < len(seasons); i++ {
		fmt.Printf("at index %v Elemnet  = %v  \n", i, seasons[i])
	}
	fmt.Printf("-------------------------------------------------\n")
	for range seasons {
		fmt.Printf("iterated \n")
	}
}

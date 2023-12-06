package main

import "fmt"

func main() {
	var bigNumber1 int16 = 32766
	var bigNumber2 int16 = 32766
	var result int16 = bigNumber1 + bigNumber2
	fmt.Println(result)
}

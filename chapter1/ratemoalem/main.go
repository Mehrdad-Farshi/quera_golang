package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var teacherName []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	teachers, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("invalid teacher")
	}

	for teacher := 0; teacher < teachers; teacher++ {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		tName := scanner.Text()
		teacherName = append(teacherName, tName)

	}

	fmt.Println(teacherName)
}

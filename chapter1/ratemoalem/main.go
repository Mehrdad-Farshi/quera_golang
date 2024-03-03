package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	teacherName          []string
	teacherStudentScores []int
	teacherDescriptive   []string
)

func main() {

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

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal()
			return
		}
		numStrings := strings.Fields(input)
		sum := 0
		count := 0
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}
			sum += num
			count++
		}
		average := float64(sum) / float64(count)
		teacherStudentScores = append(teacherStudentScores, int(average))
	}
	for _, teacherStudentScore := range teacherStudentScores {
		if teacherStudentScore < 40 {
			teacherDescriptive = append(teacherDescriptive, "Fair")
		}
		if 40 <= teacherStudentScore && teacherStudentScore < 60 {
			teacherDescriptive = append(teacherDescriptive, "Good")
		}
		if 60 <= teacherStudentScore && teacherStudentScore < 80 {
			teacherDescriptive = append(teacherDescriptive, "Very Good")
		}
		if 80 <= teacherStudentScore {
			teacherDescriptive = append(teacherDescriptive, "Excellent")
		}
	}
	for i := 0; i < len(teacherName); i++ {
		fmt.Println(teacherName[i], teacherDescriptive[i])
	}
}

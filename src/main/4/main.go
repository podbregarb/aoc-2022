package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func main() {
	fmt.Println("Test input")
	printAnswers(testInput)
	fmt.Println("Input")
	printAnswers(input)
}

func printAnswers(s string) {
	first, second := getAnswers(s)
	fmt.Printf("First answer: %d\n", first)
	fmt.Printf("Second answer: %d\n", second)
}

func getAnswers(s string) (int, int) {
	split := strings.Split(s, "\n")
	containedRanges := 0
	overlappedRanges := 0
	for i := 0; i < len(split)-1; i++ {
		splitRow := strings.Split(split[i], ",")
		range1 := strings.Split(splitRow[0], "-")
		range2 := strings.Split(splitRow[1], "-")
		range11, _ := strconv.Atoi(range1[0])
		range12, _ := strconv.Atoi(range1[1])
		range21, _ := strconv.Atoi(range2[0])
		range22, _ := strconv.Atoi(range2[1])
		if range11 >= range21 && range12 <= range22 || (range21 >= range11 && range22 <= range12) {
			containedRanges++
		}
		if range11 <= range22 && range21 <= range12 {
			overlappedRanges++
		}
	}
	return containedRanges, overlappedRanges
}

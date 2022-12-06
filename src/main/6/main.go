package main

import (
	_ "embed"
	"fmt"
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
	return getIndex(s, 4), getIndex(s, 14)
}

func getIndex(s string, numOfChars int) int {
	for i := numOfChars - 1; i <= len(s)-1; i++ {
		chars := make(map[uint8]int, 0)
		for j := 0; j < numOfChars; j++ {
			chars[s[i-j]] = 0
		}
		if len(chars) == numOfChars {
			return i + 1
		}
	}
	return -1
}

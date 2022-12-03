package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

var alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
	priority := 0
	for i := range split {
		firstHalf := split[i][0 : len(split[i])/2]
		secondHalf := split[i][len(split[i])/2 : len(split[i])]
		for i2 := range firstHalf {
			u := string(firstHalf[i2])
			if strings.Contains(secondHalf, u) {
				priority += strings.Index(alphabet, u) + 1
				break
			}
		}
	}

	priorityBadges := 0
	for i := 0; i < len(split)-1; i += 3 {
		firstLine := split[i]
		secondLine := split[i+1]
		thirdLine := split[i+2]
		for i2 := range firstLine {
			u := string(firstLine[i2])
			if strings.Contains(secondLine, u) && strings.Contains(thirdLine, u) {
				priorityBadges += strings.Index(alphabet, u) + 1
				break
			}
		}
	}
	return priority, priorityBadges
}

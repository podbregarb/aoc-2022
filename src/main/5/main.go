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
	fmt.Printf("First answer: %s\n", first)
	fmt.Printf("Second answer: %s\n", second)
}

func getAnswers(s string) (string, string) {
	return getTopCrates(s, false), getTopCrates(s, true)
}

func getTopCrates(s string, retainOrder bool) string {
	split := strings.Split(s, "\n")
	instructionsStart := getStartOfInstructions(split)
	cratesDefinition := split[:instructionsStart]
	cratesIndexes := strings.Fields(split[instructionsStart-1])
	numOfCrates := len(cratesIndexes)
	crates := make([][]string, numOfCrates)
	for crateRowIndex := len(cratesDefinition) - 2; crateRowIndex >= 0; crateRowIndex-- {
		row := cratesDefinition[crateRowIndex]
		for crateNumberIndex := range cratesIndexes {
			indexOfCrate := strings.Index(split[instructionsStart-1], cratesIndexes[crateNumberIndex])
			if len(row) > indexOfCrate && (string(row[indexOfCrate]) != " ") {
				crates[crateNumberIndex] = append(crates[crateNumberIndex], string(row[indexOfCrate]))
			}
		}
	}
	instructions := split[instructionsStart+1:]
	for instructionIndex := 0; instructionIndex < len(instructions)-1; instructionIndex++ {
		fields := strings.Fields(instructions[instructionIndex])
		howMany, _ := strconv.Atoi(fields[1])
		from, _ := strconv.Atoi(fields[3])
		to, _ := strconv.Atoi(fields[5])
		if retainOrder {
			crates[to-1] = append(crates[to-1], crates[from-1][len(crates[from-1])-howMany:]...)
		} else {
			crates[to-1] = append(crates[to-1], reverse(crates[from-1][len(crates[from-1])-howMany:])...)
		}
		crates[from-1] = crates[from-1][:len(crates[from-1])-howMany]
	}
	topCrates := ""
	for crate := range crates {
		topCrates += crates[crate][len(crates[crate])-1]
	}
	return topCrates
}

func getStartOfInstructions(s []string) int {
	for i := range s {
		if s[i] == "" {
			return i
		}
	}
	return -1
}

func reverse[S ~[]E, E any](s S) S {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

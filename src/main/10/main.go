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
	cycle := 0
	valueX := 1
	result := 0
	cycles := []int{20, 60, 100, 140, 180, 220}
	for i := 0; i < len(split); i++ {
		result = addToResult(cycles, cycle, valueX, result)
		cycle++
		if split[i] != "noop" {
			result = addToResult(cycles, cycle, valueX, result)
			cycle++
			addx, _ := strconv.Atoi(strings.Fields(split[i])[1])
			valueX += addx
		}
	}
	fmt.Println("\n")
	return result, 0
}

func addToResult(cycles []int, cycle int, valueX int, result int) int {
	if contains([]int{valueX - 1, valueX, valueX + 1}, cycle%40) {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if contains([]int{40, 80, 120, 160, 200, 240}, cycle+1) {
		fmt.Print("\n")
	}
	if contains(cycles, cycle+1) {
		result += valueX * (cycle + 1)
	}
	return result
}

/*
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
*/

func contains(s []int, e int) bool {
	for a := range s {
		if s[a] == e {
			return true
		}
	}
	return false
}

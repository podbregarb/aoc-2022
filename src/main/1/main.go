package main

import (
	_ "embed"
	"fmt"
	"sort"
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
	highestCalories, highestThree := getAnswers(s)
	fmt.Printf("First answer: %d\n", highestCalories)
	fmt.Printf("Second answer: %d\n", highestThree)
}

func getAnswers(s string) (int, int) {
	calories := make([]int, 1)
	ss := strings.Split(s, "\n")
	skrat := 0
	for i := range ss {
		atoi, err := strconv.Atoi(ss[i])
		if err != nil || ss[i] == "" {
			skrat++
			calories = append(calories, 0)
		} else {
			calories[skrat] = calories[skrat] + atoi
		}
	}
	sort.Ints(calories)
	highestCalories := calories[len(calories)-1]
	highestThree := 0
	for i := len(calories) - 1; i >= len(calories)-3; i-- {
		highestThree += calories[i]
	}
	return highestCalories, highestThree
}

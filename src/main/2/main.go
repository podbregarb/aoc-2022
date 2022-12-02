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

var Lose = 0

var Draw = 3

var Win = 6

var Scores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var OutcomeScores = map[string]map[string]int{
	"A": {"X": Draw, "Y": Win, "Z": Lose},
	"B": {"X": Lose, "Y": Draw, "Z": Win},
	"C": {"X": Win, "Y": Lose, "Z": Draw},
}

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
	ss := strings.Split(s, "\n")
	score := 0
	for i := range ss {
		if len(ss[i]) > 0 {
			outcome := strings.Split(ss[i], " ")
			score += Scores[outcome[1]] + OutcomeScores[outcome[0]][outcome[1]]
		}
	}
	score2 := 0
	for i := range ss {
		if len(ss[i]) > 0 {
			outcome := strings.Split(ss[i], " ")
			res := 0
			switch outcome[1] {
			case "X":
				res = Lose
			case "Y":
				res = Draw
			case "Z":
				res = Win
			}
			score2 += Scores[getMyOutcome(outcome, res)] + res
		}
	}
	return score, score2
}

func getMyOutcome(outcome []string, res int) string {
	me := OutcomeScores[outcome[0]]
	for s2 := range me {
		if me[s2] == res {
			return s2
		}
	}
	return ""
}

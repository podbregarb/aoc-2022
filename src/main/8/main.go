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
	visibleTrees := 0
	scenicScore := 0
	for i := 0; i < len(split); i++ {
		for j := 0; j < len(split[i]); j++ {
			visibleFromLeft, scenicScoreLeft := isVisibleFromLeft(split, i, j)
			visibleFromRight, scenicScoreRight := isVisibleFromRight(split, i, j)
			visibleFromTop, scenicScoreTop := isVisibleFromTop(split, i, j)
			visibleFromBelow, scenicScoreBelow := isVisibleFromBelow(split, i, j)
			if visibleFromLeft || visibleFromRight || visibleFromTop || visibleFromBelow {
				visibleTrees++
			}
			newScenicScore := scenicScoreLeft * scenicScoreRight * scenicScoreTop * scenicScoreBelow
			if newScenicScore > scenicScore {
				scenicScore = newScenicScore
			}
		}
	}
	return visibleTrees, scenicScore
}

func isVisibleFromTop(split []string, i int, j int) (bool, int) {
	for index := i; index > 0; index-- {
		if split[index-1][j] >= split[i][j] {
			return false, i - index + 1
		}
	}
	return true, i
}

func isVisibleFromBelow(split []string, i int, j int) (bool, int) {
	for index := i; index < len(split)-1; index++ {
		if split[index+1][j] >= split[i][j] {
			return false, index - i + 1
		}
	}
	return true, len(split) - i - 1
}

func isVisibleFromLeft(split []string, i int, j int) (bool, int) {
	for index := j; index > 0; index-- {
		if split[i][index-1] >= split[i][j] {
			return false, j - index + 1
		}
	}
	return true, j
}

func isVisibleFromRight(split []string, i int, j int) (bool, int) {
	for index := j; index < len(split[i])-1; index++ {
		if split[i][index+1] >= split[i][j] {
			return false, index - j + 1
		}
	}
	return true, len(split[i]) - j - 1
}

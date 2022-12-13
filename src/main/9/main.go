package main

import (
	_ "embed"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed testInput2.txt
var testInput2 string

//go:embed input.txt
var input string

func main() {
	fmt.Println("Test input")
	printAnswers(testInput)
	fmt.Println("Test input 2")
	printAnswers(testInput2)
	fmt.Println("Input")
	fmt.Println("2846 too high")
	printAnswers(input)
}

func printAnswers(s string) {
	first, second := getAnswers(s)
	fmt.Printf("First answer: %d\n", first)
	fmt.Printf("Second answer: %d\n", second)
}

func getAnswers(s string) (int, int) {
	return len(getVisitedPositions(s, 2)), len(getVisitedPositions(s, 10))
}

func getVisitedPositions(s string, ropeLength int) []string {
	ropePositions := make([]point, 0)
	for i := 0; i < ropeLength; i++ {
		ropePositions = append(ropePositions, point{0, 0})
	}
	split := strings.Split(s, "\n")
	visitedTailPositions := []string{"0,0"}
	for i := range split {
		instructions := strings.Fields(split[i])
		direction := instructions[0]
		step, _ := strconv.Atoi(instructions[1])
		for oneStep := 0; oneStep < step; oneStep++ {
			switch direction {
			case "R":
				ropePositions[0] = moveRight(ropePositions[0])
			case "L":
				ropePositions[0] = moveLeft(ropePositions[0])
			case "U":
				ropePositions[0] = moveUp(ropePositions[0])
			case "D":
				ropePositions[0] = moveDown(ropePositions[0])
			}
			for ropePartIndex := 1; ropePartIndex < len(ropePositions); ropePartIndex++ {
				if !arePointsTouching(ropePositions[ropePartIndex-1], ropePositions[ropePartIndex]) {
					ropePositions[ropePartIndex] = moveToHead(ropePositions[ropePartIndex], ropePositions[ropePartIndex-1], ropePartIndex)
				}
			}
			stringPoint := fmt.Sprintf("%d,%d", ropePositions[len(ropePositions)-1].x, ropePositions[len(ropePositions)-1].y)
			_, contained := contains(visitedTailPositions, stringPoint)
			if !contained {
				visitedTailPositions = append(visitedTailPositions, stringPoint)
			}
		}
	}
	return visitedTailPositions
}

func contains(s []string, e string) (int, bool) {
	for i, a := range s {
		if a == e {
			return i, true
		}
	}
	return -1, false
}

func moveToHead(t point, h point, ropePartIndex int) point {
	d := distance(t, h)
	if d > 3 {
		errors.New("distance is greater than 3")
	}
	if d >= 2 {
		if t.x-h.x > 0 {
			t = moveLeft(t)
		}
		if h.x-t.x > 0 {
			t = moveRight(t)
		}
		if t.y-h.y > 0 {
			t = moveDown(t)
		}
		if h.y-t.y > 0 {
			t = moveUp(t)
		}
	}
	return t
}

type point struct {
	x int
	y int
}

func moveRight(p point) point {
	return point{p.x + 1, p.y}
}
func moveLeft(p point) point {
	return point{p.x - 1, p.y}
}
func moveUp(p point) point {
	return point{p.x, p.y + 1}
}
func moveDown(p point) point {
	return point{p.x, p.y - 1}
}

func arePointsTouching(p point, r point) bool {
	return distance(p, r) <= 1
}

func distance(p point, r point) float64 {
	return math.Sqrt(math.Pow(float64(p.x-r.x), 2) + math.Pow(float64(p.y-r.y), 2))
}

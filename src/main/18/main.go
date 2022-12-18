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
	first, second := getAnswers(s)
	fmt.Printf("First answer: %d\n", first)
	fmt.Printf("Second answer: %d\n", second)
}

func getAnswers(s string) (int, int) {
	cubes := getCubes(s)
	covered := 0
	for _, c := range cubes {
		for _, neighbor := range neighbors(c) {
			if !contains(neighbor, cubes) {
				covered++
			}
		}
	}

	cubeSize := getCubeSize(cubes)
	filledSpace := make(map[cube]bool)
	for _, c := range cubes {
		filledSpace[c] = true
	}
	exterior := make(map[cube]bool)
	return covered, countExternalSurface(cube{0, 0, 0}, filledSpace, exterior, cubeSize)
}

func countExternalSurface(c cube, filledSpace map[cube]bool, exterior map[cube]bool, cubeSize cube) int {
	if exterior[c] || c.x < -1 || c.x > cubeSize.x+1 || c.y < -1 || c.y > cubeSize.y+1 || c.z < -1 || c.z > cubeSize.z+1 {
		return 0
	}
	if filledSpace[c] {
		return 1
	}
	exterior[c] = true
	count := 0
	for _, neighbor := range neighbors(c) {
		count += countExternalSurface(neighbor, filledSpace, exterior, cubeSize)
	}
	return count
}

func neighbors(c cube) []cube {
	return []cube{
		{x: c.x - 1, y: c.y, z: c.z},
		{x: c.x + 1, y: c.y, z: c.z},
		{x: c.x, y: c.y - 1, z: c.z},
		{x: c.x, y: c.y + 1, z: c.z},
		{x: c.x, y: c.y, z: c.z - 1},
		{x: c.x, y: c.y, z: c.z + 1},
	}
}

func getCubeSize(cubes []cube) cube {
	xes := make([]int, 0)
	yons := make([]int, 0)
	zes := make([]int, 0)
	for _, c := range cubes {
		xes = append(xes, c.x)
		yons = append(yons, c.y)
		zes = append(zes, c.z)
	}
	sort.Ints(xes)
	sort.Ints(yons)
	sort.Ints(zes)
	return cube{xes[len(xes)-1], yons[len(yons)-1], zes[len(zes)-1]}
}

func contains(c cube, cubes []cube) bool {
	for _, otherCube := range cubes {
		if c.x == otherCube.x && c.y == otherCube.y && c.z == otherCube.z {
			return true
		}
	}
	return false
}

func getCubes(s string) []cube {
	split := strings.Split(s, "\n")
	cubes := make([]cube, 0)
	for _, cubeCoordinates := range split {
		cubes = append(cubes, getCube(cubeCoordinates))
	}
	return cubes
}

func getCube(cubeCoordinates string) cube {
	coordinates := strings.Split(cubeCoordinates, ",")
	first, _ := strconv.Atoi(coordinates[0])
	second, _ := strconv.Atoi(coordinates[1])
	third, _ := strconv.Atoi(coordinates[2])
	return cube{first, second, third}
}

type cube struct {
	x int
	y int
	z int
}

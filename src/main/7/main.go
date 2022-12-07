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
	path := make([]string, 0)
	var dirSizes = make(map[string]int, 0)
	var subdirectories = make(map[string][]string)
	paths := make([]string, 0)
	split := strings.Split(s, "\n$ ")
	for i := 1; i < len(split); i++ {
		commandBlock := split[i]
		command := strings.Split(commandBlock, "\n")
		firstRow := strings.Fields(command[0])
		if firstRow[0] == "cd" {
			if firstRow[1] == ".." {
				path = path[:len(path)-1]
			} else {
				path = append(path, firstRow[1])
				dirSizes[strings.Join(path, "/")] = 0
				paths = append(paths, strings.Join(path, "/"))
			}
		} else if firstRow[0] == "ls" {
			absolutePath := strings.Join(path, "/")
			for c := 1; c < len(command); c++ {
				fields := strings.Fields(command[c])
				if fields[0] == "dir" {
					dirPath := absolutePath + "/" + fields[1]
					if subdirectories[absolutePath] != nil {
						subdirectories[absolutePath] = append(subdirectories[absolutePath], dirPath)
					} else {
						subdirectories[absolutePath] = []string{dirPath}
					}
				} else {
					atoi, _ := strconv.Atoi(fields[0])
					dirSizes[absolutePath] += atoi
				}
			}
		}
	}
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) > len(paths[j])
	})
	for s2 := range paths {
		for i := range subdirectories[paths[s2]] {
			dirSizes[paths[s2]] += dirSizes[subdirectories[paths[s2]][i]]
		}
	}
	atMost100_000 := 0
	for s2 := range dirSizes {
		if dirSizes[s2] < 100_000 {
			atMost100_000 += dirSizes[s2]
		}
	}
	smallest := dirSizes["/"]
	diskSize := dirSizes["/"]
	unusedSpace := 70000000 - diskSize
	for s2 := range dirSizes {
		if unusedSpace+dirSizes[s2] >= 30000000 && smallest > dirSizes[s2] {
			smallest = dirSizes[s2]
		}
	}
	return atMost100_000, smallest
}

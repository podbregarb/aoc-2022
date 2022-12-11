package main

import (
	_ "embed"
	"fmt"
	"regexp"
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
	split := strings.Split(s, "\n\n")
	return getMonkeyActivity(split, 20, true), getMonkeyActivity(split, 10000, false)
}

func getMonkeyActivity(split []string, rounds int, worried bool) int {
	monkeys := getMonkeys(split)
	lcm := 1
	for i := range monkeys {
		lcm *= monkeys[i].divisibleBy
	}
	doRounds(monkeys, rounds, worried, lcm)
	// Find most active monkeys
	monkeyActivity := make([]int, 0)
	for monkeyIndex := 0; monkeyIndex < len(monkeys); monkeyIndex++ {
		monkeyActivity = append(monkeyActivity, monkeys[monkeyIndex].numOfInspectedItems)
	}
	sort.Ints(monkeyActivity)
	return monkeyActivity[len(monkeyActivity)-1] * monkeyActivity[len(monkeyActivity)-2]
}

func doRounds(monkeys map[int]monkeyData, rounds int, worried bool, lcm int) {
	for i := 0; i < rounds; i++ {
		for monkeyIndex := 0; monkeyIndex < len(monkeys); monkeyIndex++ {
			monkey := monkeys[monkeyIndex]
			monkey.SetNumOfInspectedItems(monkey.NumOfInspectedItems() + len(monkeys[monkeyIndex].items))
			for itemIndex := range monkey.items {
				item := monkeys[monkeyIndex].items[itemIndex]
				worryLevel := monkey.operation(item)
				if worried {
					worryLevel /= 3
				} else {
					worryLevel = worryLevel % lcm
				}
				if worryLevel%monkey.divisibleBy == 0 {
					trueMonkey := monkeys[monkey.trueConditionMonkey]
					trueMonkey.SetItems(append(trueMonkey.items, worryLevel))
					monkeys[monkey.trueConditionMonkey] = trueMonkey
				} else {
					falseMonkey := monkeys[monkey.falseConditionMonkey]
					falseMonkey.SetItems(append(falseMonkey.items, worryLevel))
					monkeys[monkey.falseConditionMonkey] = falseMonkey
				}
			}
			monkey.SetItems(make([]int, 0))
			monkeys[monkeyIndex] = monkey
		}
	}
}

func getMonkeys(split []string) map[int]monkeyData {
	monkeys := make(map[int]monkeyData, 0)
	for oneMonkey := range split {
		oneMonkeyData := strings.Split(split[oneMonkey], "\n")
		monkeys[getMonkeyNum(oneMonkeyData)] = monkeyData{
			items:                getItems(oneMonkeyData),
			numOfInspectedItems:  0,
			divisibleBy:          getDivisibleBy(oneMonkeyData),
			trueConditionMonkey:  getTrueConditionMonkey(oneMonkeyData),
			falseConditionMonkey: getFalseConditionMonkey(oneMonkeyData),
			operation: func(old int) int {
				operationData := getOperationData(oneMonkeyData)
				firstInteger := old
				secondInteger := old
				if operationData[2] != "old" {
					second, _ := strconv.Atoi(operationData[2])
					secondInteger = second
				}
				if operationData[1] == "*" {
					return firstInteger * secondInteger

				} else {
					return firstInteger + secondInteger
				}
			},
		}
	}
	return monkeys
}

func getFalseConditionMonkey(oneMonkeyData []string) int {
	regex, _ := regexp.Compile("\\d+")
	matched := regex.FindString(oneMonkeyData[5])
	atoi, _ := strconv.Atoi(matched)
	return atoi
}

func getTrueConditionMonkey(oneMonkeyData []string) int {
	regex, _ := regexp.Compile("\\d+")
	matched := regex.FindString(oneMonkeyData[4])
	atoi, _ := strconv.Atoi(matched)
	return atoi
}

func getDivisibleBy(oneMonkeyData []string) int {
	regex, _ := regexp.Compile("\\d+")
	matched := regex.FindString(oneMonkeyData[3])
	atoi, _ := strconv.Atoi(matched)
	return atoi
}

func getOperationData(oneMonkeyData []string) []string {
	return strings.Fields(strings.Replace(oneMonkeyData[2], "  Operation: new = ", "", 1))
}

func getItems(oneMonkeyData []string) []int {
	regex, _ := regexp.Compile("\\d+")
	matched := regex.FindAllString(oneMonkeyData[1], -1)
	items := make([]int, 0)
	for i := range matched {
		item, _ := strconv.Atoi(matched[i])
		items = append(items, item)
	}
	return items
}

func getMonkeyNum(oneMonkeyData []string) int {
	regex, _ := regexp.Compile("\\d")
	matched := regex.FindString(oneMonkeyData[0])
	monkeyNum, _ := strconv.Atoi(matched)
	return monkeyNum
}

type Operation func(int) int

type monkeyData struct {
	items                []int
	numOfInspectedItems  int
	divisibleBy          int
	trueConditionMonkey  int
	falseConditionMonkey int

	operation Operation
}

func (m *monkeyData) Items() []int {
	return m.items
}

func (m *monkeyData) SetItems(items []int) {
	m.items = items
}

func (m *monkeyData) NumOfInspectedItems() int {
	return m.numOfInspectedItems
}

func (m *monkeyData) SetNumOfInspectedItems(numOfInspectedItems int) {
	m.numOfInspectedItems = numOfInspectedItems
}

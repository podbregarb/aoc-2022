package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func main() {
	fmt.Println("Test input")
	printAnswers(testInput, 10, 20)
	fmt.Println("Input")
	printAnswers(input, 2000000, 4000000)
}

func printAnswers(s string, row int, distressBeaconLocationLimit int) {
	first, second := getAnswers(s, row, distressBeaconLocationLimit)
	fmt.Printf("First answer: %d\n", first)
	fmt.Printf("Second answer: %d\n", second)
}

func getAnswers(s string, row int, distressBeaconLocationLimit int) (int, int) {
	sensorsAndBeacons := getSensorsAndBeacons(s)
	return getPositionsThatCanNotContainBeacon(sensorsAndBeacons, row), getDistressBeacon(sensorsAndBeacons, distressBeaconLocationLimit)
}

func getDistressBeacon(sensorsAndBeacons map[Point]Point, distressBeaconLocationLimit int) int {
	for y := 0; y <= distressBeaconLocationLimit; y++ {
	loop:
		for x := 0; x <= distressBeaconLocationLimit; x++ {
			for sensor, beacon := range sensorsAndBeacons {
				distance := getManhattanDistance(sensor, beacon)
				if getManhattanDistance(Point{x, y}, sensor) <= distance {
					x += distance - abs(sensor.y-y) + (sensor.x - x)
					continue loop
				}
			}
			return x*4000000 + y
		}
	}
	return -1
}

func getPositionsThatCanNotContainBeacon(sensorsAndBeacons map[Point]Point, row int) int {
	positionsThatCanNotContainBeacon := make(map[int]bool, 0)
	for sensor, beacon := range sensorsAndBeacons {
		d := getManhattanDistance(sensor, beacon) - abs(row-sensor.y)
		for i := sensor.x - d; i <= sensor.x+d; i++ {
			if !(i == beacon.x && row == beacon.y) {
				positionsThatCanNotContainBeacon[i] = true
			}
		}
	}
	return len(positionsThatCanNotContainBeacon)
}

func getManhattanDistance(p Point, r Point) int {
	return abs(p.x-r.x) + abs(p.y-r.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func getSensorsAndBeacons(s string) map[Point]Point {
	sensorsAndBeacons := make(map[Point]Point, 0)
	for _, str := range strings.Split(s, "\n") {
		split := strings.Split(str, ": ")
		sensorStr := strings.Split(split[0], ",")
		beaconStr := strings.Split(split[1], ",")
		sensorsAndBeacons[Point{getCoordinate(sensorStr[0]), getCoordinate(sensorStr[1])}] =
			Point{getCoordinate(beaconStr[0]), getCoordinate(beaconStr[1])}
	}
	return sensorsAndBeacons
}

func getCoordinate(str string) int {
	regex, _ := regexp.Compile("-?\\d+")
	matched := regex.FindString(str)
	coordinate, _ := strconv.Atoi(matched)
	return coordinate
}

type Point struct {
	x int
	y int
}

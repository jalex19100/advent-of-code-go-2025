package day3

import (
	"bufio"
	"strconv"
	"strings"
)

const Year = 2025
const Day = 3

// Part1 solves the first part of the puzzle.
func Part1(input string) int {
	return findMaxVoltage(input, 2)
}

// Part2 solves the second part of the puzzle.
func Part2(input string) int {
	return findMaxVoltage(input, 12)
}

func findMaxVoltage(input string, numberOfBatteriesNeeded int) int {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	var voltages []int
	for scanner.Scan() {
		bank := scanner.Text()
		largestBatteriesNeeded := findLargestBatteriesInOrder(bank, numberOfBatteriesNeeded)
		voltage, _ := strconv.Atoi(intsToString(largestBatteriesNeeded))
		voltages = append(voltages, voltage)
	}
	sum := 0
	for _, b := range voltages {
		sum += b
	}
	return sum
}

func findLargestBatteriesInOrder(batteryBank string, numberOfBatteriesNeeded int) []int {
	largestBatteries := make([]int, numberOfBatteriesNeeded)
	largestBatteryPosition := 0
	for b := 0; b < numberOfBatteriesNeeded; b++ {
		numberOfBatteriesStillLeft := numberOfBatteriesNeeded - b
		batteriesLeft := string(batteryBank[largestBatteryPosition : len(batteryBank)-(numberOfBatteriesStillLeft-1)])
		relativeBatteryPosition := findLargestBatteryPosition(batteriesLeft)
		largestBatteries[b], _ = strconv.Atoi(string(batteriesLeft[relativeBatteryPosition]))
		largestBatteryPosition = relativeBatteryPosition + (largestBatteryPosition + 1)
	}
	return largestBatteries
}

func findLargestBatteryPosition(batteries string) int {
	largestBattery := 0
	largestPosition := 0
	for i := 0; i < len(batteries); i++ {
		b, _ := strconv.Atoi(string(batteries[i]))
		if b > largestBattery {
			largestBattery = b
			largestPosition = i
		}
	}
	return largestPosition
}

func intsToString(numbers []int) string {
	var s string
	for _, n := range numbers {
		s += strconv.Itoa(n)
	}
	return s
}

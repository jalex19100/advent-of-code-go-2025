package day1

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const Year = 2025
const Day = 1

// Part1 solves the first part of the puzzle.
func Part1(input string) int {
	fmt.Println("Part 1")
	zeros, _ := findCombination(input)
	return zeros
}

// Part2 solves the second part of the puzzle.
func Part2(input string) int {
	fmt.Println("Part 2")
	_, passedZero := findCombination(input)
	return passedZero
}

func rotate(startingValue int, clicks int) (int, int) {
	rawNumber := startingValue + clicks
	absValue := int(math.Abs(float64(rawNumber)))
	remainder := absValue % 100
	clickQuotient := absValue / 100
	if clicks < 0 && rawNumber <= 0 && startingValue != 0 {
		clickQuotient++
	}
	if rawNumber < 0 && remainder != 0 {
		remainder = 100 - remainder
	}
	//fmt.Printf("Starting: %d, clicks: %d, remainder: %d, clickQuotient: %d\n", startingValue, clicks, remainder, clickQuotient)
	return remainder, clickQuotient
}

func findCombination(input string) (int, int) {
	combo := make([]int, strings.Count(input, "\n"))
	currentPosition := 50 // starting position is 50
	index := 0
	zeros := 0
	passedZero := 0
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text() // next line
		clicks, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			clicks = -clicks
		}
		quotient := 0
		combo[index], quotient = rotate(currentPosition, clicks)
		currentPosition = combo[index]
		if currentPosition == 0 {
			zeros++
		}
		passedZero += quotient
		//fmt.Println(line)
		index++
	}
	//fmt.Printf("Positions seen: %#v\\n", combo)
	return zeros, passedZero
}

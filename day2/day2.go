package day2

import (
	"bufio"
	"strconv"
	"strings"
)

const Year = 2025
const Day = 2

// Part1 solves the first part of the puzzle.
func Part1(input string) int {
	invalidIds := getInvalidIds(input, 2)
	var sumOfInvalidIds int
	for _, id := range invalidIds {
		sumOfInvalidIds += id
	}
	return int(sumOfInvalidIds)
}

// Part2 solves the second part of the puzzle.
func Part2(input string) int {
	invalidIds := getInvalidIds(input, 7)
	var sumOfInvalidIds int
	for _, id := range invalidIds {
		sumOfInvalidIds += id
	}
	return int(sumOfInvalidIds)
}

func getInvalidIds(input string, maxDivisions int) []int {
	var invalidIds []int
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		inputRanges := strings.Split(line, ",")
		for _, r := range inputRanges {
			inputRange := strings.Split(r, "-")
			start, _ := strconv.ParseInt(inputRange[0], 10, 64)
			end, _ := strconv.ParseInt(inputRange[1], 10, 64)
			for j := start; j <= end; j++ {
				testString := strconv.FormatInt(j, 10)
				if stringContainsMatchingSequences(testString, maxDivisions) {
					invalidIds = append(invalidIds, int(j))
				}
			}
		}
	}
	return invalidIds
}

func stringContainsMatchingSequences(s string, maxDivisions int) bool {
	allMatches := false
	for splitCount := 2; splitCount <= maxDivisions; splitCount++ {
		if len(s)%splitCount == 0 {
			seqLen := len(s) / splitCount
			sequences := make([]string, splitCount)
			for i := 0; i < splitCount; i++ {
				start := i * seqLen
				end := start + seqLen
				sequences[i] = s[start:end]
			}
			if allStringsMatch(sequences) {
				return true
			}
		}
	}
	return allMatches
}

func allStringsMatch(strings []string) bool {
	if len(strings) == 0 {
		return false
	}
	first := strings[0]
	for _, part := range strings[1:] {
		if part != first {
			return false
		}
	}
	return true
}

package day4

import "strings"

const Year = 2025
const Day = 4
const Roll = '@'

// Part1 solves the first part of the puzzle.
func Part1(input string) int {
	g := buildGrid(input)
	return len(g.getAccessibleRolls())
}

// Part2 solves the second part of the puzzle.
func Part2(input string) int {
	g := buildGrid(input)
	var rollsRemoved int
	moreRollsCanBeRemoved := true
	for moreRollsCanBeRemoved {
		accessibleRolls := g.getAccessibleRolls()
		if len(accessibleRolls) == 0 {
			moreRollsCanBeRemoved = false
		} else {
			for _, c := range accessibleRolls {
				g.Set(c, '.')
				rollsRemoved++
			}
		}
	}
	return rollsRemoved
}

type Coordinate struct {
	x int
	y int
}

var allDirections = []Coordinate{
	{-1, 0}, {-1, -1}, //  Left, Upper Left
	{0, -1}, {1, -1}, // Up, Upper Right
	{1, 0}, {1, 1}, // Right, Lower Right
	{0, 1}, {-1, 1}, // Down, Lower Left
}

type Grid map[Coordinate]rune

func NewGrid() Grid {
	return make(Grid)
}

func (g Grid) Set(c Coordinate, value rune) {
	g[Coordinate{c.x, c.y}] = value
}

func (g Grid) Get(c Coordinate) rune {
	if val, ok := g[c]; ok {
		return val
	}
	return '.'
}

func buildGrid(input string) Grid {
	g := NewGrid()
	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, value := range line {
			g.Set(Coordinate{x, y}, value)
		}
	}
	return g
}

func (g Grid) countAdjacentPositions(c Coordinate, v rune) int {
	count := 0
	for _, d := range allDirections {
		if g.Get(Coordinate{c.x + d.x, c.y + d.y}) == v {
			count++
		}
	}
	return count
}

func (g Grid) getAccessibleRolls() []Coordinate {

	var accessibleRolls []Coordinate
	// randomly iterate over the grid
	for c, v := range g {
		if v == Roll {
			adjacentRolls := g.countAdjacentPositions(c, Roll)
			if adjacentRolls < 4 {
				accessibleRolls = append(accessibleRolls, c)
			}
		}
	}
	return accessibleRolls
}

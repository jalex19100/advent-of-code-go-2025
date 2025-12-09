package day8

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const Year = 2025
const Day = 8

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type junctionBox struct {
	x, y, z int
}

func inputStringsToJunctionBoxes(input string) []junctionBox {
	positions := make([]junctionBox, 0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(strings.TrimSpace(line), ",")
		junctionBox := junctionBox{x: toInt(parts[0]), y: toInt(parts[1]), z: toInt(parts[2])}
		positions = append(positions, junctionBox)
	}
	return positions
}

type circuit struct {
	boxes []junctionBox
}

func (c circuit) contains(j junctionBox) bool {
	return slices.Contains(c.boxes, j)
}

func (c *circuit) add(j junctionBox) {
	c.boxes = append(c.boxes, j)
}

func joinCircuits(c1 circuit, c2 circuit) circuit {
	return circuit{boxes: append(c1.boxes, c2.boxes...)}
}

func addPairToCircuits(circuits []circuit, a junctionBox, b junctionBox) []circuit {
	aIndex := slices.IndexFunc(circuits, func(c circuit) bool { return c.contains(a) })
	bIndex := slices.IndexFunc(circuits, func(c circuit) bool { return c.contains(b) })
	if aIndex == -1 && bIndex == -1 { // a and b are new
		// neither junction box is in any circuit
		circuits = append(circuits, circuit{boxes: []junctionBox{a, b}})
	} else if aIndex != -1 && bIndex == -1 { // b is new
		circuits[aIndex].add(b)
	} else if aIndex == -1 { // a is new, assumes bIndex != -1
		circuits[bIndex].add(a)
	} else if aIndex != bIndex { // a and b are not new but in different circuits, assumes aIndex != -1 && bIndex != -1
		circuits[aIndex] = joinCircuits(circuits[aIndex], circuits[bIndex])
		circuits = append(circuits[:bIndex], circuits[bIndex+1:]...)
	}
	return circuits
}

func compare(a circuit, b circuit) int {
	if len(a.boxes) < len(b.boxes) {
		return -1
	}
	if len(a.boxes) > len(b.boxes) {
		return 1
	}
	return 0
}

func (j junctionBox) remove(positions []junctionBox) []junctionBox {
	idx := slices.IndexFunc(positions, func(c junctionBox) bool { return c.equal(j) })
	if idx == -1 {
		return positions
	}
	return append(positions[:idx], positions[idx+1:]...)
}

func (j junctionBox) equal(other junctionBox) bool {
	return j.x == other.x && j.y == other.y && j.z == other.z
}

func (j junctionBox) distanceApart(other junctionBox) float64 {
	return math.Sqrt(float64((j.x-other.x)*(j.x-other.x) + (j.y-other.y)*(j.y-other.y) + (j.z-other.z)*(j.z-other.z)))
}

func (j junctionBox) closestJunctionBox(positions []junctionBox) junctionBox {
	closestJunctionBox := slices.MinFunc(
		positions,
		func(a, b junctionBox) int {
			distA := math.MaxFloat64
			if !a.equal(j) {
				distA = j.distanceApart(a)
			}

			distB := math.MaxFloat64
			if !b.equal(j) {
				distB = j.distanceApart(b)
			}

			if distA < distB {
				return -1
			}
			if distA > distB {
				return 1
			}
			return 0
		},
	)
	return closestJunctionBox
}

type pairDistance struct {
	pair     [2]junctionBox
	distance int
}

func (p1 pairDistance) equal(p2 pairDistance) bool {
	return p1.distance == p2.distance && p1.pair[0].equal(p2.pair[0]) && p1.pair[1].equal(p2.pair[1])
}

func getUniquePairsSortedByDistance(boxes []junctionBox) []pairDistance {
	var results []pairDistance
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			p1, p2 := boxes[i], boxes[j]
			dist := p1.distanceApart(p2)
			results = append(results, pairDistance{
				pair:     [2]junctionBox{p1, p2},
				distance: int(dist),
			})
		}
	}

	slices.SortFunc(results, func(a, b pairDistance) int {
		return a.distance - b.distance
	})

	return results
}

// Part1 solves the first part of the puzzle.
func Part1(input string, limit int) int {
	junctionBoxes := inputStringsToJunctionBoxes(input)
	pairs := getUniquePairsSortedByDistance(junctionBoxes)
	circuits := make([]circuit, 0)
	for _, pair := range pairs[:limit] {
		circuits = addPairToCircuits(circuits, pair.pair[0], pair.pair[1])
	}

	slices.SortFunc(circuits, func(a, b circuit) int {
		return len(a.boxes) - len(b.boxes)
	})
	slices.Reverse(circuits)

	result := 1
	for _, circuit := range circuits[:3] {
		result *= len(circuit.boxes)
	}
	return result
}

// Part2 solves the second part of the puzzle.
func Part2(input string, limit int) int {
	fmt.Printf("Limit ignored: %d\n", limit)
	junctionBoxes := inputStringsToJunctionBoxes(input)
	pairs := getUniquePairsSortedByDistance(junctionBoxes)
	circuits := make([]circuit, 0)
	for _, pair := range pairs {
		circuits = addPairToCircuits(circuits, pair.pair[0], pair.pair[1])
		if len(circuits) == 1 && len(circuits[0].boxes) == len(junctionBoxes) {
			fmt.Printf("Last pair added: %v\n", pair)
			return pair.pair[0].x * pair.pair[1].x
		}
	}
	return 0
}

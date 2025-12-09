package day8

import (
	"os"
	"testing"
)

// Expected values for each scenario. Update these as you implement solutions.
// Defaults are 0 so the scaffold compiles and runs immediately.
var (
	expectedSamplePart1 = 40
	expectedSamplePart2 = 25272
	expectedInputPart1  = 244188
	expectedInputPart2  = 1
)

// readFile is a small helper to read a file or fail the test.
func readFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed reading %s: %v", path, err)
	}
	return string(b)
}

// TestDay8 runs both parts against sample_input and input.
// Each scenario has its own expected value so you can update them independently.
func TestDay8(t *testing.T) {
	sampleInput := readFile(t, "sample_input")
	puzzleInput := readFile(t, "input")

	tests := []struct {
		name  string
		input string
		limit int
		part  string
		got   func(string, int) int
		want  int
	}{
		{name: "sample_input/part1", input: sampleInput, limit: 10, part: "Part1", got: Part1, want: expectedSamplePart1},
		{name: "sample_input/part2", input: sampleInput, limit: 10, part: "Part2", got: Part2, want: expectedSamplePart2},
		{name: "input/part1", input: puzzleInput, limit: 1000, part: "Part1", got: Part1, want: expectedInputPart1},
		{name: "input/part2", input: puzzleInput, limit: 1000, part: "Part2", got: Part2, want: expectedInputPart2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.got(tt.input, tt.limit)
			if got != tt.want {
				t.Errorf("%s: got %d; want %d", tt.part, got, tt.want)
			}
		})
	}
}

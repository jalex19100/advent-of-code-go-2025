package main

import (
    _ "embed"
    "fmt"
    "os"
    "os/exec"
    "strconv"
    "strings"

    "advent-of-code-go-2025/internal/inputdl"
    "advent-of-code-go-2025/internal/scaffold"
)

var (
	version = "0.6.0"
)

// Embed the day source template so the binary can generate files without external deps.
//
//go:embed templates/day.go.tmpl
var dayTemplate string

//go:embed templates/day_test.go.tmpl
var dayTestTemplate string

func main() {
	// Expect a single positional argument: the day (1-25)
	if len(os.Args) != 2 {
		usageAndExit()
	}

	d, err := strconv.Atoi(os.Args[1])
	if err != nil || d < 1 || d > 25 {
		usageAndExit()
	}

 const year = 2025
 fmt.Printf("Advent of Code %d Day %d\n", year, d)

 // Delegate scaffolding to internal/scaffold
 // Read session from environment (preferred), or from a file path provided
 // via AOC_SESSION_FILE. The file, if set, will be added to .gitignore.
 session := strings.TrimSpace(os.Getenv("AOC_SESSION"))
 sessionFile := strings.TrimSpace(os.Getenv("AOC_SESSION_FILE"))
 if session == "" && sessionFile != "" {
     if b, err := os.ReadFile(sessionFile); err == nil {
         session = strings.TrimSpace(string(b))
     } else {
         // Non-fatal: continue without session
         fmt.Fprintf(os.Stderr, "Warning: could not read AOC_SESSION_FILE %q: %v\n", sessionFile, err)
     }
 }

 opts := scaffold.Options{Day: d, Year: year, Template: dayTemplate, TestTemplate: dayTestTemplate, SessionFile: sessionFile}
    dir, createdSrc, createdTest, _, _, err := scaffold.Run(opts)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error scaffolding day %d: %v\n", d, err)
        os.Exit(1)
    }

    // Try to populate the input file if it's empty.
    if err := inputdl.TryPopulateInput(year, d, dir, session); err != nil {
        // Non-fatal: warn and continue.
        fmt.Fprintf(os.Stderr, "Warning: could not download input: %v\n", err)
    }

	//abs, _ := filepath.Abs(dir)
	//fmt.Printf("Directory: %s\n", abs)

	// If both the source and test already existed (i.e., were not created now),
	// run the tests for that day.
	if !createdSrc && !createdTest {
		if err := runTests(dir); err != nil {
			fmt.Fprintf(os.Stderr, "Tests failed for %s: %v\n", dir, err)
			os.Exit(1)
		}
	}
}

func usageAndExit() {
	fmt.Fprintf(os.Stderr, "advent-of-code %s\n", version)
	fmt.Fprintln(os.Stderr, "Usage: aoc <day>\n  Example: aoc 1")
	os.Exit(2)
}

// runTests executes `go test` for the specified day directory (e.g., ./day1).
func runTests(dir string) error {
    cmd := exec.Command("go", "test", "./"+dir)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Env = os.Environ()
    return cmd.Run()
}

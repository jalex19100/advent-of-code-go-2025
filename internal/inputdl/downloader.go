package inputdl

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// TryPopulateInput checks dayN/input and if it's empty, attempts to download
// the real puzzle input from Advent of Code using the provided session token.
// If session is empty, it prints a warning to stderr and returns nil, leaving
// the input file empty.
func TryPopulateInput(year, day int, dir, session string) error {
	inputPath := filepath.Join(dir, "input")

	fi, err := os.Stat(inputPath)
	if err != nil {
		// If the file does not exist, do nothing (scaffold should have created it)
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	// Only act if the file is empty
	if fi.Size() > 0 {
		return nil
	}

	if session == "" {
		fmt.Fprintf(os.Stderr, "Warning: input is empty and no AOC_SESSION provided. The input file will remain empty. Set AOC_SESSION or AOC_SESSION_FILE and re-run.\n")
		return nil
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", session))
	req.Header.Set("User-Agent", "advent-of-code-go-2025 aoc-cli (+https://github.com/)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status %d downloading AoC input", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := os.WriteFile(inputPath, data, 0o644); err != nil {
		return err
	}
	// Inform the user that the input was successfully downloaded.
	fmt.Printf("Updating input from AOC: %s (%d bytes)\n", inputPath, len(data))
	return nil
}

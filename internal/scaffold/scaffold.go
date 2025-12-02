package scaffold

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Options contains the parameters for scaffolding a day.
type Options struct {
	Day          int
	Year         int
	Template     string
	TestTemplate string
	// SessionFile is an optional path to a file that contains the session
	// token. If provided, we'll try to add it to .gitignore to avoid leaks.
	SessionFile string
}

// Run creates the directory and day source file for a given day.
// It returns the directory path, a boolean indicating if the source file was created, and an error.
func Run(opts Options) (string, bool, bool, bool, bool, error) {
	dir := fmt.Sprintf("day%d", opts.Day)
	// Create the day directory if it does not exist
	if err := ensureDir(dir); err != nil {
		return "", false, false, false, false, fmt.Errorf("ensuring directory %q: %w", dir, err)
	}

	// Ensure the day source file exists (do not overwrite if already present)
	createdSrc, err := ensureDaySource(opts, dir)
	if err != nil {
		return "", false, false, false, false, fmt.Errorf("ensuring day source for %q: %w", dir, err)
	}
	if createdSrc {
		fmt.Printf("Created code from template: %s\n", filepath.Join(dir, fmt.Sprintf("day%d.go", opts.Day)))
	}

	// Ensure the day test file exists (do not overwrite if already present)
	createdTest, err := ensureDayTest(opts, dir)
	if err != nil {
		return "", false, false, false, false, fmt.Errorf("ensuring day test for %q: %w", dir, err)
	}
	if createdTest {
		fmt.Printf("Created test from template: %s\n", filepath.Join(dir, fmt.Sprintf("day%d_test.go", opts.Day)))
	}

	// Ensure input files exist (sample_input and input)
	createdSample, createdInput, err := ensureInputFiles(dir)
	if err != nil {
		return "", false, false, false, false, fmt.Errorf("ensuring input files for %q: %w", dir, err)
	}
	if createdSample {
		fmt.Printf("Created sample input file: %s\n", filepath.Join(dir, "sample_input"))
	}
	if createdInput {
		fmt.Printf("Created input file: %s\n", filepath.Join(dir, "input"))
	}

	// If a session file path was provided, ensure it is ignored by git.
	if opts.SessionFile != "" {
		if err := ensureGitignoreLines([]string{opts.SessionFile}); err != nil {
			return "", false, false, false, false, fmt.Errorf("updating .gitignore: %w", err)
		}
	}

	return dir, createdSrc, createdTest, createdSample, createdInput, nil
}

func ensureDir(path string) error {
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return nil
		}
		return fmt.Errorf("%q exists and is not a directory", path)
	}
	if os.IsNotExist(err) {
		return os.MkdirAll(path, 0o755)
	}
	return err
}

// ensureDaySource creates a dayN/dayN.go file from template if it does not already exist.
// Returns true if the file was created.
func ensureDaySource(opts Options, dir string) (bool, error) {
	filename := fmt.Sprintf("day%d.go", opts.Day)
	full := filepath.Join(dir, filename)

	// If file exists, do nothing
	if _, err := os.Stat(full); err == nil {
		return false, nil
	} else if !os.IsNotExist(err) {
		return false, err
	}

	src, err := renderDayTemplate(opts)
	if err != nil {
		return false, fmt.Errorf("rendering template: %w", err)
	}

	if err := os.WriteFile(full, []byte(src), 0o644); err != nil {
		return false, fmt.Errorf("writing file: %w", err)
	}
	return true, nil
}

// ensureDayTest creates a dayN/dayN_test.go file from template if it does not already exist.
// Returns true if the file was created.
func ensureDayTest(opts Options, dir string) (bool, error) {
	filename := fmt.Sprintf("day%d_test.go", opts.Day)
	full := filepath.Join(dir, filename)

	if _, err := os.Stat(full); err == nil {
		return false, nil
	} else if !os.IsNotExist(err) {
		return false, err
	}

	if opts.TestTemplate == "" {
		// No test template provided; do not create the file.
		return false, nil
	}

	src, err := renderTestTemplate(opts)
	if err != nil {
		return false, fmt.Errorf("rendering test template: %w", err)
	}

	if err := os.WriteFile(full, []byte(src), 0o644); err != nil {
		return false, fmt.Errorf("writing test file: %w", err)
	}
	return true, nil
}

type dayTemplateData struct {
	Package string
	Year    int
	Day     int
}

func renderDayTemplate(opts Options) (string, error) {
	data := dayTemplateData{
		Package: fmt.Sprintf("day%d", opts.Day),
		Year:    opts.Year,
		Day:     opts.Day,
	}
	t, err := template.New("day").Parse(opts.Template)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, &data); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}
	return buf.String(), nil
}

// ensureInputFiles checks for and creates empty input files used by the puzzles.
// It ensures two files under the day's directory:
//   - sample_input
//   - input
//
// Returns booleans indicating whether each file was created.
func ensureInputFiles(dir string) (bool, bool, error) {
	samplePath := filepath.Join(dir, "sample_input")
	inputPath := filepath.Join(dir, "input")

	createdSample, err := ensureEmptyFile(samplePath)
	if err != nil {
		return false, false, err
	}

	createdInput, err := ensureEmptyFile(inputPath)
	if err != nil {
		return false, false, err
	}

	return createdSample, createdInput, nil
}

// ensureEmptyFile creates an empty file if it does not already exist.
// It does not overwrite existing files.
func ensureEmptyFile(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return false, nil
	} else if !os.IsNotExist(err) {
		return false, err
	}
	if err := os.WriteFile(path, []byte(""), 0o644); err != nil {
		return false, fmt.Errorf("writing %q: %w", path, err)
	}
	return true, nil
}

func renderTestTemplate(opts Options) (string, error) {
	data := dayTemplateData{
		Package: fmt.Sprintf("day%d", opts.Day),
		Year:    opts.Year,
		Day:     opts.Day,
	}
	t, err := template.New("day_test").Parse(opts.TestTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing test template: %w", err)
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, &data); err != nil {
		return "", fmt.Errorf("executing test template: %w", err)
	}
	return buf.String(), nil
}

// ensureGitignoreLines appends the provided lines to .gitignore if they are
// not already present. It is idempotent.
func ensureGitignoreLines(lines []string) error {
	const fname = ".gitignore"
	var existing string
	if b, err := os.ReadFile(fname); err == nil {
		existing = string(b)
	} else if !os.IsNotExist(err) {
		return err
	}

	add := make([]string, 0, len(lines))
	for _, ln := range lines {
		if ln == "" {
			continue
		}
		// Normalize to forward slashes in gitignore entries
		norm := filepath.ToSlash(ln)
		if !containsLine(existing, norm) {
			add = append(add, norm)
		}
	}
	if len(add) == 0 {
		return nil
	}
	// Ensure file ends with newline
	if existing != "" && existing[len(existing)-1] != '\n' {
		existing += "\n"
	}
	updated := existing + strings.Join(add, "\n") + "\n"
	return os.WriteFile(fname, []byte(updated), 0o644)
}

func containsLine(content, line string) bool {
	for _, l := range strings.Split(content, "\n") {
		if strings.TrimSpace(l) == strings.TrimSpace(line) {
			return true
		}
	}
	return false
}

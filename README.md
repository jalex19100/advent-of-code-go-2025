# advent-of-code-go-2025
Advent of Code 2025 with Go

[Advent of Code](https://adventofcode.com/) is an Advent calendar of small programming puzzles for a variety of skill levels that can be solved in any programming language you like.

The puzzles start at midnight on December 1st so that the day numbers make sense (Day 1 = Dec 1), and puzzles come out every day (ending mid-December).

The URL for each day is in the format of https://adventofcode.com/2025/day/1, where the year is 2025 and the day is 1.


## CLI usage

Build the CLI:

```
## Using Make
make
## Alternatively, run go build directly.
go build -o aoc ./...

```

Run for a specific day using a single positional argument:

```
# Execute binary
./aoc 1
# or during development
go run ./cmd/aoc 1
# Use Makefile
make run DAY=1
```

### What the CLI does
On first run for a given day `N`:
- Ensures a `dayN/` directory exists in the current working directory (creates it if missing).
- Creates a starter Go file `dayN/dayN.go` from a template with stub functions `Part1` and `Part2` that each return `0`.
- Creates a basic test file `dayN/dayN_test.go` that asserts both parts return `0` for empty input and that their results are equal.
- Creates two empty files under the day directory:
  - `dayN/sample_input`
  - `dayN/input` (created empty by scaffold)
- After scaffolding, if `dayN/input` is still empty and an Advent of Code session is available, the CLI will download the real puzzle input and write it to `dayN/input`.
- The templates are embedded in the binary (`cmd/aoc/templates/day.go.tmpl`, `cmd/aoc/templates/day_test.go.tmpl`), so no extra files are needed at runtime.

Subsequent runs will not overwrite existing files. If both the source and test files already exist for the day, the CLI will automatically run `go test ./dayN`.

### Supplying your Advent of Code session (for input download)
The puzzle input endpoint requires your personal AoC `session` cookie. Provide it at runtime via environment variables:

- Preferred: set `AOC_SESSION` to the session token value.
- Or set `AOC_SESSION_FILE` to a file path containing the token (the file’s contents will be read, and the path is added to `.gitignore`).

Examples:

```
# Direct env var
export AOC_SESSION="<your-session-token>"
./aoc 1

# From a file
echo -n "<your-session-token>" > .aoc_session
export AOC_SESSION_FILE=".aoc_session"
./aoc 1
```

Behavior:
- If `dayN/input` is empty and a session is provided, the CLI downloads the input and prints a confirmation like:
  `Updating input from AOC: dayN/input (<bytes> bytes)`
- If no session is available, a warning is printed and the input file remains empty. You can set the session and re-run the command.

### Git ignore and secrets hygiene
- Downloaded inputs are ignored by default via `day*/input` in `.gitignore`.
- If you use `AOC_SESSION_FILE`, its path is appended to `.gitignore` (idempotent) to avoid committing secrets.

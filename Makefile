SHELL := /bin/bash

.PHONY: all build run clean

# Default target builds the CLI binary at project root named "aoc"
all: build

build:
	GO111MODULE=on go build -o aoc ./cmd/aoc

# Example: make run DAY=1
run: build
	@if [ -z "$(DAY)" ]; then \
		echo "Usage: make run DAY=<1-25>" >&2; \
		exit 2; \
	fi
	./aoc $(DAY)

clean:
	rm -f aoc
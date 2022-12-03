package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	score := 0
	for _, line := range lines {
		elf := line[0:1]
		us := line[2:3]
		if elf == "A" && us == "X" {
			score += 1 + 3
		} else if elf == "A" && us == "Y" {
			score += 2 + 6
		} else if elf == "A" && us == "Z" {
			score += 3 + 0
		} else if elf == "B" && us == "X" {
			score += 1 + 0
		} else if elf == "B" && us == "Y" {
			score += 2 + 3
		} else if elf == "B" && us == "Z" {
			score += 3 + 6
		} else if elf == "C" && us == "X" {
			score += 1 + 6
		} else if elf == "C" && us == "Y" {
			score += 2 + 0
		} else if elf == "C" && us == "Z" {
			score += 3 + 3
		}
	}
	fmt.Printf("Score part1: %d\n", score)
}

func part2(lines []string) {
	score := 0
	for _, line := range lines {
		elf := line[0:1]
		us := line[2:3]
		if elf == "A" && us == "X" {
			score += 3 + 0
		} else if elf == "A" && us == "Y" {
			score += 1 + 3
		} else if elf == "A" && us == "Z" {
			score += 2 + 6
		} else if elf == "B" && us == "X" {
			score += 1 + 0
		} else if elf == "B" && us == "Y" {
			score += 2 + 3
		} else if elf == "B" && us == "Z" {
			score += 3 + 6
		} else if elf == "C" && us == "X" {
			score += 2 + 0
		} else if elf == "C" && us == "Y" {
			score += 3 + 3
		} else if elf == "C" && us == "Z" {
			score += 1 + 6
		}
	}
	fmt.Printf("Score part2: %d\n", score)
}

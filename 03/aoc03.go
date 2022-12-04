package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./input3.txt")
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
		b1 := line[0 : len(line)/2]
		b2 := line[len(line)/2 : len(line)]
		match := ' '
		for _, c1 := range b1 {
			for _, c2 := range b2 {
				if c1 == c2 {
					match = c1
					break
				}
			}
			if match != ' ' {
				break
			}
		}
		if int(match) > 96 {
			score += int(match) - 96
			//fmt.Printf(" -- %s -- %s --  %c -- %d\n", b1, b2, match, int(match)-96)
		} else {
			score += int(match) - 64 + 26
			//fmt.Printf(" -- %s -- %s --  %c -- %d\n", b1, b2, match, int(match)-64+26)
		}
	}
	fmt.Printf("Score part1: %d\n", score)
}

func part2(lines []string) {
	score := 0
	for i := 0; i < len(lines); i += 3 {
		match := ' '
		for _, c1 := range lines[i] {
			for _, c2 := range lines[i+1] {
				for _, c3 := range lines[i+2] {
					if c1 == c2 && c2 == c3 {
						match = c1
						break
					}
				}
				if match != ' ' {
					break
				}
			}
			if match != ' ' {
				break
			}
		}
		if int(match) > 96 {
			score += int(match) - 96
			//fmt.Printf(" -- %s -- %s --  %c -- %d\n", b1, b2, match, int(match)-96)
		} else {
			score += int(match) - 64 + 26
			//fmt.Printf(" -- %s -- %s --  %c -- %d\n", b1, b2, match, int(match)-64+26)
		}
	}
	fmt.Printf("Score part2: %d\n", score)
}

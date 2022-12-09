package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./input8.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	sum := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if isVisible(lines, j, i) {
				sum++
			}
		}
	}
	fmt.Printf("Part 1, visible sum: %d\n", sum)
}

func part2(lines []string) {
	best := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			score := sightScore(lines, j, i)
			if score > best {
				best = score
			}
		}
	}
	fmt.Printf("Part 2, best score: %d\n", best)
}

func isVisible(lines []string, x int, y int) bool {
	if x == 0 ||
		y == 0 ||
		x == len(lines[0])-1 ||
		y == len(lines)-1 {
		return true
	}

	for i := 0; i <= x; i++ {
		if i == x {
			return true
		}
		if lines[y][i] >= lines[y][x] {
			break
		}
	}

	for i := len(lines[y]) - 1; i >= x; i-- {
		if i == x {
			return true
		}
		if lines[y][i] >= lines[y][x] {
			break
		}
	}

	for i := 0; i <= y; i++ {
		if i == y {
			return true
		}
		if lines[i][x] >= lines[y][x] {
			break
		}
	}

	for i := len(lines) - 1; i >= y; i-- {
		if i == y {
			return true
		}
		if lines[i][x] >= lines[y][x] {
			break
		}
	}

	return false
}

func sightScore(lines []string, x int, y int) int {
	if x == 0 ||
		y == 0 ||
		x == len(lines[0])-1 ||
		y == len(lines)-1 {
		return 0
	}

	scorex1 := 0
	for i := x + 1; i < len(lines[y]); i++ {
		scorex1++
		if lines[y][i] >= lines[y][x] {
			break
		}
	}

	scorex2 := 0
	for i := x - 1; i >= 0; i-- {
		scorex2++
		if lines[y][i] >= lines[y][x] {
			break
		}
	}

	scorey1 := 0
	for i := y + 1; i < len(lines); i++ {
		scorey1++
		if lines[i][x] >= lines[y][x] {
			break
		}
	}

	scorey2 := 0
	for i := y - 1; i >= 0; i-- {
		scorey2++
		if lines[i][x] >= lines[y][x] {
			break
		}
	}

	return scorex1 * scorex2 * scorey1 * scorey2
}

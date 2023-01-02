package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./input10.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	var instr []int
	for _, line := range lines {
		if line == "noop" {
			instr = append(instr, 0)
		} else {
			var i int
			n, err := fmt.Sscanf(line, "addx %d", &i)
			if n == 1 && err == nil {
				instr = append(instr, i)
			}
		}
	}
	part1(instr)
	part2(instr)
}

func part1(instr []int) {
	fmt.Println(instr)
	curcycle := 0
	x := 1
	signal := 0
	for _, i := range instr {
		if i == 0 {
			curcycle++
		} else {
			curcycle += 2
			x += i
		}
		if (curcycle-20)%40 == 0 {
			fmt.Printf("Cycle: %d, x: %d, i: %d\n", curcycle, x-i, i)
			signal += (x - i) * curcycle
		} else if ((curcycle-20)%40 == 1) && i != 0 {
			fmt.Printf("Cycle: %d, x: %d\n", curcycle-1, x-i)
			signal += (x - i) * (curcycle - 1)
		}

	}
	fmt.Printf("Signal part1: %d\n", signal)
}

func part2(instr []int) {
	fmt.Printf("Signal part2: %d\n", 0)
}

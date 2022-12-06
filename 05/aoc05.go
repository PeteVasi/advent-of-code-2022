package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type mov struct {
	count int
	from  int
	to    int
}

func Push(s string, v string) string {
	return s + v
}

func Pop(s string) (string, string) {
	l := len(s)
	return s[0 : l-1], s[l-1 : l]
}

func main() {
	/*
		    [D]
		[N] [C]
		[Z] [M] [P]
		 1   2   3
	*/
	var sampleStacks []string
	sampleStacks = append(sampleStacks, "ZN")
	sampleStacks = append(sampleStacks, "MCD")
	sampleStacks = append(sampleStacks, "P")

	/*
		[F]         [L]     [M]
		[T]     [H] [V] [G] [V]
		[N]     [T] [D] [R] [N]     [D]
		[Z]     [B] [C] [P] [B] [R] [Z]
		[M]     [J] [N] [M] [F] [M] [V] [H]
		[G] [J] [L] [J] [S] [C] [G] [M] [F]
		[H] [W] [V] [P] [W] [H] [H] [N] [N]
		[J] [V] [G] [B] [F] [G] [D] [H] [G]
		 1   2   3   4   5   6   7   8   9
	*/
	var stacks []string
	stacks = append(stacks, "JHGMZNTF")
	stacks = append(stacks, "VWJ")
	stacks = append(stacks, "GVLJBTH")
	stacks = append(stacks, "BPJNCDVL")
	stacks = append(stacks, "FWSMPRG")
	stacks = append(stacks, "GHCFBNVM")
	stacks = append(stacks, "DHGMR")
	stacks = append(stacks, "HNMVZD")
	stacks = append(stacks, "GNFH")

	raw, err := os.ReadFile("./input5.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	var moves []mov
	for _, line := range lines {
		var m mov
		n, err := fmt.Sscanf(line, "move %d from %d to %d", &m.count, &m.from, &m.to)
		if n == 3 && err == nil {
			moves = append(moves, m)
		}
	}
	//fmt.Println(moves)
	//part1(stacks, moves)
	part2(stacks, moves)
}

func part1(stacks []string, moves []mov) {
	var v string
	for _, m := range moves {
		for i := 0; i < m.count; i++ {
			stacks[m.from-1], v = Pop(stacks[m.from-1])
			stacks[m.to-1] = Push(stacks[m.to-1], v)
		}
	}
	fmt.Println("Stacks part1:")
	fmt.Println(stacks)
}

func part2(stacks []string, moves []mov) {
	var v string
	tmp := ""
	for _, m := range moves {
		for i := 0; i < m.count; i++ {
			stacks[m.from-1], v = Pop(stacks[m.from-1])
			tmp = Push(tmp, v)
		}
		for i := 0; i < m.count; i++ {
			tmp, v = Pop(tmp)
			stacks[m.to-1] = Push(stacks[m.to-1], v)
		}
	}
	fmt.Println("Stacks part2:")
	fmt.Println(stacks)
}

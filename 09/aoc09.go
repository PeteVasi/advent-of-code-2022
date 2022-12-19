package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type mov struct {
	dir  string
	dist int
}

func main() {
	raw, err := os.ReadFile("./input9.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	var moves []mov
	for _, line := range lines {
		var m mov
		n, err := fmt.Sscanf(line, "%s %d", &m.dir, &m.dist)
		if n == 2 && err == nil {
			moves = append(moves, m)
		}
	}
	part1(moves)
	part2(moves)
}

func part1(moves []mov) {
	var board [900][900]int
	hx := 450
	hy := 450
	tx := 450
	ty := 450
	board[ty][tx] = 1
	for _, m := range moves {
		//fmt.Printf("at: %d, %d  (%d, %d)\n", hx, hy, tx, ty)
		for i := 0; i < m.dist; i++ {
			if m.dir == "R" {
				hx++
				if tx < hx-1 {
					tx++
					ty = hy
				}
			} else if m.dir == "U" {
				hy++
				if ty < hy-1 {
					ty++
					tx = hx
				}
			} else if m.dir == "L" {
				hx--
				if tx > hx+1 {
					tx--
					ty = hy
				}
			} else if m.dir == "D" {
				hy--
				if ty > hy+1 {
					ty--
					tx = hx
				}
			}
			board[ty][tx] = 1
		}
	}
	visits := 0
	for _, by := range board {
		for _, bx := range by {
			if bx == 1 {
				visits++
			}
		}
	}
	fmt.Printf("Visits part1: %d\n", visits)
}

func part2(moves []mov) {
	var board [900][900]int
	var tx [10]int
	var ty [10]int
	for i := 0; i < 10; i++ {
		tx[i] = 450
		ty[i] = 450
	}
	board[ty[9]][tx[9]] = 1
	for _, m := range moves {
		for i := 0; i < m.dist; i++ {
			if m.dir == "R" {
				tx[0]++
				for t := 1; t < 10; t++ {
					if tx[t] < tx[t-1]-1 {
						tx[t]++
						ty[t] = ty[t-1]
					}
				}
			} else if m.dir == "U" {
				ty[0]++
				for t := 1; t < 10; t++ {
					if ty[t] < ty[t-1]-1 {
						ty[t]++
						tx[t] = tx[t-1]
					}
				}
			} else if m.dir == "L" {
				tx[0]--
				for t := 1; t < 10; t++ {
					if tx[t] > tx[t-1]+1 {
						tx[t]--
						ty[t] = ty[t-1]
					}
				}
			} else if m.dir == "D" {
				ty[0]--
				for t := 1; t < 10; t++ {
					if ty[t] > ty[t-1]+1 {
						ty[t]--
						tx[t] = tx[t-1]
					}
				}
			}
			board[ty[9]][tx[9]] = 1
		}
	}
	visits := 0
	for _, by := range board {
		for _, bx := range by {
			if bx == 1 {
				visits++
			}
		}
	}
	fmt.Printf("Visits part2: %d\n", visits)
}

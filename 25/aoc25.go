package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./input25.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	var vals []int
	for _, line := range lines {
		vals = append(vals, snafuToDec(line))
	}
	part1(vals)
}

func snafuToDec(n string) int {
	result := 0
	m := 1
	for i := len(n) - 1; i >= 0; i-- {
		digit := 0
		if n[i] == '1' {
			digit = 1
		} else if n[i] == '2' {
			digit = 2
		} else if n[i] == '-' {
			digit = -1
		} else if n[i] == '=' {
			digit = -2
		}
		result += digit * m
		m *= 5
	}
	return result
}

func decToSnafu(n int) string {
	result := ""
	cur := n
	for m := 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5; m > 0; m /= 5 {
		bestt := 0
		bestdist := math.MaxInt
		for t := -2; t <= 2; t++ {
			if Abs(cur-(t*m)) < bestdist {
				bestt = t
				bestdist = Abs(cur - (t * m))
			}
		}
		cur -= bestt * m
		if bestt == -2 {
			result += "="
		} else if bestt == -1 {
			result += "-"
		} else {
			result += fmt.Sprintf("%d", bestt)
		}
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(vals []int) {
	//fmt.Println(vals)
	sum := 0
	for _, n := range vals {
		sum += n
	}
	fmt.Printf("Sum part1: %d: %s\n", sum, decToSnafu((sum)))
}

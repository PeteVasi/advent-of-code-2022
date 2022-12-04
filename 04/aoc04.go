package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cast"
)

type sects struct {
	min1 int
	max1 int
	min2 int
	max2 int
}

func main() {
	raw, err := os.ReadFile("./input4.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	var sections []sects
	for _, line := range lines {
		e := strings.Split(line, ",")
		r1 := strings.Split(e[0], "-")
		r2 := strings.Split(e[1], "-")
		s := sects{min1: cast.ToInt(r1[0]), max1: cast.ToInt(r1[1]), min2: cast.ToInt(r2[0]), max2: cast.ToInt((r2[1]))}
		sections = append(sections, s)
	}
	//fmt.Println(sections)
	part1(sections)
	part2(sections)
}

func part1(sections []sects) {
	contains := 0
	for _, sec := range sections {
		if (sec.min1 <= sec.min2 && sec.max1 >= sec.max2) ||
			(sec.min2 <= sec.min1 && sec.max2 >= sec.max1) {
			contains++
		}
	}
	fmt.Printf("Contains part1: %d\n", contains)
}

func part2(sections []sects) {
	overlap := 0
	for _, sec := range sections {
		if !((sec.max1 < sec.min2) ||
			(sec.max2 < sec.min1)) {
			overlap++
		}
	}
	fmt.Printf("Overlap part1: %d\n", overlap)
}

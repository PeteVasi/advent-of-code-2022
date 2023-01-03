package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type monkey struct {
	name  string
	op    string
	left  string
	right string
	val   int
}

func main() {
	raw, err := os.ReadFile("./input21.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	var monkies []monkey
	for _, line := range lines {
		var m monkey
		n, err := fmt.Sscanf(line, "%s %s %s %s", &m.name, &m.left, &m.op, &m.right)
		//fmt.Printf("%s  --  %d -- %s  -- %v\n", line, n, err, m)
		if n == 4 && err == nil {
			m.name = m.name[0:4]
			monkies = append(monkies, m)
		} else {
			n, err := fmt.Sscanf(line, "%s %d", &m.name, &m.val)
			if n == 2 && err == nil {
				m.name = m.name[0:4]
				monkies = append(monkies, m)
			} else {
				log.Fatal(err)
			}
		}
	}
	part1(monkies)
}

func part1(monkies []monkey) {
	//fmt.Println(monkies)
	total := 0
	for _, m := range monkies {
		//fmt.Println(m.name)
		if m.name == "root" {
			total = recurPart1(m, monkies)
			break
		}
	}
	fmt.Printf("Total part1: %d\n", total)
}

func recurPart1(m monkey, monkies []monkey) int {
	if m.val != 0 {
		return m.val
	}

	lval := 0
	rval := 0
	for _, mo := range monkies {
		if mo.name == m.left {
			lval = recurPart1(mo, monkies)
		} else if mo.name == m.right {
			rval = recurPart1(mo, monkies)
		}
	}

	if m.op == "+" {
		return lval + rval
	} else if m.op == "-" {
		return lval - rval
	} else if m.op == "*" {
		return lval * rval
	} else if m.op == "/" {
		return lval / rval
	}

	log.Fatal("No op")
	return 0
}

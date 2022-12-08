package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Push(s string, v string) string {
	return s + v
}

func Pop(s string) (string, string) {
	l := len(s)
	return s[0 : l-1], s[l-1 : l]
}

func main() {
	raw, err := os.ReadFile("./input6.txt")
	if err != nil {
		log.Fatal(err)
	}
	line := string(raw)
	part1(line)
	part2(line)
}

func part1(line string) {
	test := line[0:4]
	for i := 4; i < len(line); i++ {
		found := true
		for j := 0; j < len(test); j++ {
			if strings.Count(test, test[j:j+1]) > 1 {
				found = false
				break
			}
		}
		if found {
			fmt.Printf("Part 1 Found, index: %d\n", i)
			break
		}
		test = test[1:] + line[i:i+1]
	}
}

func part2(line string) {
	test := line[0:14]
	for i := 14; i < len(line); i++ {
		found := true
		for j := 0; j < len(test); j++ {
			if strings.Count(test, test[j:j+1]) > 1 {
				found = false
				break
			}
		}
		if found {
			fmt.Printf("Part 2 Found, index: %d\n", i)
			break
		}
		test = test[1:] + line[i:i+1]
	}
}

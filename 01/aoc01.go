package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")

	var nums []int
	for _, line := range lines {
		if len(line) == 0 {
			nums = append(nums, -1) // cheaty, but I know my inputs
		} else {
			n, _ := strconv.Atoi(line)
			nums = append(nums, n)
		}
	}
	part1(nums)
	part2(nums)
}

func part1(nums []int) {
	largest := 0
	runningsum := 0
	for _, n := range nums {
		if n == -1 {
			if runningsum > largest {
				largest = runningsum
			}
			runningsum = 0
		} else {
			runningsum += n
		}
	}
	if runningsum > largest {
		largest = runningsum
	}
	fmt.Printf("Largest: %d\n", largest)
}

func part2(nums []int) {
	var sums []int
	runningsum := 0
	for _, n := range nums {
		if n == -1 {
			sums = append(sums, runningsum)
			runningsum = 0
		} else {
			runningsum += n
		}
	}
	if runningsum != 0 {
		sums = append(sums, runningsum)
	}
	sort.Ints(sums)
	//fmt.Println("Sums: ", sums)

	top3sum := sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]
	fmt.Printf("Sum fo top 3 sums: %d\n", top3sum)
}

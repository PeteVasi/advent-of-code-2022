package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type cube struct {
	x int
	y int
	z int
}

func main() {
	file, err := os.Open("input18sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	lines, _ := csvReader.ReadAll()
	var cubes []cube
	for _, line := range lines {
		var c cube
		c.x, _ = strconv.Atoi(line[0])
		c.y, _ = strconv.Atoi(line[1])
		c.z, _ = strconv.Atoi(line[2])
		cubes = append(cubes, c)
	}
	part1(cubes)
	//part2(cubes)
}

func contains(cubes []cube, c cube) bool {
	for _, v := range cubes {
		if v == c {
			return true
		}
	}
	return false
}

func part1(cubes []cube) {
	//fmt.Println(cubes)
	sum := 6 * len(cubes)
	for _, c := range cubes {
		if contains(cubes, cube{x: c.x + 1, y: c.y, z: c.z}) {
			sum--
		}
		if contains(cubes, cube{x: c.x - 1, y: c.y, z: c.z}) {
			sum--
		}
		if contains(cubes, cube{x: c.x, y: c.y + 1, z: c.z}) {
			sum--
		}
		if contains(cubes, cube{x: c.x, y: c.y - 1, z: c.z}) {
			sum--
		}
		if contains(cubes, cube{x: c.x, y: c.y, z: c.z + 1}) {
			sum--
		}
		if contains(cubes, cube{x: c.x, y: c.y, z: c.z - 1}) {
			sum--
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type dir struct {
	path           string
	size           int
	recursive_size int
}

func main() {
	raw, err := os.ReadFile("./input7.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	var dirs []dir
	curdir := ""
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "$ cd /" {
			curdir = ""
		} else if line == "$ cd .." {
			curdir = curdir[0:strings.LastIndex(curdir, "/")]
		} else if line[0:4] == "$ cd" {
			curdir += "/" + line[5:]
		} else if line == "$ ls" {
			i++
			size := 0
			for i < len(lines) && lines[i][0] != '$' {
				var fname string
				var fsize int
				n, err := fmt.Sscanf(lines[i], "%d %s", &fsize, &fname)
				if n == 2 && err == nil {
					size += fsize
				}
				i++
			}
			i--
			d := dir{path: curdir, size: size}
			dirs = append(dirs, d)
		}
	}

	for i := range dirs {
		for _, dir2 := range dirs {
			if dirs[i].path == "" || (dir2.path != "" && strings.Contains(dir2.path, dirs[i].path)) {
				dirs[i].recursive_size += dir2.size
				//fmt.Printf("%s in %s, add: %d\n", dirs[i].path, dir2.path, dirs[i].recursive_size)
			}
		}
	}

	//fmt.Println(dirs)
	part1(dirs)
	part2(dirs)
}

func part1(dirs []dir) {
	sum := 0
	for _, dir := range dirs {
		if dir.recursive_size <= 100000 {
			sum += dir.recursive_size
		}
	}
	fmt.Printf("Part 1, sizes sum: %d\n", sum)
}

func part2(dirs []dir) {
	unused := 70000000 - dirs[0].recursive_size
	need := 30000000 - unused
	best := dirs[0].recursive_size
	for _, dir := range dirs {
		if dir.recursive_size >= need && dir.recursive_size < best {
			best = dir.recursive_size
		}
	}
	fmt.Printf("Part 2, delete size: %d\n", best)
}

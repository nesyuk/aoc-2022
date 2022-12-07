package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dirSize(sc *bufio.Scanner, sizeSum *int, minDelSize int, delDirSize *int) int {
	sc.Scan()
	line := sc.Text()
	size := 0
	dirs := 0
	if line == "$ ls" {
		for sc.Scan() {
			line = sc.Text()
			if strings.HasPrefix(line, "dir") {
				dirs++
			} else if !strings.HasPrefix(line, "$") {
				fileSize, _ := strconv.Atoi(strings.Split(line, " ")[0])
				size += fileSize
			} else {
				break
			}
		}
	}
	for strings.HasPrefix(line, "$ cd") {
		if strings.HasPrefix(line, "$ cd ..") {
			if size <= 100000 {
				*sizeSum += size
			}
			if size >= minDelSize && size < *delDirSize {
				*delDirSize = size
			}
			return size
		}
		size += dirSize(sc, sizeSum, minDelSize, delDirSize)

		sc.Scan()
		line = sc.Text()
	}
	if size <= 100000 {
		*sizeSum += size
	}
	if size >= minDelSize && size < *delDirSize {
		*delDirSize = size
	}
	return size
}

func main() {
	inputFile := "./day07/data/input.txt"
	r, _ := os.Open(inputFile)
	sc := bufio.NewScanner(r)

	sum := 0
	dummy := 0
	totalSize := dirSize(sc, &sum, dummy, &dummy)
	fmt.Println("part1", sum)
	r.Close()

	r, _ = os.Open(inputFile)
	defer r.Close()
	sc = bufio.NewScanner(r)

	delDirSize := totalSize
	minDelSize := 30000000 - (70000000 - delDirSize)
	_ = dirSize(sc, &sum, minDelSize, &delDirSize)
	fmt.Println("part2", delDirSize)
}

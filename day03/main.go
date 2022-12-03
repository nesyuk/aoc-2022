package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func part2(r io.Reader) int {
	sc := bufio.NewScanner(r)
	priorities := 0
	seenGroup := make(map[rune]int)
	count := 0

	for sc.Scan() {
		items := sc.Text()
		seen := make(map[rune]bool)
		for _, item := range items {
			if !seen[item] {
				seenGroup[item]++
				seen[item] = true
			}
		}
		count++
		if count == 3 {
			for item, times := range seenGroup {
				if times == 3 {
					if item <= 'Z' {
						priorities += int(item - 'A' + 27)
					} else {
						priorities += int(item - 'a' + 1)
					}
				}
			}
			count = 0
			seenGroup = make(map[rune]int)
		}
	}
	return priorities
}

func part1(r io.Reader) int {
	sc := bufio.NewScanner(r)
	priorities := 0

	for sc.Scan() {
		items := sc.Text()
		seen := make(map[rune]bool)
		for _, item := range items[:len(items)/2] {
			seen[item] = true
		}
		for _, item := range items[len(items)/2:] {
			if seen[item] {
				if item <= 'Z' {
					priorities += int(item - 'A' + 27)
				} else {
					priorities += int(item - 'a' + 1)
				}
				break
			}
		}
	}

	return priorities
}

func main() {
	r, _ := os.Open("./day03/data/input.txt")
	defer r.Close()

	//fmt.Println(part1(r))
	fmt.Println(part2(r))
}

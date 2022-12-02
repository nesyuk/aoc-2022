package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func topN(sc *bufio.Scanner, n int) int {
	max := make([]*int, n)
	for i := 0; i < n; i++ {
		max[i] = new(int)
	}
	total := 0
	for sc.Scan() {
		s := sc.Text()
		num, _ := strconv.Atoi(s)
		total += num
		if s == "" {
			updateMaxN(total, max...)
			total = 0
		}
	}
	// update last
	updateMaxN(total, max...)

	maxSum := 0
	for _, maxNth := range max {
		maxSum += *maxNth
	}
	return maxSum
}

func updateMaxN(n int, m ...*int) {
	for i := 0; i < len(m); i++ {
		if *m[i] < n {
			for j := len(m) - 1; j > i; j-- {
				*m[j] = *m[j-1]
			}
			*m[i] = n
			break
		}
	}
}

func main() {
	r, _ := os.Open("./day01/data/input.txt")
	defer r.Close()
	sc := bufio.NewScanner(r)

	// fmt.Println(topN(sc, 1)) // part1
	fmt.Println(topN(sc, 3)) // part2
}

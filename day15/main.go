package main

import (
	"aoc/2022/lib"
	"fmt"
	"io"
	"os"
	"sort"
)

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.x, c.y)
}

func (c coord) manhattan(c2 coord) int {
	return lib.Abs(c.x-c2.x) + lib.Abs(c.y-c2.y)
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals)-1; {
		itl1, itl2 := intervals[i], intervals[i+1]

		if itl2[0] > itl1[1]+1 {
			i++
			continue
		}
		// merge
		intervals[i][0] = lib.Min(intervals[i][0], intervals[i+1][0])
		intervals[i][1] = lib.Max(intervals[i][1], intervals[i+1][1])
		intervals = append(intervals[:i+1], intervals[i+2:]...)
	}
	return intervals
}

func findBeacon(rows map[int][][]int, from, to int) (int, int) {
	for y := from; y <= to; y++ {
		if len(rows[y]) > 1 {
			if len(rows[y]) != 2 {
				panic("wtf")
			}
			return rows[y][0][1] + 1, y
		}
	}
	return -1, -1
}

func count(rows map[int][][]int, row int, grid map[coord]int) int {
	seen := map[coord]bool{}
	for _, interval := range rows[row] {
		for x := interval[0]; x <= interval[1]; x++ {
			c := coord{x, row}
			if _, exist := grid[coord{x, row}]; !exist {
				seen[c] = true
			}
		}
	}
	return len(seen)
}

func scanPositions(r io.Reader) (map[int][][]int, map[coord]int) {
	grid := map[coord]int{}
	rows := map[int][][]int{}
	var s, b coord
	_, err := fmt.Fscanf(r, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x, &s.y, &b.x, &b.y)
	for err == nil {
		grid[s] = 0
		grid[b] = 1
		distance := s.manhattan(b)
		top, bottom := s.y, s.y
		for start, end := s.x-distance, s.x+distance; start <= end; start, end = start+1, end-1 {
			rows[top] = merge(append(rows[top], []int{start, end}))
			rows[bottom] = merge(append(rows[bottom], []int{start, end}))
			top--
			bottom++
		}
		_, err = fmt.Fscanf(r, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x, &s.y, &b.x, &b.y)

		// print progress
		fmt.Print(".")
	}
	fmt.Println()
	return rows, grid
}

func main() {
	r, _ := os.Open("./day15/data/input.txt")
	defer r.Close()

	rows, grid := scanPositions(r)

	// part 1
	fmt.Println(count(rows, 2000000, grid))

	// part 2
	x, y := findBeacon(rows, 0, 4000000)
	fmt.Println(x*4000000 + y)
}

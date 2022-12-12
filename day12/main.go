package main

import (
	"bufio"
	"fmt"
	"os"
)

func shortestDistance(grid [][]int, sx, sy int, checkHeightFn func(height int) bool, checkTargetFn func(x, y int) bool) int {
	visit := [][]int{{sx, sy}}
	seen := map[string]string{fmt.Sprintf("%d-%d", sx, sy): ""}
	steps := 0

	for len(visit) > 0 {
		count := len(visit)
		for c := 0; c < count; c++ {
			x, y := visit[0][0], visit[0][1]
			visit = visit[1:]

			if checkTargetFn(x, y) {
				return steps
			}
			// top
			check(grid, x, y, x-1, y, checkHeightFn, &visit, &seen)
			// bottom
			check(grid, x, y, x+1, y, checkHeightFn, &visit, &seen)
			// left
			check(grid, x, y, x, y-1, checkHeightFn, &visit, &seen)
			// right
			check(grid, x, y, x, y+1, checkHeightFn, &visit, &seen)
		}
		steps++
	}
	return steps
}

func check(grid [][]int, sx, sy int, tx, ty int, checkHeight func(height int) bool, visit *[][]int, seen *map[string]string) {
	if tx < 0 || ty < 0 || tx > len(grid)-1 || ty > len(grid[0])-1 {
		return
	}
	idx := fmt.Sprintf("%d-%d", tx, ty)
	if _, exist := (*seen)[idx]; exist {
		return
	}
	if checkHeight(grid[tx][ty] - grid[sx][sy]) {
		*visit = append(*visit, []int{tx, ty})
		(*seen)[idx] = fmt.Sprintf("%d-%d", sx, sy)
	}
	return
}

func scanGrid(sc *bufio.Scanner) (grid [][]int, sx, sy int, tx, ty int) {
	grid = make([][]int, 0)
	i := 0
	for sc.Scan() {
		row := []rune(sc.Text())
		grid = append(grid, make([]int, len(row)))
		for j, ch := range row {
			if ch == 'S' {
				ch = 'a'
				sx, sy = i, j
			}
			if ch == 'E' {
				ch = 'z'
				tx, ty = i, j
			}
			height := int(ch - 'a' + 1)
			grid[len(grid)-1][j] = height
		}
		i++
	}
	return
}

func main() {
	r, _ := os.Open("./day12/data/input.txt")
	defer r.Close()
	sc := bufio.NewScanner(r)

	grid, sx, sy, tx, ty := scanGrid(sc)

	// part 1
	fmt.Println(shortestDistance(grid, sx, sy,
		func(height int) bool {
			return height <= 1
		}, func(x, y int) bool {
			return x == tx && y == ty
		}))

	// part 2
	fmt.Println(shortestDistance(grid, tx, ty,
		func(height int) bool {
			return height >= -1
		}, func(x, y int) bool {
			return grid[x][y] == 1
		}))
}

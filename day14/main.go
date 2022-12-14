package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func stones(sc *bufio.Scanner) (map[coord]int, int) {
	stonesMap := map[coord]int{}
	floor := 0
	for sc.Scan() {
		points := strings.Split(sc.Text(), " -> ")
		var prevX, prevY *int
		for _, point := range points {
			coords := strings.Split(point, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			if floor < y {
				floor = y
			}
			if prevX != nil && prevY != nil {
				if *prevX == x {
					from, to := orderCoords(*prevY, y)
					for i := from; i <= to; i++ {
						stonesMap[coord{x, i}] = 0
					}
				}
				if *prevY == y {
					from, to := orderCoords(*prevX, x)
					for i := from; i <= to; i++ {
						stonesMap[coord{i, y}] = 0
					}
				}
			}

			prevX = &x
			prevY = &y
		}
	}
	return stonesMap, floor
}

func orderCoords(n, m int) (from int, to int) {
	if n < m {
		return n, m
	}
	return m, n
}

func sandFlow(occupied map[coord]int, floor int) int {
	fallen := 0
	for {
		sand := coord{500, 0}
		_, isOccupied := occupied[sand]
		if isOccupied {
			break
		}
		for !isOccupied {
			if sand.y > floor {
				// part 1
				// return fallen

				// part 2
				isOccupied = true
				continue
			}
			down := coord{sand.x, sand.y + 1}
			if _, isOccupied = occupied[down]; !isOccupied {
				sand = down
				continue
			}
			left := coord{sand.x - 1, sand.y + 1}
			if _, isOccupied = occupied[left]; !isOccupied {
				sand = left
				continue
			}
			right := coord{sand.x + 1, sand.y + 1}
			if _, isOccupied = occupied[right]; !isOccupied {
				sand = right
				continue
			}
		}
		occupied[sand] = 1
		fallen++
	}
	return fallen
}

func main() {
	r, _ := os.Open("./day14/data/input.txt")
	defer r.Close()
	sc := bufio.NewScanner(r)

	stonesMap, floor := stones(sc)
	// part 1
	// floor = 10000
	fmt.Println(sandFlow(stonesMap, floor))
}

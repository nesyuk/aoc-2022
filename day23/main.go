package main

import (
	"bufio"
	"fmt"
	"os"
)

type Elf struct {
	r, c int
}

func (e Elf) String() string {
	return fmt.Sprintf("%d-%d", e.r, e.c)
}

func (e Elf) hasNeighbours(elves map[Elf]bool, direction [][]int) bool {
	for i := 0; i < 3; i++ {
		if _, exist := elves[e.move(direction[i])]; exist {
			return true
		}
	}
	return false
}

func (e Elf) move(direction []int) Elf {
	return Elf{e.r + direction[0], e.c + direction[1]}
}

func scan(sc *bufio.Scanner) map[Elf]bool {
	elves := map[Elf]bool{}
	row := 0
	for sc.Scan() {
		line := sc.Text()
		for col, ch := range line {
			if ch == '#' {
				elves[Elf{row, col}] = true
			}
		}
		row++
	}
	return elves
}

func grid(elves map[Elf]bool) (int, int, int, int) {
	minR, maxR, minC, maxC := 1000, -1, 1000, -1
	for elf := range elves {
		if minR > elf.r {
			minR = elf.r
		}
		if maxR < elf.r {
			maxR = elf.r
		}
		if minC > elf.c {
			minC = elf.c
		}
		if maxC < elf.c {
			maxC = elf.c
		}
	}
	return minR, maxR, minC, maxC
}

func move(elves map[Elf]bool, rounds int) (map[Elf]bool, int) {
	directions := [][][]int{
		{{-1, 0}, {-1, -1}, {-1, 1}},
		{{1, 0}, {1, -1}, {1, 1}},
		{{0, -1}, {-1, -1}, {1, -1}},
		{{0, 1}, {-1, 1}, {1, 1}}}

	for round := 1; round <= rounds; round++ {
		proposed := map[Elf][]Elf{} // to => from
		haveToMove := map[Elf]bool{}
		for elf := range elves {
			for _, direction := range directions {
				if elf.hasNeighbours(elves, direction) {
					haveToMove[elf] = true
				}
			}
		}
		if len(haveToMove) == 0 {
			return elves, round
		}
	ELVES:
		for elf := range haveToMove {
			for _, direction := range directions {
				to := elf.move(direction[0])
				if !elf.hasNeighbours(elves, direction) {
					if _, exist := proposed[to]; !exist {
						proposed[to] = make([]Elf, 0)
					}
					proposed[to] = append(proposed[to], elf)
					continue ELVES
				}
			}
		}
		for to, from := range proposed {
			if len(from) > 1 {
				continue
			}
			// move
			delete(elves, from[0])
			elves[to] = true
		}
		directions = append(directions[1:], directions[0])
	}
	return elves, 0
}

func main() {
	r, _ := os.Open("./day23/data/input.txt")
	defer r.Close()
	sc := bufio.NewScanner(r)
	elves := scan(sc)

	// part 1
	// elves, _ = move(elves, 10)
	// x1, x2, y1, y2 := grid(elves)
	// fmt.Println((x2-x1+1)*(y2-y1+1) - len(elves))

	// part 2
	_, rounds := move(elves, 1000000)
	fmt.Println(rounds)
}

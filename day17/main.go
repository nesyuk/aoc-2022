package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type input struct {
	r    int
	j    int
	cave string
}

type output struct {
	input
	counter  int
	cut      uint64
	totalCut uint64
}

func fallRocks(cave Rock, rocks []Rock, jets string, maxRocks int) uint64 {
	j := 0
	totalCut := uint64(0)
	start := input{0, 0, ""}
	memo := map[input]output{}
	for counter := 0; counter < maxRocks; counter++ {
		r := counter % len(rocks)
		if result, exist := memo[start]; exist {
			// found a cycle

			cycleLen := counter - result.counter
			cyclesLeft := (maxRocks - counter) / cycleLen

			// cycles left * cycle cut
			totalCut += uint64(cyclesLeft) * (totalCut - result.totalCut)

			// move counter & finish
			counter += cyclesLeft * cycleLen
			left := uint64(0)
			for ; counter < maxRocks; counter++ {
				_, _, cave, totalCut = memo[start].r, memo[start].j, load(memo[start].cave), totalCut+memo[start].cut
				left += memo[start].cut
				start = memo[start].input
			}
			return uint64(len(cave)) + totalCut
		}
		rock := rocks[r]
		var cut uint64
		cave, cut, j = moveRock(cave, rock, jets, j)
		totalCut += cut
		memo[start] = output{input{r, j, cave.export()}, counter, cut, totalCut - cut}
		start = memo[start].input
	}
	return uint64(len(cave)) + totalCut
}

func moveRock(cave Rock, rock Rock, jets string, j int) (Rock, uint64, int) {
	// add 0s
	for i := 0; i < len(rock)+3; i++ {
		if len(cave) < i+1 || cave[i] != 0 {
			cave = append([]int{0}, cave...)
		}
	}
	row := 0
	for ; row+len(rock)-1 < len(cave) && !rock.overlaps(cave[row:row+len(rock)]); row++ {
		moved := rock.move(rune(jets[j]))
		j = (j + 1) % len(jets)
		if !moved.overlaps(cave[row : row+len(rock)]) {
			rock = moved
		}
	}
	for i := 0; i < len(rock); i++ {
		cave[row+i-1] |= rock[i]
	}
	// trim 0s
	for cave[0] == 0 {
		cave = cave[1:]
	}
	floor := 0
	i := 0
	for ; i < len(cave) && floor != 127; i++ {
		floor |= cave[i]
	}
	cut := uint64(len(cave)) - uint64(i)
	cave = cave[:i]
	return cave, cut, j
}

type Rock []int

func newRock(s string) Rock {
	rock := Rock{}
	for _, row := range strings.Split(s, "\n") {
		row = ".." + row + strings.Repeat(".", 7-len(row)-2)
		rockRow := 0
		for i := 0; i < len(row); i++ {
			if row[i] == '#' {
				rockRow |= 1 << (len(row) - i - 1)
			}
		}
		rock = append(rock, rockRow)
	}
	return rock
}

func load(s string) Rock {
	rows := make([]int, 0)
	for _, row := range strings.Split(s, "\n") {
		d, _ := strconv.Atoi(row)
		rows = append(rows, d)
	}
	return rows
}

func (r Rock) export() string {
	strs := make([]string, 0)
	for _, row := range r {
		strs = append(strs, fmt.Sprintf("%d", row))
	}
	return strings.Join(strs, "\n")
}

func (r Rock) move(direction rune) Rock {
	if direction == '<' {
		return r.left()
	}
	if direction == '>' {
		return r.right()
	}
	panic("unknown direction")
}

func (r Rock) left() Rock {
	moved := make([]int, len(r))
	for i, row := range r {
		if row&(1<<6) != 0 {
			return r
		}
		moved[i] = row << 1
	}
	return moved
}

func (r Rock) right() Rock {
	moved := make([]int, len(r))
	for i, row := range r {
		if row&1 != 0 {
			return r
		}
		moved[i] = row >> 1
	}
	return moved
}

func (r Rock) overlaps(cave Rock) bool {
	for i, row := range r {
		if cave[i]&row != 0 {
			return true
		}
	}
	return false
}

func main() {
	r, _ := os.Open("./day17/data/input.txt")
	defer r.Close()
	jets := ""
	fmt.Fscanf(r, "%s", &jets)
	rocks := make([]Rock, 0)
	for _, s := range strings.Split("####\n\n.#.\n###\n.#.\n\n..#\n..#\n###\n\n#\n#\n#\n#\n\n##\n##", "\n\n") {
		rocks = append(rocks, newRock(s))
	}
	cave := make([]int, 0)
	fmt.Println(fallRocks(cave, rocks, jets, 2022))          // part 1
	fmt.Println(fallRocks(cave, rocks, jets, 1000000000000)) // part 2
}

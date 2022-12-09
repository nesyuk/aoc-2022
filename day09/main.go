package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type knot struct {
	x, y int
	next *knot
}

func (c *knot) isTail() bool {
	return c.next == nil
}

func (c *knot) distance(c2 *knot) int {
	w, h := c.x-c2.x, c.y-c2.y
	return int(math.Sqrt(float64(w*w + h*h)))
}

func (c *knot) step(dir rune) {
	switch dir {
	case 'U':
		c.y += 1
	case 'D':
		c.y -= 1
	case 'L':
		c.x -= 1
	case 'R':
		c.x += 1
	}
}

// -1: left or down, 1: up or right
func direction(from, to int) int {
	if to == from {
		return 0
	}
	if to > from {
		return 1
	}
	return -1
}

func (c *knot) moveTowards(c2 *knot) {
	c.y += direction(c.y, c2.y)
	c.x += direction(c.x, c2.x)
}

func moveTail(r io.Reader, knots int) int {
	tail := new(knot)
	head := tail
	for k := 1; k < knots; k++ {
		tail.next = new(knot)
		tail = tail.next
	}
	var dir rune
	var steps int
	visited := map[string]bool{"0-0": true}
	_, err := fmt.Fscanf(r, "%c %d", &dir, &steps)
	for err == nil {
		for step := 0; step < steps; step++ {
			head.step(dir)
			prev := head
			next := head.next
			for next != nil {
				if next.distance(prev) >= 2 {
					next.moveTowards(prev)
				}
				if next.isTail() {
					visited[fmt.Sprintf("%d-%d", next.x, next.y)] = true
				}
				prev, next = next, next.next
			}
		}
		_, err = fmt.Fscanf(r, "%c %d", &dir, &steps)
	}
	return len(visited)
}

func main() {
	r, _ := os.Open("./day09/data/input.txt")
	defer r.Close()

	//fmt.Println(moveTail(r, 1+1)) // part 1
	fmt.Println(moveTail(r, 1+9)) // part 2
}

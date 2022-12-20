package main

import (
	"aoc/2022/lib"
	"fmt"
	"io"
	"os"
)

type node struct {
	n    int
	next *node
	prev *node
	size *int
}

func (n *node) move(steps int) *node {
	if lib.Abs(steps) > *n.size-1 {
		steps = steps % (*n.size - 1)
	}
	m := n
	if steps > 0 {
		for i := 0; i < steps; i++ {
			m = m.next
		}
	}
	if steps < 0 {
		for i := 0; i >= steps; i-- {
			m = m.prev
		}
	}
	return m
}

func decrypt(queue []*node) {
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n.n == 0 {
			continue
		}
		m := n.move(n.n)
		if n == m {
			continue
		}
		// cut
		n.prev.next, n.next.prev = n.next, n.prev

		// insert
		m.next, n.prev, m.next.prev, n.next = n, m, n, m.next
	}
}

func part1(zero *node, queue []*node) int {
	decrypt(queue)
	return zero.move(1000).n + zero.move(2000).n + zero.move(3000).n
}

func part2(zero *node, initial []*node) int {
	times := 10
	key := 811589153
	for i, n := range initial {
		initial[i].n = n.n * key
	}
	for t := 0; t < times; t++ {
		queue := make([]*node, len(initial))
		for i, n := range initial {
			queue[i] = n
		}
		decrypt(queue)
	}
	return zero.move(1000).n + zero.move(2000).n + zero.move(3000).n
}

func input(r io.Reader) (*node, []*node) {
	queue := make([]*node, 0)
	var zero *node
	var n int
	_, err := fmt.Fscanf(r, "%d", &n)
	cur := new(node)
	first := cur
	size := 0
	for err == nil {
		size++
		cur.next = &node{n, nil, nil, &size}
		cur.next.prev = cur
		cur = cur.next
		if n == 0 {
			zero = cur
		}
		queue = append(queue, cur)
		_, err = fmt.Fscanf(r, "%d", &n)
	}
	head := first.next
	cur.next = head
	head.prev = cur
	return zero, queue
}

func main() {
	r, _ := os.Open("./day20/data/input.txt")
	defer r.Close()
	zero, queue := input(r)

	// fmt.Println(part1(zero, queue))
	fmt.Println(part2(zero, queue))
}

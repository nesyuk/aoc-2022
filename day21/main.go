package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type op string

func (o op) exec(l, r int) int {
	switch o {
	case " + ":
		return l + r
	case " - ":
		return l - r
	case " * ":
		return l * r
	case " / ":
		return l / r
	default:
		panic("unknown op")
	}
}

func (o op) revert(l, r int, isLeft bool) int {
	if isLeft {
		switch o {
		case " + ":
			return l - r
		case " - ":
			return l + r
		case " * ":
			return l / r
		case " / ":
			return l * r
		}
	}
	switch o {
	case " + ":
		return l - r
	case " - ":
		return r - l
	case " * ":
		return l / r
	case " / ":
		return r / l
	}
	panic("unknown op")
}

type monkey string

type calculation interface {
}

type expression struct {
	m1, m2 monkey
	op     op
}

type number struct {
	n int
}

func yellNumbers(m monkey, monkeys *map[monkey]calculation) int {
	num, ok := (*monkeys)[m].(*number)
	if ok {
		return num.n
	}
	exp := (*monkeys)[m].(*expression)
	n := exp.op.exec(yellNumbers(exp.m1, monkeys), yellNumbers(exp.m2, monkeys))
	(*monkeys)[m] = &number{n}
	return n
}

func guessNumbers(m monkey, monkeys *map[monkey]calculation) int {
	(*monkeys)["root"].(*expression).op = op(" - ")
	return guess(m, monkeys, &number{0})
}

func guess(m monkey, monkeys *map[monkey]calculation, result *number) int {
	if m == "humn" {
		return result.n
	}
	exp, ok := (*monkeys)[m].(*expression)
	if !ok {
		panic("wtf!")
	}
	left := tryYell(exp.m1, monkeys)
	right := tryYell(exp.m2, monkeys)
	if left == nil && right == nil || left != nil && right != nil {
		panic("wtf!!")
	}
	var newMonkey monkey
	var newResult int
	if left == nil {
		newMonkey = exp.m1
		newResult = exp.op.revert(result.n, *right, true)
	}

	if right == nil {
		newMonkey = exp.m2
		newResult = exp.op.revert(result.n, *left, false)
	}
	return guess(newMonkey, monkeys, &number{newResult})
}

func tryYell(m monkey, monkeys *map[monkey]calculation) *int {
	if m == "humn" {
		return nil
	}
	num, ok := (*monkeys)[m].(*number)
	if ok {
		return &num.n
	}
	exp := (*monkeys)[m].(*expression)
	l := tryYell(exp.m1, monkeys)
	r := tryYell(exp.m2, monkeys)
	if l == nil || r == nil {
		return nil
	}
	n := exp.op.exec(*l, *r)
	(*monkeys)[m] = &number{n}
	return &n
}

func readMonkeys(sc *bufio.Scanner) map[monkey]calculation {
	monkeys := map[monkey]calculation{}
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, ": ")
		n, err := strconv.Atoi(parts[1])
		if err == nil {
			monkeys[monkey(parts[0])] = &number{n}
		}

		for _, o := range []string{" + ", " - ", " * ", " / "} {
			if strings.Contains(parts[1], o) {
				m := strings.Split(parts[1], o)
				monkeys[monkey(parts[0])] = &expression{monkey(m[0]), monkey(m[1]), op(o)}
			}
		}
	}
	return monkeys
}

func main() {
	r, _ := os.Open("./day21/data/input.txt")
	defer r.Close()
	sc := bufio.NewScanner(r)
	monkeys := readMonkeys(sc)

	// fmt.Println(yellNumbers("root", &monkeys)) // part 1
	fmt.Println(guessNumbers("root", &monkeys)) // part 2
}

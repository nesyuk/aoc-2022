package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const WIDTH = 40

func signalStrength(sc *bufio.Scanner) int {
	strengthSum := 0
	sumCycleStrengthFn := sumCycleStrength()
	cycle := 0
	register := 1
	for sc.Scan() {
		cmd := sc.Text()
		cycle++

		// part 2
		printSymbol(cycle, register)

		// part 1
		strengthSum += sumCycleStrengthFn(cycle, register)

		if cmd != "noop" {
			cycle++

			// part 2
			printSymbol(cycle, register)

			// part 1
			strengthSum += sumCycleStrengthFn(cycle, register)

			n, _ := strconv.Atoi(strings.Replace(cmd, "addx ", "", 1))
			register += n
		}
	}
	return strengthSum
}

func sumCycleStrength() func(cycle int, register int) int {
	watermark := 20
	return func(cycle int, register int) int {
		if cycle == watermark {
			watermark += WIDTH
			return cycle * register
		}
		return 0
	}
}

func printSymbol(cycle int, register int) {
	fmt.Print(symbol(cycle, register))
	if cycle%WIDTH == 0 {
		fmt.Println()
	}
}

func symbol(cycle int, register int) string {
	pos := cycle%WIDTH - 1
	if pos >= register-1 && pos <= register+1 {
		return "#"
	}
	return "."
}

func main() {
	r, _ := os.Open("./day10/data/input.txt")
	defer r.Close()

	sc := bufio.NewScanner(r)
	fmt.Println(signalStrength(sc))
}

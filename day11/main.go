package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const MONKEYS = 8
const ROUNDS = 10000

type monkey struct {
	items   []int64
	op      func(worry int64) int64
	test    func(worry int64) *monkey
	handled int64
}

func (m *monkey) String() string {
	return fmt.Sprintf("%v", m.items)
}

func (m *monkey) addItem(item int64) {
	m.items = append(m.items, item)
}

func (m *monkey) inspectNextItem() (*monkey, int64) {
	if len(m.items) == 0 {
		panic("wtf")
	}
	worry := m.items[0]
	m.items = m.items[1:]
	// part 1
	// worry = m.op(worry) / 3
	worry = m.op(worry)
	toMonkey := m.test(worry)
	m.handled++
	return toMonkey, worry
}

func scanInput(r io.Reader) ([]*monkey, int64) {
	sc := bufio.NewScanner(r)

	monkeys := make([]*monkey, MONKEYS)
	for m := 0; m < MONKEYS; m++ {
		monkeys[m] = &monkey{}
	}
	mod := int64(1)
	for m := 0; m < MONKEYS; m++ {
		// skip 1
		sc.Scan()

		// items
		sc.Scan()
		str := sc.Text()
		itemsStrs := strings.Split(strings.Replace(str, "  Starting items: ", "", 1), ", ")
		for i := 0; i < len(itemsStrs); i++ {
			item, _ := strconv.ParseInt(itemsStrs[i], 10, 64)
			monkeys[m].items = append(monkeys[m].items, item)
		}

		// op
		sc.Scan()
		str = sc.Text()
		exp := strings.Split(strings.Replace(str, "  Operation: new = ", "", 1), " ")
		monkeys[m].op = func(worry int64) int64 {
			num := worry
			if exp[2] != "old" {
				var err error
				if num, err = strconv.ParseInt(exp[2], 10, 64); err != nil {
					panic("wtf")
				}
			}
			switch exp[1] {
			case "+":
				return worry + num
			case "*":
				return worry * num
			default:
				panic("wtf")
			}
		}

		// test
		sc.Scan()
		str = sc.Text()
		div, _ := strconv.ParseInt(strings.Replace(str, "  Test: divisible by ", "", 1), 10, 64)
		mod *= div

		// test: true
		sc.Scan()
		str = sc.Text()
		mTrue, _ := strconv.ParseInt(strings.Replace(str, "    If true: throw to monkey ", "", 1), 10, 64)

		// test: false
		sc.Scan()
		str = sc.Text()
		mFalse, _ := strconv.ParseInt(strings.Replace(str, "    If false: throw to monkey ", "", 1), 10, 64)

		monkeys[m].test = func(worry int64) *monkey {
			if worry%div == 0 {
				return monkeys[mTrue]
			}
			return monkeys[mFalse]
		}
	}
	return monkeys, mod
}

func main() {
	r, _ := os.Open("./day11/data/input.txt")
	defer r.Close()

	monkeys, mod := scanInput(r)
	for round := 1; round <= ROUNDS; round++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				toMonkey, item := m.inspectNextItem()
				toMonkey.addItem(item % mod)
			}
		}
		/*		if round == 1 || round == 20 || round%1000 == 0 {
				fmt.Printf("ROUND %d\n", round)
				for i, m := range monkeys {
					fmt.Printf(" Monkey %d inspected %v items \n", i, m.handled)
				}
			}*/
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].handled > monkeys[j].handled
	})
	fmt.Println(monkeys[0].handled * monkeys[1].handled)
}

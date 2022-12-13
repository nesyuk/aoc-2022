package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Comparable interface {
	Compare(Comparable) int
}

type number struct {
	num int
}

func (n *number) String() string {
	return fmt.Sprintf("%d", n.num)
}

func (n *number) Compare(other Comparable) int {
	switch n1 := other.(type) {
	case *number:
		if n.num < n1.num {
			return -1
		}
		if n.num > n1.num {
			return 1
		}
		return 0
	case *array:
		if len(n1.nums) == 0 {
			return 1
		}
		result := n.Compare(n1.nums[0])
		if result != 0 {
			return result
		}
		return -1
	}
	panic("wtf")
}

type array struct {
	nums []Comparable
}

func (n *array) String() string {
	str := make([]string, 0)
	for _, num := range n.nums {
		str = append(str, fmt.Sprintf("%v", num))
	}
	return "[" + strings.Join(str, ",") + "]"
}

func (n *array) Compare(other Comparable) int {
	switch n1 := other.(type) {
	case *number:
		return -1 * n1.Compare(n) // revert
	case *array:
		i := 0
		for ; i < len(n.nums); i++ {
			if i == len(n1.nums) {
				return 1
			}
			result := n.nums[i].Compare(n1.nums[i])
			if result != 0 {
				return result
			}
		}
		if i == len(n1.nums) {
			return 0
		}
		return -1
	}
	panic(fmt.Sprintf("wtf: %T\n", other))
}
func readArray(str string) Comparable {
	stack := make([]*array, 0)
	var arr *array
	for i := 0; i < len(str); {
		ch := str[i]
		switch ch {
		case '[':
			arr = &array{}
			stack = append(stack, arr)
			i++
		case ']':
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				return last
			}
			stack[len(stack)-1].nums = append(stack[len(stack)-1].nums, last)
			i++
		case ',':
			i++
		default:
			num := ""
			for ; str[i] <= '9' && str[i] >= '0'; i++ {
				num += string(str[i])
			}
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			stack[len(stack)-1].nums = append(stack[len(stack)-1].nums, &number{n})
		}
	}
	return arr
}

func part1(sc *bufio.Scanner) int {
	idx := 1
	count := 0
	for sc.Scan() {

		p1 := sc.Text()
		sc.Scan()
		p2 := sc.Text()
		// skip next
		sc.Scan()
		n1 := readArray(p1)
		n2 := readArray(p2)
		if n1.Compare(n2) == -1 {
			fmt.Println(idx)
			count += idx
		} else {
		}
		idx++
	}
	return count
}

func part2(sc *bufio.Scanner) int {
	div1 := &array{nums: []Comparable{&array{nums: []Comparable{&number{2}}}}}
	div2 := &array{nums: []Comparable{&array{nums: []Comparable{&number{6}}}}}
	comparables := []Comparable{div1, div2}
	for sc.Scan() {
		p1 := sc.Text()
		sc.Scan()
		p2 := sc.Text()
		comparables = append(comparables, readArray(p1))
		comparables = append(comparables, readArray(p2))
		// skip next
		sc.Scan()
	}
	sort.Slice(comparables, func(i, j int) bool {
		return comparables[i].Compare(comparables[j]) == -1
	})
	result := 1
	for idx, c := range comparables {
		if c == div1 || c == div2 {
			result *= idx + 1
		}
	}
	return result
}

func main() {
	r, _ := os.Open("./day13/data/input.txt")
	defer r.Close()
	sc := bufio.NewScanner(r)
	// fmt.Println(part1(sc))
	fmt.Println(part2(sc))
}

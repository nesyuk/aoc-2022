package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// read stacks
	rStacks, _ := os.Open("./day05/data/input_stacks.txt")
	defer rStacks.Close()

	stacksLen := 9

	stacks := make([][]string, stacksLen)
	for i := 0; i < stacksLen; i++ {
		stacks[i] = make([]string, 0)
	}
	sc := bufio.NewScanner(rStacks)
	for sc.Scan() {
		s := sc.Text()
		stackIdx := 0
		for j := 1; j < len(s); j += 4 {
			if s[j] == ' ' {
				stackIdx++
				continue
			}
			stacks[stackIdx] = append([]string{string(s[j])}, stacks[stackIdx]...)
			stackIdx++
		}
	}

	// read commands
	rCommands, _ := os.Open("./day05/data/input_commands.txt")
	defer rCommands.Close()

	var amount, from, to int
	_, err := fmt.Fscanf(rCommands, "move %d from %d to %d", &amount, &from, &to)
	for err == nil {
		from--
		to--
		// part 1
		//for i := 0; i < amount; i++ {
		//	stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
		//	stacks[from] = stacks[from][:len(stacks[from])-1]
		//}

		// part 2
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-amount:len(stacks[from])]...)
		stacks[from] = stacks[from][:len(stacks[from])-amount]

		_, err = fmt.Fscanf(rCommands, "move %d from %d to %d", &amount, &from, &to)
	}

	for i := 0; i < stacksLen; i++ {
		fmt.Print(stacks[i][len(stacks[i])-1])
	}
}

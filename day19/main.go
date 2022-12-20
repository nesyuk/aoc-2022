package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type resource int

const (
	ore resource = iota
	clay
	obsidian
	geode
)

func maxGain(sc *bufio.Scanner, t int) int {
	results := make(chan result, 100)
	scenario := 1
	wg := &sync.WaitGroup{}
	for sc.Scan() {
		blueprint := [4][4]int{} // robot -> resource -> count
		parts := strings.Split(sc.Text(), ".")
		blueprint[ore][ore], _ = strconv.Atoi(strings.Split(strings.Split(parts[0], "costs ")[1], " ")[0])
		blueprint[clay][ore], _ = strconv.Atoi(strings.Split(strings.Split(parts[1], "costs ")[1], " ")[0])
		blueprint[obsidian][ore], _ = strconv.Atoi(strings.Split(strings.Split(parts[2], "costs ")[1], " ")[0])
		blueprint[obsidian][clay], _ = strconv.Atoi(strings.Split(strings.Split(parts[2], "costs ")[1], " ")[3])
		blueprint[geode][ore], _ = strconv.Atoi(strings.Split(strings.Split(parts[3], "costs ")[1], " ")[0])
		blueprint[geode][obsidian], _ = strconv.Atoi(strings.Split(strings.Split(parts[3], "costs ")[1], " ")[3])

		robots := [4]int{1, 0, 0, 0}
		resources := [4]int{0, 0, 0, 0}
		memo := map[string]int{}

		wg.Add(1)
		go func(scenario int, results chan result) {

			defer wg.Done()
			results <- result{scenario: scenario, count: simulate(blueprint, resources, robots, t, &memo)}
		}(scenario, results)

		scenario++
	}
	wg.Wait()
	close(results)

	sum := 0
	for res := range results {
		fmt.Println(res)
		sum += res.count * res.scenario
	}
	return sum
}

func add(r1, r2 [4]int) [4]int {
	return [4]int{r1[0] + r2[0], r1[1] + r2[1], r1[2] + r2[2], r1[3] + r2[3]}
}

func cpy(r [4]int) [4]int {
	return [4]int{r[0], r[1], r[2], r[3]}
}

type result struct {
	scenario int
	count    int
}

func enough(r1, r2 [4]int) bool {
	return r1[0] >= r2[0] && r1[1] >= r2[1] && r1[2] >= r2[2] && r1[3] >= r2[3]
}

func sub(r1, r2 [4]int) [4]int {
	return [4]int{r1[0] - r2[0], r1[1] - r2[1], r1[2] - r2[2], r1[3] - r2[3]}
}

func simulate(blueprint [4][4]int, resources [4]int, robots [4]int, t int, memo *map[string]int) int {
	if t == 0 {
		return resources[geode]
	}
	memoId := fmt.Sprintf("%v-%v-%d", robots, resources, t)
	if maxSeen, seen := (*memo)[memoId]; seen {
		return maxSeen
	}
	max := 0
	for robot := 0; robot < 4; robot++ {
		if enough(resources, blueprint[robot]) {
			newRobots := cpy(robots)
			newRobots[robot]++
			newResources := add(sub(resources, blueprint[robot]), robots)
			count := simulate(blueprint, newResources, newRobots, t-1, memo)
			if max < count {
				max = count
			}
		}
	}
	newResources := add(resources, robots)
	count := simulate(blueprint, newResources, robots, t-1, memo)
	if max < count {
		max = count
	}
	(*memo)[memoId] = max
	return max
}

func main() {
	r, _ := os.Open("./day19/data/input.txt")
	defer r.Close()

	sc := bufio.NewScanner(r)
	fmt.Println(maxGain(sc, 32)) // 1349
}

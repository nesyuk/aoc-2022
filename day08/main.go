package main

import (
	"bufio"
	"fmt"
	"os"
)

func to2D(sc *bufio.Scanner) [][]int {
	arr := make([][]int, 0)
	for sc.Scan() {
		arr = append(arr, make([]int, 0))
		for _, ch := range []rune(sc.Text()) {
			arr[len(arr)-1] = append(arr[len(arr)-1], int(ch-'0'))
		}
	}
	return arr
}

func initSeen(n int) []int {
	maxSeen := make([]int, n)
	for i := 0; i < n; i++ {
		maxSeen[i] = -1
	}
	return maxSeen
}

func countVisible(sc *bufio.Scanner) int {
	arr := to2D(sc)
	n := len(arr)
	seen := make(map[string]bool)

	// top-bottom
	maxSeen := initSeen(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] > maxSeen[j] {
				seen[fmt.Sprintf("%d-%d", i, j)] = true
			}
			maxSeen[j] = max(maxSeen[j], arr[i][j])
		}
	}

	// bottom-top
	maxSeen = initSeen(n)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if arr[i][j] > maxSeen[j] {
				seen[fmt.Sprintf("%d-%d", i, j)] = true
			}
			maxSeen[j] = max(maxSeen[j], arr[i][j])
		}
	}

	// left-right
	maxSeen = initSeen(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[j][i] > maxSeen[j] {
				seen[fmt.Sprintf("%d-%d", j, i)] = true
			}
			maxSeen[j] = max(maxSeen[j], arr[j][i])
		}
	}

	// right-left
	maxSeen = initSeen(n)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if arr[j][i] > maxSeen[j] {
				seen[fmt.Sprintf("%d-%d", j, i)] = true
			}
			maxSeen[j] = max(maxSeen[j], arr[j][i])
		}
	}
	return len(seen)
}

func countVisiblePerTree(sc *bufio.Scanner) int {
	arr := to2D(sc)
	maxScore := 0
	n := len(arr)
	scores := make([][]int, n)
	for i := 0; i < n; i++ {
		scores[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s := score(arr, i, j)
			scores[i][j] = s
			if maxScore < s {
				maxScore = s
			}
		}
	}
	return maxScore
}

func score(arr [][]int, i, j int) int {
	// top
	t, b, l, r := i-1, i+1, j-1, j+1
	for ; t > 0 && arr[t][j] < arr[i][j]; t-- {
	}
	// bottom
	for ; b < len(arr)-1 && arr[b][j] < arr[i][j]; b++ {
	}
	// left
	for ; l > 0 && arr[i][l] < arr[i][j]; l-- {
	}
	// right
	for ; r < len(arr[0])-1 && arr[i][r] < arr[i][j]; r++ {
	}
	return (i - t) * (j - l) * (b - i) * (r - j)
}

func max(n, k int) int {
	if n > k {
		return n
	}
	return k
}

func main() {
	r, _ := os.Open("./day08/data/input.txt")
	defer r.Close()
	sc := bufio.NewScanner(r)
	// fmt.Println(countVisible(sc)) // part 1
	fmt.Println(countVisiblePerTree(sc)) // part 2
}

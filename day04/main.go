package main

import (
	"fmt"
	"io"
	"os"
)

func countOverlaps(r io.Reader) int {
	var s1, e1, s2, e2 int
	_, err := fmt.Fscanf(r, "%d-%d,%d-%d\n", &s1, &e1, &s2, &e2)
	count := 0
	for err == nil {
		// part 1 (includes)
		// if (s1 >= s2 && e1 <= e2) || (s2 >= s1 && e2 <= e1) {
		// 	count++
		// }

		// part 2 (overlaps)
		if (s1 <= e2 && s1 >= s2) || (s2 <= e1 && s2 >= s1) {
			count++
		}
		_, err = fmt.Fscanf(r, "%d-%d,%d-%d", &s1, &e1, &s2, &e2)
	}
	return count
}

func main() {
	r, _ := os.Open("./day04/data/input.txt")
	defer r.Close()

	fmt.Println(countOverlaps(r))
}

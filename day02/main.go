package main

import (
	"fmt"
	"io"
	"os"
)

func part2(r io.Reader) int {
	myTotalScore := 0

	var partner, outcome rune
	_, err := fmt.Fscanf(r, "%c %c\n", &partner, &outcome)
	for err == nil {
		partnerScore := int(partner - 'A' + 1)
		if outcome == 'Z' {
			// win
			myTotalScore += 6
			myTotalScore += (partnerScore % 3) + 1
		} else if outcome == 'Y' {
			// draw
			myTotalScore += 3
			myTotalScore += partnerScore
		} else {
			// lose
			myTotalScore += (partnerScore+1)%3 + 1
		}
		_, err = fmt.Fscanf(r, "%c %c\n", &partner, &outcome)
	}
	return myTotalScore
}

func part1(r io.Reader) int {
	var partner, me rune
	myTotalScore := 0

	_, err := fmt.Fscanf(r, "%c %c\n", &partner, &me)
	for err == nil {
		partnerScore := int(partner - 'A' + 1)
		myScore := int(me - 'X' + 1)
		myTotalScore += myScore

		if myScore == partnerScore {
			myTotalScore += 3
		} else if myScore == (partnerScore%3)+1 {
			myTotalScore += 6
		}

		_, err = fmt.Fscanf(r, "%c %c\n", &partner, &me)
	}
	return myTotalScore
}

func main() {
	r, _ := os.Open("./day02/data/input.txt")
	defer r.Close()

	// fmt.Println(part1(r))
	fmt.Println(part2(r)) // 13490
}

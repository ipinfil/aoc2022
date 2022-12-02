package main

import (
	"fmt"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

var choiceRank = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func getRoundScore(line []string) int {
	opponentChoiceRank := choiceRank[line[0]]

	var myChoiceRank int
	if line[1] == "X" {
		myChoiceRank = opponentChoiceRank - 1
	} else if line[1] == "Y" {
		myChoiceRank = opponentChoiceRank
	} else {
		myChoiceRank = opponentChoiceRank + 1
	}

	if myChoiceRank > 3 {
		myChoiceRank = 1
	} else if myChoiceRank < 1 {
		myChoiceRank = 3
	}

	if opponentChoiceRank == myChoiceRank { // draw
		return myChoiceRank + 3
	} else if (opponentChoiceRank < myChoiceRank && myChoiceRank - opponentChoiceRank == 1) || (myChoiceRank == 1 && opponentChoiceRank == 3) {
		return myChoiceRank + 6
	}

	return myChoiceRank
}

func main() {
	file := loading.LoadFile("2-1/input.txt")
	defer file.Close()

	scanner := loading.NewLineScanner(file)

	score := 0
	for scanner.Scan() {
		choices := strings.Split(scanner.Text(), " ")

		score += getRoundScore(choices)
	}
	fmt.Println(score)
}

package main

import (
	"fmt"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

var choiceRank = map[string]int{
	"A": 1, "X": 1, // rock
	"B": 2, "Y": 2, // paper
	"C": 3, "Z": 3, // scissors
}

func getRoundScore(line []string) int {
	opponentChoiceRank, myChoiceRank := choiceRank[line[0]], choiceRank[line[1]]

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

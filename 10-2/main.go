package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

type hashset map[int]struct{}

var X int = 1
var importantCycles hashset = hashset{
	40: struct{}{},
	80: struct{}{},
	120: struct{}{},
	160: struct{}{},
	200: struct{}{},
	240: struct{}{},
}
var cycles int

func AddCycle() {
	cycles++
	spritePositions := hashset{X - 1: struct{}{}, X: struct{}{}, X + 1: struct{}{}}

	if _, ok := spritePositions[(cycles - 1) % 40]; ok {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if _, ok := importantCycles[cycles]; ok {
		fmt.Print("\n")
	}
}

func main() {
	file := loading.LoadFile("10-1/input.txt")
	defer file.Close()
	scanner := loading.NewLineScanner(file)

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")

		if text[0] == "noop" {
			AddCycle()
			continue
		} else {
			number, _ := strconv.Atoi(text[1])

			for i := 0; i < 2; i++ {
				AddCycle()
			}

			X += number
		}
	}
}
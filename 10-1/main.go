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
	20: struct{}{},
	60: struct{}{},
	100: struct{}{},
	140: struct{}{},
	180: struct{}{},
	220: struct{}{},
}
var cycles int
var sum int

func AddCycle() {
	cycles++

	if _, ok := importantCycles[cycles]; ok {
		sum += cycles * X
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

	fmt.Println(sum)
}
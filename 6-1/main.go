package main

import (
	"fmt"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

type hashset map[string]struct{}

func main() {
	file := loading.LoadFile("6-1/input.txt")
	defer file.Close()
	scanner := loading.NewLineScanner(file)

	scanner.Scan()
	text := strings.Split(scanner.Text(), "")

	for i := 0; i < len(text); i++ {
		markers := hashset{}
		markers[text[i]] = struct{}{}
		markers[text[i + 1]] = struct{}{}
		markers[text[i + 2]] = struct{}{}
		markers[text[i + 3]] = struct{}{}

		if len(markers) == 4 {
			fmt.Println(markers)
			fmt.Println(i + 4)
			break
		}
	}
}
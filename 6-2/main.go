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

		for j := i; j < i + 14; j++ {
			if _, ok := markers[text[j]]; ok {
				break
			}

			markers[text[j]] = struct{}{}
		}

		if len(markers) == 14 {
			fmt.Println(i + 14)
			break
		}
	}
}
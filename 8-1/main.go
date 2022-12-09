package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

var treeMap [][]int

func IsVisible(x int, y int, height int) bool {
	visible := true
	for i := 0; i < y; i++ {
		if treeMap[i][x] >= height {
			visible = false
		}
	}

	if visible {
		return true;
	}
	visible = true

	for i := y + 1; i < len(treeMap); i++ {
		if treeMap[i][x] >= height {
			visible = false
		}
	}

	if visible {
		return true;
	}
	visible = true

	for i := 0; i < x; i++ {
		if treeMap[y][i] >= height {
			visible = false
		}
	}

	if visible {
		return true;
	}
	visible = true

	for i := x + 1; i < len(treeMap[y]); i++ {
		if treeMap[y][i] >= height {
			visible = false
		}
	}

	return visible
}

func main() {
	file := loading.LoadFile("8-1/input.txt")
	defer file.Close()

	scanner := loading.NewLineScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, "")

		row := []int{}

		for _, el := range split {
			num, err := strconv.Atoi(el)
			if err != nil {
				panic(err)
			}

			row = append(row, num)
		}

		treeMap = append(treeMap, row)
	}

	sum := 0

	for y, row := range treeMap {
		if y == 0 || y == len(treeMap) - 1 {
			sum += len(row)
			continue
		}

		for x, col := range row {
			if x == 0 || x == len(row) - 1 {
				if y == 1 {
					sum += len(treeMap) - 2
				}

				continue
			}

			if IsVisible(x, y, col) {
				sum++
			}
		}
	}

	fmt.Println(sum)
}
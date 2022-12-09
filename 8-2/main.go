package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

var treeMap [][]int

func GetScenicScore(x int, y int, height int) int {
	score := []int{}

	tmp := 0
	for i := y - 1; i >= 0; i-- {
		tmp++
		if treeMap[i][x] >= height {
			break
		}
	}
	score = append(score, tmp)
	tmp = 0

	for i := y + 1; i < len(treeMap); i++ {
		tmp++
		if treeMap[i][x] >= height {
			break
		}
	}
	score = append(score, tmp)
	tmp = 0

	for i := x - 1; i >= 0; i-- {
		tmp++
		if treeMap[y][i] >= height {
			break
		}
	}
	score = append(score, tmp)
	tmp = 0

	for i := x + 1; i < len(treeMap[y]); i++ {
		tmp++
		if treeMap[y][i] >= height {
			break
		}
	}
	score = append(score, tmp)
	tmp = 0

	for i, el := range score {
		if i == 0 {
			tmp = el
			continue
		}

		tmp = tmp * el
	}

	return tmp
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

	maxScore := 0

	for y, row := range treeMap {
		for x, col := range row {
			score := GetScenicScore(x, y, col)

			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println(maxScore)
}
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

func ToIntSlice(slice []string) []int {
	intSlice := []int{}

	for _, val := range slice {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		intSlice = append(intSlice, intVal)
	}

	return intSlice
}

func main() {
	file := loading.LoadFile("4-1/input.txt")
	scanner := loading.NewLineScanner(file)

	sum := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		first, second := split[0], split[1]
		firstRange, secondRange := ToIntSlice(strings.Split(first, "-")), ToIntSlice(strings.Split(second, "-"))

		if (firstRange[0] >= secondRange[0] && firstRange[1] <= secondRange[1]) || (secondRange[0] >= firstRange[0] && secondRange[1] <= firstRange[1]) {
			sum++
		}
	}

	fmt.Println(sum)
}
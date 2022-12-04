package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

type hashset map[int]struct{}

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

func ToHashSet(sectionRange []int) hashset {
	result := make(hashset)

	for i := sectionRange[0]; i <= sectionRange[1]; i++ {
		result[i] = struct{}{}
	}

	return result
}

func main() {
	file := loading.LoadFile("4-1/input.txt")
	scanner := loading.NewLineScanner(file)

	sum := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		first, second := split[0], split[1]
		firstRange, secondRange := ToIntSlice(strings.Split(first, "-")), ToIntSlice(strings.Split(second, "-"))
		firstRangeHashset := ToHashSet(firstRange)

		for i := secondRange[0]; i <= secondRange[1]; i++ {
			if _, ok := firstRangeHashset[i]; ok {
				sum++
				break
			}
		}
	}

	fmt.Println(sum)
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var numericRegex = regexp.MustCompile(`^[0-9]+$`)

func isNumeric(value string) bool {
	return numericRegex.MatchString(value)
}

func main() {
	file, err := os.Open("1-2/input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	actualElf := 0
	var caloriesPerElf []int
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			caloriesPerElf = append(caloriesPerElf, actualElf)
			actualElf = 0
		} else if isNumeric(text) {
			caloryValue, err := strconv.Atoi(text)
			if err != nil {
				continue
			}

			actualElf += caloryValue
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(caloriesPerElf)))
	caloriesPerElf = caloriesPerElf[0:3]
	result := 0

	for i := 0; i < len(caloriesPerElf); i++ {
		result += caloriesPerElf[i]
	}

	fmt.Println(result)
}
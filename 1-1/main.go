package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var numericRegex = regexp.MustCompile(`^[0-9]+$`)

func isNumeric(value string) bool {
	return numericRegex.MatchString(value)
}

func main() {
	file, err := os.Open("1/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	actualElf := 0
	max := 0
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			if actualElf > max {
				max = actualElf
			}

			actualElf = 0
		} else if isNumeric(text) {
			caloryValue, err := strconv.Atoi(text)
			if err != nil {
				continue
			}

			actualElf += caloryValue
		}
	}

	fmt.Println(max)
}
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

var stacks = make(map[int][]string)

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
	file := loading.LoadFile("5-1/input.txt")
	defer file.Close()
	scanner := loading.NewLineScanner(file)

	// initialize stacks
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "")

		if len(text) != 0 && text[1] == "1" {
			continue
		}

		if len(text) == 0 {
			break
		}

		for i := 1; i < len(text); i += 4 {
			if _, ok := stacks[i % 4 + i / 4]; !ok {
				stacks[i % 4 + i / 4] = []string{}
			}

			if text[i] != " " {
				stacks[i % 4 + i / 4] = append([]string{text[i]}, stacks[i % 4 + i / 4]...)
			}
		}
	}

	for scanner.Scan() {
		instruction := strings.ReplaceAll(scanner.Text(), "move ", "")
		instruction = strings.ReplaceAll(instruction, "from ", "")
		instruction = strings.ReplaceAll(instruction, "to ", "")
		splitInstruction := ToIntSlice(strings.Split(instruction, " "))
		fmt.Println(stacks)
		fmt.Println(splitInstruction)
		stacks[splitInstruction[2]] = append(stacks[splitInstruction[2]], stacks[splitInstruction[1]][len(stacks[splitInstruction[1]]) - splitInstruction[0]:]...)
		stacks[splitInstruction[1]] = stacks[splitInstruction[1]][:len(stacks[splitInstruction[1]]) - splitInstruction[0]]
	}

	fmt.Println(stacks)
	// HNSNMTLHQ
}
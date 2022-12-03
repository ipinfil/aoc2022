package main

import (
	"fmt"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

var characterPriorities = make(map[string]int, 52)

func PrepareCharacterPriorities() {
	priority := 1
	for i := 97; i <= 122; i++ {
		characterPriorities[string(rune(i))] = priority
		priority++
	}

	for i := 65; i <= 90; i++ {
		characterPriorities[string(rune(i))] = priority
		priority++
	}
}

func GetIntersection(arr1 []string, arr2 []string) string {
	characterMap1 := make(map[string]bool)

	for _, character := range arr1 {
		if _, ok := characterMap1[character]; ok {
			continue
		}
		characterMap1[character] = true;
	}

	for _, character := range arr2 {
		if _, ok := characterMap1[character]; ok {
			return character
		}
	}

	panic("No intersection found.")
}

func main() {
	file := loading.LoadFile("3-1/input.txt")
	defer file.Close()

	scanner := loading.NewLineScanner(file)

	PrepareCharacterPriorities()
	prioritySum := 0
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "")
		first, second := items[:len(items) / 2], items[len(items)/2:]

		prioritySum += characterPriorities[GetIntersection(first, second)]
	}

	fmt.Println(prioritySum)
}

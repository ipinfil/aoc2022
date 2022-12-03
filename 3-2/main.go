package main

import (
	"fmt"
	"log"
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

func ToMap(arr []string) map[string]bool {
	result := make(map[string]bool)
	for _, val := range arr {
		result[val] = true
	}

	return result
}

func GetIntersection(maps []map[string]bool) string {
	characterMap := make(map[string]int)

	for _, mappedArray := range maps {
		for character, _ := range mappedArray {
			if _, ok := characterMap[character]; ok {
				characterMap[character] = characterMap[character] + 1
				continue
			}
			characterMap[character] = 1;
		}
	}

	for character, value := range characterMap {
		if value == len(maps) {
			return character
		}
	}

	log.Fatal(maps, characterMap)
	panic("No intersection found.")
}

func main() {
	file := loading.LoadFile("3-1/input.txt")
	defer file.Close()

	scanner := loading.NewLineScanner(file)

	PrepareCharacterPriorities()
	prioritySum := 0
	rucksackNum := 0
	var rucksacks []map[string]bool

	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "")
		rucksacks = append(rucksacks, ToMap(items))
		rucksackNum++

		if rucksackNum == 3 {
			prioritySum += characterPriorities[GetIntersection(rucksacks)]
			rucksacks = rucksacks[:0] // clear array, keep allocated memory
			rucksackNum = 0
		}
	}

	fmt.Println(prioritySum)
}

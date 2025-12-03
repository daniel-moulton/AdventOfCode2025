package main

import (
	"fmt"
	"strings"
	"os"
)


func main() {
	inputData, _ := os.ReadFile("inputs/day2.txt")
	inputs := strings.Split(string(inputData), ",")
	result1 := part1(inputs)
	result2 := part2(inputs)

	fmt.Printf("Part 1 Result: %d\n", result1)
	fmt.Printf("Part 2 Result: %d\n", result2)
}

func getBounds(input string) (int, int) {
	var start, end int
	fmt.Sscanf(input, "%d-%d", &start, &end)
	return start, end
}

func invalidID(id int) bool {
	idStr := fmt.Sprintf("%d", id)
	firstHalf, secondHalf := idStr[:len(idStr)/2], idStr[len(idStr)/2:]
	if firstHalf == secondHalf {
		return true
	}
	return false
}

func invalidIDPart2(id int) bool {
	idStr := fmt.Sprintf("%d", id)
	doubled_idStr := idStr + idStr
	doubled_idStr = doubled_idStr[1 : len(doubled_idStr)-1]
	return strings.Contains(doubled_idStr, idStr)
}

func part1(inputs []string) int {
	total := 0
	for _, input := range inputs {
		start, end := getBounds(input)
		for i := start; i <= end; i++ {
			if invalidID(i) {
				total += i
			}
		}
	}
	return total
}

func part2(inputs []string) int {
	total := 0
	for _, input := range inputs {
		start, end := getBounds(input)
		for i := start; i <= end; i++ {
			if invalidIDPart2(i) {
				total += i
			}
		}
	}
	return total
}
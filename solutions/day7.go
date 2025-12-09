package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	inputData, _ := os.ReadFile("inputs/day7.txt")
	inputs := strings.Split(string(inputData), "\n")
	result1 := part1(inputs)
	fmt.Printf("Part 1 Result: %d\n", result1)
}

func getStartingPoint(inputs []string) []int {
	for i:= 0; i < len(inputs); i++ {
		for j:= 0; j < len(inputs[i]); j++ {
			char := rune(inputs[i][j])
			if char == 'S' {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func part1(inputs []string) int {
	// startingPoint := getStartingPoint(inputs)
	numSplits := 0
	for i:= 0; i < len(inputs); i++ {
		for j:= 0; j < len(inputs[i]); j++ {
			char := rune(inputs[i][j])
			if char != 'S' && char != '|' {
				continue
			}
			if i+1 >= len(inputs) {
				continue
			}
			below_char := rune(inputs[i+1][j])
			if below_char == '.' {
				inputs[i+1] = inputs[i+1][:j] + "|" + inputs[i+1][j+1:]
			} else if below_char == '^' {
				// Add | to left and right of ^
				if j-1 >= 0 && rune(inputs[i][j-1]) == '.' {
					inputs[i+1] = inputs[i+1][:j-1] + "|" + inputs[i+1][j:]
				}
				if j+1 < len(inputs[i]) && rune(inputs[i][j+1]) == '.' {
					inputs[i+1] = inputs[i+1][:j+1] + "|" + inputs[i+1][j+2:]
				}
				numSplits++
			}
		}
	}
	return numSplits
}

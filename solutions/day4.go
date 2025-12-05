package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	inputData, _ := os.ReadFile("inputs/day4.txt")
	inputs := strings.Split(string(inputData), "\n")
	result1 := part1(inputs)
	fmt.Printf("Part 1 Result: %d\n", result1)
	result2 := part2(inputs)
	fmt.Printf("Part 2 Result: %d\n", result2)
}

func part1(inputs []string) int {
	total := 0
	for i, input := range inputs {
		for j, char := range input {
			if char == '.' {
				continue
			}
			numAdjacent := getAdjacentCount(inputs, i, j)
			if numAdjacent < 4 {
				total++
			}
		}
	}
	return total
}

func part2(inputs []string) int {
	total := 0
	repeat := true
	for repeat {
		prevTotal := total
		for i := range inputs {
			for j := range inputs[i] {
				char := rune(inputs[i][j])
				if char == '.' {
					continue
				}
				if char != '@' {
					continue
				}
				numAdjacent := getAdjacentCount(inputs, i, j)
				if numAdjacent < 4 {
					total++
					// Switch the character
					inputs[i] = inputs[i][:j] + "." + inputs[i][j+1:]
				}
			}
		}
		if total == prevTotal {
			repeat = false
		}
	}
	return total
}

func getAdjacentCount(grid []string, i int, j int) int {
	count := 0
	// Check all 8 possible directions
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},          {0, 1},
		{1, -1},  {1, 0}, {1, 1},
	}
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]

		val, err := getValueAt(grid, ni, nj)
		if err != nil {
			continue
		}
		if val == '@' {
			count++
		}
	}
	return count
}

func getValueAt(grid []string, i int, j int) (rune, error) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return 0, fmt.Errorf("out of bounds")
	}
	return rune(grid[i][j]), nil
}

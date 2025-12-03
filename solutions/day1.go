// Advent of Code - Day1 - Go Solution

package main

import (
	"fmt"
	"strings"
	"os"
)

type Instruction struct {
	direction string
	steps     int
}

func main() {
	inputData, _ := os.ReadFile("inputs/day1.txt")
	inputs := strings.Split(string(inputData), "\n")
	
	result1 := part1(inputs)
	result2 := part2(inputs)

	fmt.Printf("Part 1 Result: %d\n", result1)
	fmt.Printf("Part 2 Result: %d\n", result2)
}

func convertToInstructions(inputs string) Instruction {
	dir := string(inputs[0])
	steps := 0
	fmt.Sscanf(inputs[1:], "%d", &steps)
	return Instruction{direction: dir, steps: steps}
}

func performInstruction(instruction Instruction, start int) (int, int) {
	// Calculate the new position after the move
	new_position := start
	switch instruction.direction {
	case "L":
		new_position -= instruction.steps
	case "R":
		new_position += instruction.steps
	}

	// Calculate final position on the dial (Go doesn't handle negative mod well...)
	final_position := ((new_position % 100) + 100) % 100

	// Count how many times we step on position 0
	zero_crossings := 0
	current := start
	step := 1
	if instruction.direction == "L" {
		step = -1
	}
	
	for i := 0; i < instruction.steps; i++ {
		current += step
		if ((current % 100) + 100) % 100 == 0 {
			zero_crossings++
		}
	}

	return final_position, zero_crossings
}

func part1(inputs []string) int {
	count := 0
	start := 50
	for _, input := range inputs {
		if len(input) == 0 {
			continue // Skip empty lines
		}
		instruction := convertToInstructions(input)
		start, _ = performInstruction(instruction, start)

		if start == 0 {
			count++
		}
	}
	return count
}

func part2(inputs []string) int {
	count := 0
	start := 50
	for _, input := range inputs {
		if len(input) == 0 {
			continue // Skip empty lines
		}
		instruction := convertToInstructions(input)
		newStart, zero_crossings := performInstruction(instruction, start)
		start = newStart
		count += zero_crossings
	}
	return count
}
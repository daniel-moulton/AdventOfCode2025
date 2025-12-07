package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	inputData, _ := os.ReadFile("inputs/day6.txt")
	inputs := strings.Split(string(inputData), "\n")
	var allNumbers [][]int
	for i := 0; i < len(inputs) - 1; i++ {
		var numbers []int
		numStrs := strings.Fields(inputs[i])
		for _, ns := range numStrs {
			var num int
			fmt.Sscanf(ns, "%d", &num)
			numbers = append(numbers, num)
		}
		allNumbers = append(allNumbers, numbers)
	}
	var operators []rune
	operatorLine := inputs[len(inputs)-1]
	for _, char := range operatorLine {
		if char == '+' || char == '*' {
			operators = append(operators, char)
		}
	}
	result1 := part1(allNumbers, operators)
	fmt.Printf("Part 1 Result: %d\n", result1)

}

func part1(allNumbers [][]int, operators []rune) int {
	total := 0
	// fmt.Printf("Operators: %v\n", operators)

	for i:=0; i<len(operators); i++ {
		var nums []int
		for j:=0; j<len(allNumbers); j++ {
			// fmt.Printf("allNumbers[%d][%d]\n", j, i)
			nums = append(nums, allNumbers[j][i])
		}
		var result int
		switch operators[i] {
		case '+':
			result = 0
			for _, n := range nums {
				result += n
			}
		case '*':
			result = 1
			for _, n := range nums {
				result *= n
			}
		}
		total += result
	}
	return total

}
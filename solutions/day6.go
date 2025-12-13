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

	result2 := part2(inputs)
	fmt.Printf("Part 2 Result: %d\n", result2)
}

func part1(allNumbers [][]int, operators []rune) int {
	total := 0
	// fmt.Printf("Operators: %v\n", operators)

	for i:=0; i<len(operators); i++ {
		var nums []int
		for j:=0; j<len(allNumbers); j++ {
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

func part2(inputs []string) int {
	total := 0
	maxLineLength := 0
	for _, line := range inputs {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}

	var numsBlock []int
	var operator rune
	for i := maxLineLength - 1; i >= 0; i-- {
		nums := ""
		if isEmptyColumn(inputs, i) {
			// reset nums and operator
			numsBlock = numsBlock[:0]
			operator = 0
			continue
		}
		for j := 0; j < len(inputs); j++ {
			if i < len(inputs[j]) {
				if inputs[j][i] >= '0' && inputs[j][i] <= '9' {
					nums = nums + string(inputs[j][i])
				} else if inputs[j][i] == '+' || inputs[j][i] == '*' {
					operator = rune(inputs[j][i])
				}
			}
		}
		if nums != ""{
			var colNum int
			fmt.Sscanf(nums, "%d", &colNum)
			numsBlock = append(numsBlock, colNum)
		}
		if operator != 0 {
			switch operator {
			case '+':
				for _, n := range numsBlock {
					total += n
				}
			case '*':
				prod := 1
				for _, n := range numsBlock {
					prod *= n
				}
				total += prod
			}
		}
	}
	return total
}

func isEmptyColumn(inputs []string, col int) bool {
	for _, line := range inputs {
		if col < len(line) {
			if line[col] != ' ' {
				return false
			}
		}
	}
	return true
}
package main

import (
	"fmt"
	"strings"
	"os"
	"sort"
)

func main() {
	inputData, _ := os.ReadFile("inputs/day3.txt")
	inputs := strings.Split(string(inputData), "\n")
	result1 := part1(inputs)
	fmt.Printf("Part 1 Result: %d\n", result1)
	result2 := part2(inputs)
	fmt.Printf("Part 2 Result: %d\n", result2)
}

func part1(inputs []string) int {
	total := 0
	for _, input := range inputs {
		// Convert string to array of ints
		var numbers []int
		for _, numStr := range strings.Split(input, "") {
			var num int
			fmt.Sscanf(numStr, "%d", &num)
			numbers = append(numbers, num)
		}
		biggest := -1
		biggestIndex := -1
		for i, num := range numbers {
			if num > biggest && i < len(numbers)-1 {
				biggest = num
				biggestIndex = i
			}
		}

		secondBiggest := -1
		for j := biggestIndex + 1; j < len(numbers); j++ {
			if numbers[j] > secondBiggest {
				secondBiggest = numbers[j]
			}
		}
		total += biggest * 10 + secondBiggest
	}
	return total
}

func part2(inputs []string) int {
	total := 0
	for _, input := range inputs {
		fmt.Println("Processing input:", input)
		// Convert string to array of ints
		var numbers []int
		for _, numStr := range strings.Split(input, "") {
			var num int
			fmt.Sscanf(numStr, "%d", &num)
			numbers = append(numbers, num)
		}

		nums_count := 12
		nums_dict := make(map[int]int) // index, value
		// Loop to get the nums_count biggest numbers
		for i := 0; i < nums_count; i++ {
			biggest := -1
			biggestIndex := -1
			// Same logic as part1
			for j := len(numbers) - 1; j >= 0; j-- {
				num := numbers[j]
				if num > biggest {
					if _, exists := nums_dict[j]; exists {
						continue
					}
					biggest = num
					biggestIndex = j
				}
			}
			nums_dict[biggestIndex] = biggest
		}
		fmt.Println(nums_dict)

		// Sort the keys of nums_dict in ascending order
		sorted_keys := make([]int, 0, len(nums_dict))
		for k := range nums_dict {
			sorted_keys = append(sorted_keys, k)
		}
		sort.Ints(sorted_keys)

		// Build final number using values in nums_dict based on sorted keys
		final_number := 0
		for _, k := range sorted_keys {
			final_number = final_number*10 + nums_dict[k]
		}
		fmt.Println("Final number constructed:", final_number)
		total += final_number
	}
	return total
}
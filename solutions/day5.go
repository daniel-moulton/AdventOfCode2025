package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	inputData, _ := os.ReadFile("inputs/day5.txt")
	inputs := strings.Split(string(inputData), "\n")
	emptyLineIndex := -1
	for i, line := range inputs {
		if line == "" {
			emptyLineIndex = i
			break
		}
	}
	ranges := inputs[:emptyLineIndex]
	ids := inputs[emptyLineIndex+1:]
	result1 := part1(ids, ranges)
	fmt.Printf("Part 1 Result: %d\n", result1)

	result2 := part2(ranges)
	fmt.Printf("Part 2 Result: %d\n", result2)
}

func part1(inputs []string, ranges []string) int {
	total := 0
	for _, input := range inputs {
		for _, r := range ranges {
			var start, end, id int
			fmt.Sscanf(r, "%d-%d", &start, &end)
			fmt.Sscanf(input, "%d", &id)
			if id >= start && id <= end {
				total++
				break
			}
		}
	}
	return total
}

// Merge existing range tuples with each other
func mergeExistingRanges(rangeTuples [][2]int) ([][2]int, bool) {
	changed := false
	for i := 0; i < len(rangeTuples); i++ {
		for j := i + 1; j < len(rangeTuples); j++ {
			// Check if ranges i and j overlap
			if (rangeTuples[i][1] >= rangeTuples[j][0] && rangeTuples[i][0] <= rangeTuples[j][1]) ||
			   (rangeTuples[j][1] >= rangeTuples[i][0] && rangeTuples[j][0] <= rangeTuples[i][1]) {
				// Merge the ranges
				newStart := min(rangeTuples[i][0], rangeTuples[j][0])
				newEnd := max(rangeTuples[i][1], rangeTuples[j][1])
				
				// Replace range i with merged range
				rangeTuples[i][0] = newStart
				rangeTuples[i][1] = newEnd
				
				// Remove range j
				rangeTuples = append(rangeTuples[:j], rangeTuples[j+1:]...)
				changed = true
				j--
			}
		}
	}
	return rangeTuples, changed
}

// Add new ranges from strings to existing range tuples
func addNewRanges(rangeTuples [][2]int, ranges []string) [][2]int {
	for j := 0; j < len(ranges); j++ {
		var start, end int
		fmt.Sscanf(ranges[j], "%d-%d", &start, &end)
		foundOverlap := false
		// fmt.Printf("Ranges before adding (%d, %d): %v\n", start, end, rangeTuples)

		for i := 0; i < len(rangeTuples); i++ {
			if (end >= rangeTuples[i][0] && start <= rangeTuples[i][1]) {
				rangeTuples[i][0] = min(start, rangeTuples[i][0])
				rangeTuples[i][1] = max(end, rangeTuples[i][1])
				foundOverlap = true
				break
			}
		}
		if !foundOverlap {
			rangeTuples = append(rangeTuples, [2]int{start, end})
		}
	}
	return rangeTuples
}

func part2(ranges []string) int {
	total := 0
	var rangeTuples [][2]int
	
	rangeTuples = addNewRanges(rangeTuples, ranges)
	
	for {
		var changed bool
		rangeTuples, changed = mergeExistingRanges(rangeTuples)
		if !changed {
			break
		}
		// fmt.Printf("After merge iteration: %v\n", rangeTuples)
	}

	for _, r := range rangeTuples {
		total += r[1] - r[0] + 1
	}
	return total
}
package main

import (
	"fmt"
	"strings"
	"os"
	"math"
	"slices"
)

type JunctionBox struct {
	x int
	y int
	z int
	parent *JunctionBox // pointer to parent box in the circuit
	size int // size of the circuit this box belongs to (for root boxes)
}

func calculateDistance(jb1, jb2 JunctionBox) float64 {
	// Euclidean distance
	return math.Sqrt(float64((jb1.x-jb2.x)*(jb1.x-jb2.x) + (jb1.y-jb2.y)*(jb1.y-jb2.y) + (jb1.z-jb2.z)*(jb1.z-jb2.z)))
}

func findClosestPair(junctionBoxes []JunctionBox, closestPairs []struct {
	box1 *JunctionBox
	box2 *JunctionBox
}) (*JunctionBox, *JunctionBox) {
	minDistance := math.MaxFloat64
	var box1, box2 *JunctionBox
	found := false
	
	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			distance := calculateDistance(junctionBoxes[i], junctionBoxes[j])
			if distance < minDistance && !closestPairsContains(closestPairs, &junctionBoxes[i], &junctionBoxes[j]) {
				minDistance = distance
				box1 = &junctionBoxes[i]
				box2 = &junctionBoxes[j]
				found = true
			}
		}
	}
	
	if !found {
		// Return nil if no valid pair found
		return nil, nil
	}
	return box1, box2
}

func closestPairsContains(closestPairs []struct {
	box1 *JunctionBox
	box2 *JunctionBox
}, jb1, jb2 *JunctionBox) bool {
	for _, pair := range closestPairs {
		if (pair.box1 == jb1 && pair.box2 == jb2) || (pair.box1 == jb2 && pair.box2 == jb1) {
			return true
		}
	}
	return false
}

func getCircuitSizes(junctionBoxes []JunctionBox) []int {
	circuitSizeMap := make(map[*JunctionBox]int)
	for i := range junctionBoxes {
		root := find(&junctionBoxes[i])
		if _, exists := circuitSizeMap[root]; !exists {
			circuitSizeMap[root] = root.size
		}
	}
	var sizes []int
	for _, size := range circuitSizeMap {
		sizes = append(sizes, size)
	}
	return sizes
}

func getNBiggestCircuitSizes(sizes []int, n int) []int {
	if n > len(sizes) {
		n = len(sizes)
	}
	
	sortedSizes := make([]int, len(sizes))
	copy(sortedSizes, sizes)
	
	// Sort in ascending order then reverse for descending
	slices.Sort(sortedSizes)
	slices.Reverse(sortedSizes)

	return sortedSizes[:n]
}

func find(jb *JunctionBox) *JunctionBox {
	// Get the root of the circuit
	if jb.parent != jb {
		jb.parent = find(jb.parent)
	}
	return jb.parent
}

func union(jb1, jb2 *JunctionBox) {
	// Merge the circuits of jb1 and jb2
	root1 := find(jb1)
	root2 := find(jb2)
	if root1 != root2 {
		root2.parent = root1
		root1.size += root2.size
	}
}

func main() {
	inputData, _ := os.ReadFile("inputs/day8.txt")
	inputs := strings.Split(string(inputData), "\n")

	var junctionBoxes []JunctionBox
	for _, line := range inputs {
		var jb JunctionBox
		fmt.Sscanf(line, "%d,%d,%d", &jb.x, &jb.y, &jb.z)
		jb.parent = &jb
		jb.size = 1  // Initialize size to 1
		junctionBoxes = append(junctionBoxes, jb)
	}

	// Print the first box to verify parsing
	fmt.Println("First Junction Box:", junctionBoxes[0])
	result1 := part1(junctionBoxes)
	fmt.Println("Part 1 Result:", result1)
	result2 := part2(junctionBoxes)
	fmt.Println("Part 2 Result:", result2)
}

func part1(junctionBoxes []JunctionBox) int {
	numPairs := 1000
	topN := 3
	var closestPairs []struct {
		box1 *JunctionBox
		box2 *JunctionBox
	}
	for i := 0; i < numPairs; i++ {
		box1, box2 := findClosestPair(junctionBoxes, closestPairs)
		closestPairs = append(closestPairs, struct {
			box1    *JunctionBox
			box2    *JunctionBox
		}{box1, box2})
		union(box1, box2)
	}

	circuitSizes := getCircuitSizes(junctionBoxes)

	biggestSizes := getNBiggestCircuitSizes(circuitSizes, topN)

	product := 1
	for _, size := range biggestSizes {
		product *= size
	}

	return product
}

func part2(junctionBoxes []JunctionBox) int {
	var closestPairs []struct {
		box1 *JunctionBox
		box2 *JunctionBox
	}
	finish := false
	for !finish {
		box1, box2 := findClosestPair(junctionBoxes, closestPairs)
		closestPairs = append(closestPairs, struct {
			box1    *JunctionBox
			box2    *JunctionBox
		}{box1, box2})
		union(box1, box2)

		// Check if all boxes are now in the same circuit
		root := find(&junctionBoxes[0])
		finish = true
		for i := 1; i < len(junctionBoxes); i++ {
			if find(&junctionBoxes[i]) != root {
				finish = false
				break
			}
		}

		if finish {
			x1, x2 := box1.x, box2.x
			product := x1 * x2
			fmt.Println("Final Closest Pair:", box1, box2)
			fmt.Println("Product of their x-coordinates:", product)
			return product
		}
	}

	return 0
}

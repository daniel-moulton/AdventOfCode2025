package main

import (
	"fmt"
	"strings"
	"os"
)

type Tiles struct {
	x int
	y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func part1(tiles []Tiles) int {
	largestArea := 0
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			if tiles[i].x == tiles[j].x || tiles[i].y == tiles[j].y {
				continue
			}
			// Plus one to include boundaries
			area := (abs(tiles[i].x-tiles[j].x) + 1) * (abs(tiles[i].y-tiles[j].y) + 1)
			if area > largestArea {
				largestArea = area
			}
		}
	}
	return largestArea
}

func main() {
	inputData, _ := os.ReadFile("inputs/day9.txt")
	inputs := strings.Split(string(inputData), "\n")

	var tiles []Tiles
	for _, line := range inputs {
		var t Tiles
		fmt.Sscanf(line, "%d,%d", &t.x, &t.y)
		tiles = append(tiles, t)
	}

	result1 := part1(tiles)
	fmt.Printf("Part 1 Result: %d\n", result1)
}

package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
)

type Machine struct {
	endState int
	buttons  [][]int
	joltage  []int
}

func parseEndState(stateString string) int {
	stateStr := strings.Trim(stateString, "[]")
	
	endStateStr := ""
	for _, char := range stateStr {
		if char == '.' {
			endStateStr += "0"
		} else {
			endStateStr += "1"
		}
	}
	
	endStateInt, _ := strconv.ParseInt(endStateStr, 2, 64)
	return int(endStateInt)
}

func parseButtonList(buttonString string) []int {
	buttonStr := strings.Trim(buttonString, "()")
	
	button := []int{}
	for _, num := range strings.Split(buttonStr, ",") {
		val, _ := strconv.Atoi(strings.TrimSpace(num))
		button = append(button, val)
	}
	return button
}

func parseJoltageList(joltageString string) []int {
	joltageStr := strings.Trim(joltageString, "{}")
	
	joltage := []int{}
	for _, num := range strings.Split(joltageStr, ",") {
		val, _ := strconv.Atoi(strings.TrimSpace(num))
		joltage = append(joltage, val)
	}
	return joltage
}

func splitInput(line string) (int, [][]int, []int) {
	parts := strings.Split(line, " ")
	
	// fmt.Println("Part 0:", parts[0])
	
	endState := parseEndState(parts[0])
	
	var buttons [][]int
	for i := 1; i < len(parts)-1; i++ {
		button := parseButtonList(parts[i])
		buttons = append(buttons, button)
	}
	
	joltage := parseJoltageList(parts[len(parts)-1])
	
	return endState, buttons, joltage
}

func getButtonCombination(permutation int, buttons [][]int) [][]int {
	var buttonsToPress [][]int
	n := len(buttons)
	for j := 0; j < n; j++ {
		bitmask := 1 << j
		if (permutation & bitmask) != 0 { // Check if bit j is set
			buttonsToPress = append(buttonsToPress, buttons[j])
		}
	}
	return buttonsToPress
}

func simulateButtonPresses(buttonsToPress [][]int, machineLen int) int {
	startingVal := 0
	for _, button := range buttonsToPress {
		for _, lightIndex := range button {
			// Index needs to be reversed so we can bit shift correctly
			indx := machineLen - 1 - lightIndex
			startingVal ^= (1 << indx)
		}
	}
	return startingVal
}

func findMinimumPresses(machine Machine) int {
	lowestPresses := -1
	machineLen := len(machine.joltage)
	n := len(machine.buttons)
	numPermutations := (1 << n) - 1 // 2^n - 1 non-empty combos (assumes lights will never start in the end state)
	
	/*
	Logic here is that pressng a button twice will cancel out so we only
	need to consider pressing a button once or not at all.
	So we can represent this as a binary number where each bit represents if
	we press the button or not. So 9 (1001) would mean press button 0 and 3.
	So for each permutation we can bitwise AND each bit to see if we press
	the button or not (credit to https://takeuforward.org/bit-manipulation/check-if-kth-bit-is-set-or-not).
	Then for each case, we will know what buttons to press, and can then for
	each light that button affects, we can toggle it by XORing our startingVal
	with the indexes of the lights affected.
	*/
	for i := 1; i <= numPermutations; i++ {
		buttonsToPress := getButtonCombination(i, machine.buttons)
		finalState := simulateButtonPresses(buttonsToPress, machineLen)
		
		if finalState == machine.endState {
			presses := len(buttonsToPress)
			if lowestPresses == -1 || presses < lowestPresses {
				lowestPresses = presses
			}
		}
	}
	return lowestPresses
}

func part1(machines []Machine) int {
	totalPresses := 0
	for _, machine := range machines {
		minPresses := findMinimumPresses(machine)
		totalPresses += minPresses
	}
	return totalPresses
}

func main() {
	inputData, _ := os.ReadFile("inputs/day10.txt")
	inputs := strings.Split(string(inputData), "\n")

	machines := []Machine{}
	for _, line := range inputs {
		endState, buttons, joltage := splitInput(line)
		machine := Machine{
			endState: endState,
			buttons:  buttons,
			joltage:  joltage,
		}
		machines = append(machines, machine)
	}

	result1 := part1(machines)
	fmt.Println("Part 1 Result:", result1)
}
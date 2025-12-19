package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/BooleanCat/go-functional/v2/it"
)

//go:embed input.txt
var input string //= "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}\n[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}\n[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}"

func parseLights(lights string) []bool {

	return slices.Collect(it.Map(slices.Values([]byte(lights)), func(ch byte) bool {
		return ch == '#'
	}))
}

func parseButtons(buttons []string) [][]int {
	return slices.Collect(it.Map(slices.Values(buttons), func(button string) []int {
		buttonItems := strings.Split(button[1:len(button)-1], ",")
		return slices.Collect(it.Map(slices.Values(buttonItems), func(digit string) int {
			return int(digit[0] - '0')
		}))
	}))
}

func parseJoltage(joltage string) []int {
	joltageItems := strings.Split(joltage[1:len(joltage)-1], ",")
	return slices.Collect(it.Map(slices.Values(joltageItems), func(digit string) int {
		v, _ := strconv.Atoi(digit)
		return v
	}))
}

type Machine struct {
	lights  []bool
	buttons [][]int
	joltage []int
}

type Part1Solution struct {
	state          []bool
	buttonsPressed []int
}

func pressButton(state []bool, button []int) {
	for _, light := range button {
		state[light] = !state[light]
	}
}

func solvePart1(m *Machine) []int {
	solutions := make([]Part1Solution, 0, len(m.buttons))
	for buttonIdx, button := range m.buttons {
		state := make([]bool, len(m.lights))
		pressButton(state, button)
		solutions = append(solutions, Part1Solution{state: state, buttonsPressed: []int{buttonIdx}})
	}
	//fmt.Println(solutions)

	for {
		curLen := len(solutions)
		for i := 0; i < curLen; i++ {
			solution := solutions[i]
			lastButton := solution.buttonsPressed[len(solution.buttonsPressed)-1]
			butLastButton := -1
			if len(solution.buttonsPressed) > 1 {
				butLastButton = solution.buttonsPressed[len(solution.buttonsPressed)-2]
			}
			if slices.Equal(m.lights, solution.state) {
				return solution.buttonsPressed
			}

			toAppend := make([]Part1Solution, 0, len(m.buttons))
			for buttonIdx, button := range m.buttons {
				if buttonIdx != butLastButton || buttonIdx != lastButton {
					nextState := slices.Clone(solution.state)
					pressButton(nextState, button)
					nextButtonsPressed := slices.Clone(solution.buttonsPressed)
					nextButtonsPressed = append(nextButtonsPressed, buttonIdx)
					nextSolution := Part1Solution{state: nextState, buttonsPressed: nextButtonsPressed}
					toAppend = append(toAppend, nextSolution)
				}
			}
			solutions[i] = toAppend[0]
			solutions = append(solutions, toAppend[1:]...)
		}

		//fmt.Println(len(solutions))
	}
}

func main() {
	lines := strings.Split(input, "\n")

	machines := make([]Machine, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		lights := parseLights(parts[0][1 : len(parts[0])-1])
		buttons := parseButtons(parts[1 : len(parts)-1])
		joltage := parseJoltage(parts[len(parts)-1])

		machines = append(machines, Machine{lights: lights, buttons: buttons, joltage: joltage})
	}
	//fmt.Println(machines)

	for _, m := range machines {
		fmt.Println(m)
		fmt.Println(solvePart1(&m))
	}
}

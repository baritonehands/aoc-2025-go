package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/BooleanCat/go-functional/v2/it"
	pq "github.com/baritonehands/aoc-2025-go/utils/priority_queue"
)

// go:embed input.txt
var input string = "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}\n[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}\n[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}"

func parseButtons(buttons []string) [][]int {
	return slices.Collect(it.Map(slices.Values(buttons), func(button string) []int {
		buttonItems := strings.Split(button[1:len(button)-1], ",")
		return slices.Collect(it.Map(slices.Values(buttonItems), func(digit string) int {
			return int(digit[0] - '0')
		}))
	}))
}

type Machine struct {
	lights  string
	buttons [][]int
	joltage string
}

type Part1Solution struct {
	state          string
	buttonsPressed []int
}

type ButtonPress struct {
	dest   string
	button int
}

func pressButton(state string, button []int) string {
	mutState := []byte(state)
	for _, light := range button {
		if mutState[light] == '.' {
			mutState[light] = '#'
		} else {
			mutState[light] = '.'
		}
	}
	return string(mutState)
}

func (m *Machine) walkPath(cameFrom map[ButtonPress]ButtonPress, current ButtonPress) []int {
	ret := []int{}
	for {
		if next, ok := cameFrom[current]; ok {
			current = next
			if current.button != -1 {
				ret = append(ret, current.button)
			}
		} else {
			break
		}
	}
	return ret
}

func solvePart1(m *Machine) []int {
	emptyState := strings.Repeat(".", len(m.lights))
	fScore := map[string]int{emptyState: math.MaxInt}
	fScoreFn := func(state ButtonPress) int { return fScore[state.dest] + 1 }
	openSet := pq.NewQueue[int, ButtonPress](fScoreFn, ButtonPress{dest: emptyState, button: -1})

	cameFrom := map[ButtonPress]ButtonPress{}
	gScore := map[string]int{emptyState: 0}

	for {
		if openSet.Len() == 0 {
			panic("Shouldn't happen")
		}

		current := openSet.Peek()

		if current.dest == m.lights {
			// Walk path
			return append(m.walkPath(cameFrom, current), current.button)
		} else {
			openSet.Poll()

			// For each neighbor of current
			buttonPresses := []ButtonPress{}
			for buttonIdx, button := range m.buttons {
				nextState := pressButton(current.dest, button)
				buttonPresses = append(buttonPresses, ButtonPress{dest: nextState, button: buttonIdx})
			}
			for _, buttonPress := range buttonPresses {
				state := buttonPress.dest
				g := gScore[current.dest] + 1
				gState, found := gScore[state]
				if !found || g < gState {
					fScore[state] = g + 1
					cameFrom[buttonPress] = current
					gScore[state] = g
					openSet.Append(buttonPress)
				}
			}
		}

	}

	panic("Shouldn't happen")
}

func stateToIntsPart2(state string) []int {
	nums := strings.Split(state[1:len(state)-1], " ")
	ints := make([]int, len(nums))
	for i, s := range nums {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func pressButtonPart2(state string, button []int) string {
	ints := stateToIntsPart2(state)
	for _, light := range button {
		ints[light] += 1
	}
	return fmt.Sprint(ints)
}

func compareJoltage(lhs, rhs string) int {
	lhsInts := stateToIntsPart2(lhs)
	rhsInts := stateToIntsPart2(rhs)
	return slices.Compare(lhsInts, rhsInts)
}

func solvePart2(m *Machine) []int {
	emptyState := fmt.Sprint(slices.Repeat([]int{0}, len(m.lights)))
	fScore := map[string]int{emptyState: math.MaxInt}
	fScoreFn := func(state ButtonPress) int { return fScore[state.dest] + 1 }
	openSet := pq.NewQueue[int, ButtonPress](fScoreFn, ButtonPress{dest: emptyState, button: -1})

	cameFrom := map[ButtonPress]ButtonPress{}
	gScore := map[string]int{emptyState: 0}

	i := 0
	for {
		if openSet.Len() == 0 {
			panic("Shouldn't happen")
		}

		current := openSet.Peek()

		if i%10000 == 0 {
			fmt.Println(i)
		}
		if current.dest == m.joltage {
			// Walk path
			return append(m.walkPath(cameFrom, current), current.button)
		} else {
			openSet.Poll()

			// For each neighbor of current
			buttonPresses := []ButtonPress{}
			for buttonIdx, button := range m.buttons {
				nextState := pressButtonPart2(current.dest, button)
				if compareJoltage(nextState, m.joltage) <= 0 {
					buttonPresses = append(buttonPresses, ButtonPress{dest: nextState, button: buttonIdx})
				}
			}
			slices.SortFunc(buttonPresses, func(a, b ButtonPress) int {
				return -compareJoltage(a.dest, b.dest)
			})
			for _, buttonPress := range buttonPresses {
				state := buttonPress.dest
				g := gScore[current.dest] + 1
				gState, found := gScore[state]
				if !found || g < gState {
					fScore[state] = g + 1
					cameFrom[buttonPress] = current
					gScore[state] = g
					openSet.Append(buttonPress)
				}
			}
			i++
		}

	}

	panic("Shouldn't happen")
}

func main() {
	lines := strings.Split(input, "\n")

	machines := make([]Machine, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		lights := parts[0][1 : len(parts[0])-1]
		buttons := parseButtons(parts[1 : len(parts)-1])
		joltageRaw := parts[len(parts)-1]
		joltage := "[" + strings.ReplaceAll(joltageRaw[1:len(joltageRaw)-1], ",", " ") + "]"

		machines = append(machines, Machine{lights: lights, buttons: buttons, joltage: joltage})
	}
	//fmt.Println(machines)

	part1 := 0
	for _, m := range machines {
		part1 += len(solvePart1(&m))
	}
	fmt.Println("Part1:", part1)

	part2 := 0
	for _, m := range machines {
		fmt.Println(m)
		part2 += len(solvePart2(&m))
	}
	fmt.Println("Part2:", part2)
}

package main

import (
	_ "embed"
	"fmt"
	"github.com/baritonehands/aoc-2025-go/utils"
	"slices"
	"strings"
)

//go:embed input.txt
var input string //= "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@."

func main() {
	lines := strings.Split(input, "\n")
	xMax := len(lines[0]) - 1
	yMax := len(lines) - 1

	part1 := 0
	part2 := 0
	curGen := lines
	nextGen := make([]string, len(curGen))
	for {
		for row := 0; row <= yMax; row++ {
			curRow := curGen[row]
			nextRow := make([]byte, len(curRow))
			for col := 0; col <= xMax; col++ {
				curPoint := utils.Point{col, row}
				if curGen[row][col] == '@' {
					adj := 0
					neighbors := curPoint.AllNeighbors(xMax, yMax)
					for _, p := range neighbors {
						if curGen[p.Y][p.X] == '@' {
							adj++
						}
					}
					if adj < 4 {
						part1++
						nextRow[col] = '.'
					} else {
						nextRow[col] = '@'
					}
				} else {
					nextRow[col] = '.'
				}
			}
			nextGen[row] = string(nextRow)
		}
		fmt.Println("Part1", part1)

		part2 += part1

		if slices.Equal(curGen, nextGen) {
			break
		}
		curGen = nextGen
		nextGen = make([]string, len(curGen))
	}

}

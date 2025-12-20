package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/baritonehands/aoc-2025-go/utils"
)

// go:embed input.txt
var input string = "0:\n###\n##.\n##.\n\n1:\n###\n##.\n.##\n\n2:\n.##\n###\n##.\n\n3:\n##.\n###\n##.\n\n4:\n###\n#..\n###\n\n5:\n###\n.#.\n###\n\n4x4: 0 0 0 0 2 0\n12x5: 1 0 1 0 2 2\n12x5: 1 0 1 0 3 2"

var rotations = map[utils.Point][]utils.Point{
	{X: 0, Y: 0}: {{2, 0}, {2, 2}, {0, 2}},
	{X: 1, Y: 0}: {{2, 1}, {1, 2}, {0, 1}},
	{X: 2, Y: 0}: {{2, 2}, {0, 2}, {0, 0}},
	{X: 0, Y: 1}: {{1, 0}, {2, 1}, {1, 2}},
	{X: 1, Y: 1}: {{1, 1}, {1, 1}, {1, 1}},
	{X: 2, Y: 1}: {{1, 2}, {0, 1}, {1, 0}},
	{X: 0, Y: 2}: {{0, 0}, {2, 0}, {2, 2}},
	{X: 1, Y: 2}: {{0, 1}, {1, 0}, {2, 1}},
	{X: 2, Y: 2}: {{0, 2}, {0, 0}, {2, 0}},
}

type Shape utils.Set[utils.Point]

func (s Shape) String() string {
	ret := []string{}
	for row := range 3 {
		rowStr := []byte{'.', '.', '.'}
		for col := range 3 {
			if s[utils.Point{col, row}] {
				rowStr[col] = '#'
			}
		}
		ret = append(ret, string(rowStr))
	}
	return strings.Join(ret, "\n")
}

func parseShape(input []string) []Shape {
	shape := Shape{}
	for ri, row := range input {
		for ci, col := range row {
			if col == '#' {
				shape[utils.Point{X: ci, Y: ri}] = true
			}
		}
	}
	fmt.Println(shape.String())
	fmt.Println()
	ret := []Shape{shape}
	seen := utils.Set[string]{shape.String(): true}
	for rot := range 3 {
		rotShape := Shape{}
		for point := range shape {
			rotShape[rotations[point][rot]] = true
		}
		shapeStr := rotShape.String()
		if !seen.Contains(shapeStr) {
			fmt.Println(shapeStr)
			fmt.Println()
			seen[shapeStr] = true
			ret = append(ret, rotShape)
		} else {
			fmt.Println("skipping duplicate shape")
		}
	}
	return ret
}

type Region struct {
	w, h     int
	presents []int
}

func main() {
	lines := strings.Split(input, "\n")

	shapes := [][]Shape{}
	regions := []Region{}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			continue
		}

		if line[1] == ':' {
			shapes = append(shapes, parseShape(lines[i+1:i+4]))
			i += 3
		} else {
			parts := strings.Split(line, ": ")
			dimParts := strings.Split(parts[0], "x")
			w, _ := strconv.Atoi(dimParts[0])
			h, _ := strconv.Atoi(dimParts[1])
			region := Region{w: w, h: h, presents: make([]int, len(shapes))}

			presents := strings.Split(parts[1], " ")
			for pi, p := range presents {
				region.presents[pi], _ = strconv.Atoi(p)
			}
			regions = append(regions, region)
		}
	}
	//fmt.Println(shapes)
	//fmt.Println(regions)

	//for _, region := range regions {
	//	taken := utils.Set[utils.Point]{}
	//	remainingPresents := slices.Clone(region.presents)
	//
	//	cur := utils.Point{X: 0, Y: 0}
	//	for {
	//
	//	}
	//}
}

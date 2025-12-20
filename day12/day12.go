package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/BooleanCat/go-functional/v2/it"
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

type Attempt struct {
	cur   utils.Point
	shape Shape
	idx   int
}

func printRegion(region Region, attempts []Attempt) {
	ret := [][]byte{}
	for _ = range region.h {
		rowBytes := make([]byte, region.w)
		for col := range region.w {
			rowBytes[col] = '.'
		}
		ret = append(ret, rowBytes)
	}
	for idx, attempt := range attempts {
		for point := range attempt.shape {
			ret[point.Y][point.X] = byte('A' + idx)
		}
	}
	fmt.Println(strings.Join(slices.Collect(it.Map(slices.Values(ret), func(bytes []byte) string {
		return string(bytes)
	})), "\n"))
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
			fmt.Println(len(shapes))
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

	part1 := 0
	for _, region := range regions[1:2] {
		taken := utils.Set[utils.Point]{}
		takenAttempts := []Attempt{}
		emptyPresents := make([]int, len(region.presents))
		remainingPresents := slices.Clone(region.presents)

		computeShapesToTry := func() [][]Shape {
			ret := make([][]Shape, len(region.presents))
			for i, n := range remainingPresents {
				if n > 0 {
					ret[i] = shapes[i]
				} else {
					ret[i] = []Shape{}
				}
			}
			return ret
		}

		cur := utils.Point{X: 0, Y: 0}
		shapesToTry := map[utils.Point][][]Shape{}

		findNextShape := func(cur utils.Point) (bool, int, Shape) {
			toTry := shapesToTry[cur]
			for i, try := range toTry {
				if len(try) > 0 {
					toTry[i] = toTry[i][1:]
					fmt.Printf("trying %v:\n%v\n", cur, try[0].String())
					return true, i, try[0]
				}
			}
			return false, 0, nil
		}

		computeNextPoint := func(cur utils.Point) (bool, utils.Point) {
			for row := cur.Y; row < region.h-2; row++ {
				xStart := cur.X + 1
				if row != cur.Y {
					xStart = 0
				}
				for col := xStart; col < region.w-2; col++ {
					point := utils.Point{X: col, Y: row}
					if !taken[point] {
						return true, point
					}
				}
			}
			return false, cur
		}

		backtrack := func() {
			lastAttempt := takenAttempts[len(takenAttempts)-1]
			takenAttempts = takenAttempts[:len(takenAttempts)-1]
			for point := range lastAttempt.shape {
				delete(taken, point)
			}
			delete(shapesToTry, cur)
			cur = lastAttempt.cur
			remainingPresents[lastAttempt.idx]++
		}

		for _ = range 1000 {
			if shapesToTry[cur] == nil {
				shapesToTry[cur] = computeShapesToTry()
			}

			found, shapeIdx, shapeToTry := findNextShape(cur)
			if found {
				valid := true
				adjustedShape := Shape{}
				for point := range shapeToTry {
					point.X += cur.X
					point.Y += cur.Y
					if (point.X >= region.w || point.Y >= region.h) || taken.Contains(point) {
						valid = false
						break
					}
					adjustedShape[point] = true
				}

				if valid {
					remainingPresents[shapeIdx]--
					takenAttempts = append(takenAttempts, Attempt{cur, adjustedShape, shapeIdx})
					if slices.Equal(emptyPresents, remainingPresents) {
						printRegion(region, takenAttempts)
						part1++
						break
					}
					for point := range adjustedShape {
						taken[point] = true
					}
					if found, cur = computeNextPoint(cur); !found {
						fmt.Println("backtrack: next point", len(taken))
						printRegion(region, takenAttempts)
						backtrack()
					}
				} else {
					if found, cur = computeNextPoint(cur); !found {
						fmt.Println("invalid: next shape", len(taken))
						printRegion(region, takenAttempts)
					}
				}
			} else {
				fmt.Println("backtrack: next shape", len(taken))
				printRegion(region, takenAttempts)
				backtrack()
			}
		}
		fmt.Println("part1", part1)
	}
}

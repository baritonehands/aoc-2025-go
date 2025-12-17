package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/baritonehands/aoc-2025-go/utils"
	"maps"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string //= "162,817,812\n57,618,57\n906,360,560\n592,479,940\n352,342,300\n466,668,158\n542,29,236\n431,825,988\n739,650,466\n52,470,668\n216,146,977\n819,987,18\n117,168,530\n805,96,715\n346,949,466\n970,615,88\n941,993,340\n862,61,35\n984,92,344\n425,690,689"

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(len(lines))

	points := it.Fold(slices.Values(lines), func(ret []utils.Point3D, s string) []utils.Point3D {
		parts := strings.Split(s, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		return append(ret, utils.Point3D{x, y, z})
	}, []utils.Point3D{})

	pairs := slices.Collect(utils.Permutations[utils.Point3D](points))

	slices.SortFunc(pairs, func(a, b utils.Pair[utils.Point3D]) int {
		return cmp.Compare(a.P1.DistanceTo(a.P2), b.P1.DistanceTo(b.P2))
	})

	circuits := map[utils.Point3D]utils.Set[utils.Point3D]{}
	processPair := func(pair utils.Pair[utils.Point3D]) {
		c1, ok1 := circuits[pair.P1]
		c2, ok2 := circuits[pair.P2]

		if !ok1 && !ok2 {
			newCircuit := utils.Set[utils.Point3D]{pair.P1: true, pair.P2: true}
			circuits[pair.P1] = newCircuit
			circuits[pair.P2] = newCircuit
			c1 = newCircuit
			c2 = newCircuit
		} else if ok1 && !ok2 {
			c1[pair.P2] = true
			c2 = c1
			circuits[pair.P2] = c2
		} else if ok2 && !ok1 {
			c2[pair.P1] = true
			c1 = c2
			circuits[pair.P1] = c1
		} else {
			maps.Insert(c1, maps.All(c2))
			for point := range c1 {
				circuits[point] = c1
			}
		}
	}
	for i := range 1000 {
		pair := pairs[i]
		processPair(pair)
	}

	part1 := map[string]int{}
	for _, circuit := range circuits {
		sortedPoints := slices.SortedFunc(maps.Keys(circuit), func(d utils.Point3D, d2 utils.Point3D) int {
			return d.Compare(d2)
		})
		part1[fmt.Sprint(sortedPoints)] = len(sortedPoints)
	}
	threeLargest := slices.SortedFunc(maps.Values(part1), func(i int, i2 int) int {
		return -cmp.Compare(i, i2)
	})[:3]

	part1Product := 1
	for _, cnt := range threeLargest {
		part1Product *= cnt
	}
	fmt.Println("part1", part1Product)

	circuits = map[utils.Point3D]utils.Set[utils.Point3D]{}
	for i := 0; i < len(pairs); i++ {
		pair := pairs[i]
		processPair(pair)

		if len(circuits) == len(points) {
			fmt.Println("part2", pair.P1.X*pair.P2.X)
			break
		}
	}
}

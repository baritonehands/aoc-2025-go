package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/baritonehands/aoc-2025-go/utils"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Square utils.Pair[utils.Point]

func (sq Square) String() string {
	return fmt.Sprintf("{%v, %v <-> %d}", sq.P1, sq.P2, sq.Area())
}

func (sq Square) Area() int {
	if sq.P1 == sq.P2 {
		return 1
	}

	if sq.P1.X > sq.P2.X {
		return Square{sq.P2, sq.P1}.Area()
	}

	return (sq.P2.X - sq.P1.X + 1) * (sq.P2.Y - sq.P1.Y + 1)
}

func main() {
	lines := strings.Split(input, "\n")

	points := []utils.Point{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, utils.Point{x, y})
	}

	pairs := slices.Collect(it.Map(utils.Permutations(points), func(pair utils.Pair[utils.Point]) Square {
		return Square(pair)
	}))
	slices.SortFunc(pairs, func(a, b Square) int {
		return cmp.Compare(a.Area(), b.Area())
	})
	fmt.Println("part1", pairs[len(pairs)-1:])
}

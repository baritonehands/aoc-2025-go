package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string //= "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32"

type Range struct {
	start, end int64
}

func (r Range) InRange(n int64) bool {
	return n >= r.start && n <= r.end
}

func (r Range) Size() int64 {
	return r.end - r.start + 1
}

func (r Range) Combine(other Range) (bool, Range) {
	if r == other {
		return true, r
	}

	if other.start >= r.start && other.start <= r.end ||
		other.end >= r.start && other.end <= r.end {
		combinedStart := int64(math.Min(float64(r.start), float64(other.start)))
		combinedEnd := int64(math.Max(float64(other.end), float64(r.end)))
		return true, Range{combinedStart, combinedEnd}
	}
	return false, Range{}
}

type Part2Result struct {
	output   []Range
	combined bool
	last     *Range
}

func main() {
	lines := strings.Split(input, "\n")

	ranges := make([]Range, 0)
	tests := make([]int64, 0)

	cur := 0
	var line string
	for cur, line = range lines {
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end, _ := strconv.ParseInt(parts[1], 10, 64)
		ranges = append(ranges, Range{start, end})
	}

	for _, testStr := range lines[cur+1:] {
		test, _ := strconv.ParseInt(testStr, 10, 64)
		tests = append(tests, test)
	}

	part1 := 0
	for _, test := range tests {
		for _, r := range ranges {
			if r.InRange(test) {
				part1++
				break
			}
		}
	}
	fmt.Println("part1", part1)

	part2 := int64(0)

	// Sort to make combination faster
	slices.SortFunc(ranges, func(lhs, rhs Range) int {
		return cmp.Or(cmp.Compare(lhs.start, rhs.start), cmp.Compare(lhs.end, rhs.end))
	})

	// Combine overlapping ranges
	curResult := Part2Result{output: ranges}
	fmt.Println("current size", len(ranges))
	for {
		nextResult := it.Fold(slices.Values(curResult.output), func(res Part2Result, r Range) Part2Result {
			// First iteration
			if res.last == nil {
				res.last = &r
				return res
			}

			combined, output := res.last.Combine(r)
			if combined {
				// Overwrite last range as combined, and mark as combined
				res.last = &output
				res.combined = true
			} else {
				// Append last range and continue to next
				res.output = append(res.output, *res.last)
				res.last = &r
			}
			return res
		}, Part2Result{output: make([]Range, 0)})
		fmt.Println("current size", len(nextResult.output))

		nextResult.output = append(nextResult.output, *nextResult.last)
		if !nextResult.combined {
			// Stop when we didn't do any combinations
			curResult = nextResult
			break
		}

		curResult = Part2Result{output: nextResult.output}
	}

	// Compute size from combined ranges
	for _, r := range curResult.output {
		part2 += r.Size()
	}
	fmt.Println("part2", part2)
}

package main

import (
	_ "embed"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"

	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/baritonehands/aoc-2025-go/utils"
)

//go:embed input.txt
var input string //= "aaa: you hhh\nyou: bbb ccc\nbbb: ddd eee\nccc: ddd eee fff\nddd: ggg\neee: out\nfff: out\nggg: out\nhhh: ccc fff iii\niii: out"

type PathSegment struct {
	path []string
	seen utils.Set[string]
}

func (ps *PathSegment) Downstream(connections map[string][]string) iter.Seq[string] {
	last := ps.path[len(ps.path)-1]

	return it.Filter(slices.Values(connections[last]), func(s string) bool {
		return !ps.seen.Contains(s)
	})
}

func main() {
	lines := strings.Split(input, "\n")

	connections := map[string][]string{}
	for _, l := range lines {
		parts := strings.Split(l, ": ")
		outputs := strings.Split(parts[1], " ")

		connections[parts[0]] = outputs
	}
	fmt.Println(connections)

	part1 := 0
	segments := []PathSegment{{[]string{"you"}, utils.Set[string]{"you": true}}}
	for {
		if len(segments) == 0 {
			break
		}

		segment := segments[0]
		segments = segments[1:]

		for downstream := range segment.Downstream(connections) {
			if downstream == "out" {
				part1++
				continue
			}

			nextSegment := PathSegment{path: slices.Clone(segment.path), seen: maps.Clone(segment.seen)}
			nextSegment.path = append(nextSegment.path, downstream)
			nextSegment.seen[downstream] = true

			segments = append(segments, nextSegment)
		}
	}
	fmt.Println("Part1: ", part1)
}

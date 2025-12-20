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
var input string //= "svr: aaa bbb\naaa: fft\nfft: ccc\nbbb: tty\ntty: ccc\nccc: ddd eee\nddd: hub\nhub: fff\neee: dac\ndac: fff\nfff: ggg hhh\nggg: out\nhhh: out"

type PathSegment struct {
	path []string
	seen utils.Set[string]
}

func (ps *PathSegment) Downstream(connections map[string]utils.Set[string]) iter.Seq[string] {
	last := ps.path[len(ps.path)-1]

	return it.Filter(maps.Keys(connections[last]), func(s string) bool {
		return !ps.seen.Contains(s)
	})
}

func main() {
	lines := strings.Split(input, "\n")

	connections := map[string]utils.Set[string]{}
	for _, l := range lines {
		parts := strings.Split(l, ": ")
		outputs := strings.Split(parts[1], " ")

		connections[parts[0]] = utils.SeqSet(slices.Values(outputs))
	}
	//fmt.Println(connections)

	findAllPaths := func(start string, doneFn func(*PathSegment, string) bool) {
		segments := []PathSegment{{[]string{start}, utils.Set[string]{start: true}}}
		for {
			if len(segments) == 0 {
				break
			}

			segment := segments[0]
			segments = segments[1:]

			for downstream := range segment.Downstream(connections) {
				if doneFn(&segment, downstream) {
					//fmt.Println(len(segments))
					continue
				}

				nextSegment := PathSegment{path: slices.Clone(segment.path), seen: maps.Clone(segment.seen)}
				nextSegment.path = append(nextSegment.path, downstream)
				nextSegment.seen[downstream] = true

				segments = append(segments, nextSegment)
			}
		}
	}

	part1 := 0
	part1DoneFn := func(_ *PathSegment, downstream string) bool {
		if downstream == "out" {
			part1++
			return true
		}
		return false
	}
	findAllPaths("you", part1DoneFn)
	fmt.Println("Part1: ", part1)

	graph := utils.NewDirectedGraph[string]()
	for k, v := range connections {
		graph.AddVertex(k)
		for vv := range v {
			graph.AddVertex(vv)
			graph.AddEdge(k, vv)
		}
	}

	fmt.Println(graph.Vertices["svr"])
	fmt.Println(graph.Vertices["fft"])
	fmt.Println(graph.Vertices["dac"])
	fmt.Println(graph.Vertices["out"])
}

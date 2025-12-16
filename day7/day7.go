package main

import (
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/baritonehands/aoc-2025-go/utils"
	"maps"
	"regexp"
	"slices"
	"strings"
)

//go:embed input.txt
var input string //= ".......S.......\n...............\n.......^.......\n...............\n......^.^......\n...............\n.....^.^.^.....\n...............\n....^.^...^....\n...............\n...^.^...^.^...\n...............\n..^...^.....^..\n...............\n.^.^.^.^.^...^.\n..............."

func main() {
	lines := strings.Split(input, "\n")

	start := strings.Index(lines[0], "S")
	xMax := len(lines[0])

	splitRegex := regexp.MustCompile(`\^`)
	beamSet := utils.Set[int]{start: true}
	part1 := 0
	for _, line := range lines[1:] {
		splits := splitRegex.FindAllStringIndex(line, -1)
		if len(splits) > 0 {
			nextBeamSet := utils.Set[int]{}
			splitSet := utils.SeqSet(it.Map(slices.Values(splits), func(split []int) int {
				return split[0]
			}))

			for beam := range beamSet {
				if splitSet.Contains(beam) {
					part1++
					if beam > 0 {
						nextBeamSet[beam-1] = true
					}
					if beam <= xMax && splitSet.Contains(beam) {
						nextBeamSet[beam+1] = true
					}
				} else {
					nextBeamSet[beam] = true
				}
			}

			beamSet = nextBeamSet
		}
	}
	fmt.Println("part1", part1)

	part2 := map[int]int64{start: 1}
	toAdd := map[int]int64{}
	toDelete := utils.Set[int]{}
	fmt.Println(part2)
	for _, line := range lines[1:] {
		splits := splitRegex.FindAllStringIndex(line, -1)
		if len(splits) > 0 {
			splitSet := utils.SeqSet(it.Map(slices.Values(splits), func(split []int) int {
				return split[0]
			}))

			clear(toAdd)
			clear(toDelete)
			for last, cnt := range part2 {
				if splitSet.Contains(last) {
					if last > 0 {
						if _, ok := toAdd[last-1]; !ok {
							toAdd[last-1] = 0
						}
						toAdd[last-1] += cnt
					}
					if last <= xMax {
						if _, ok := toAdd[last+1]; !ok {
							toAdd[last+1] = 0
						}
						toAdd[last+1] += cnt
					}
					toDelete[last] = true
				}
			}
			for last, cnt := range toAdd {
				part2[last] += cnt
			}
			for last := range toDelete {
				delete(part2, last)
			}
			//fmt.Println(part2)
		}
	}
	part2Sum := int64(0)
	for v := range maps.Values(part2) {
		part2Sum += v
	}
	fmt.Println("part2", part2Sum)
}

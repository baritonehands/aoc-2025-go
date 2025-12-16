package main

import (
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/baritonehands/aoc-2025-go/utils"
	"regexp"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

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
			fmt.Println(nextBeamSet.Values())
		}
	}
	fmt.Println("part1", part1)
}

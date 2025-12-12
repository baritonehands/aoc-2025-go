package main

import (
	_ "embed"
	"fmt"
	"github.com/baritonehands/aoc-2025-go/utils"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string // = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

func parseLine(line string) (int, int) {
	parts := strings.Split(line, "-")
	lhs, _ := strconv.ParseInt(parts[0], 10, 64)
	rhs, _ := strconv.ParseInt(parts[1], 10, 64)
	return int(lhs), int(rhs)
}

func isInvalid(n int) bool {
	str := fmt.Sprint(n)
	if len(str)%2 != 0 {
		return false
	}
	if str[0] == '0' {
		return false
	}

	p := len(str) / 2
	parts := slices.Collect(utils.Partition([]byte(str), p, p))
	return slices.Equal(parts[0], parts[1])
}

func isInvalidPart2(n int) bool {
	str := fmt.Sprint(n)
	if str[0] == '0' {
		return false
	}

	for p := 1; p <= len(str)/2; p++ {
		if len(str)%p == 0 {
			parts := slices.Collect(utils.Partition([]byte(str), p, p))
			first := parts[0]
			cnt := 1
			for _, v := range parts[1:] {
				if slices.Equal(v, first) {
					cnt++
				}
			}
			if cnt == len(parts) {
				return true
			}
		}
	}

	return false
}

func main() {
	lines := strings.Split(input, ",")
	fmt.Println(lines)

	part1 := 0
	part2 := 0
	for _, line := range lines {
		start, end := parseLine(line)
		end++
		for n := start; n <= end; n++ {
			if isInvalid(n) {
				part1 += n
			}
			if isInvalidPart2(n) {
				fmt.Println(n, isInvalidPart2(n))
				part2 += n
			}
		}
	}
	fmt.Println("Part1", part1)
	fmt.Println("Part2", part2)
}

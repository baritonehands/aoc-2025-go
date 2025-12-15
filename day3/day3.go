package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string //= "987654321111111\n811111111111119\n234234234234278\n818181911112111"

func largestJoltage(vals []int) int {
	tens := slices.Max(vals[:len(vals)-1])
	tensIdx := slices.Index(vals, tens)
	ones := slices.Max(vals[tensIdx+1:])
	return tens*10 + ones
}

func largestJoltage2(vals []int, size int) int {
	first := slices.Max(vals[:len(vals)-size-1])
	firstIdx := slices.Index(vals, first)
	ones := slices.Max(vals[firstIdx+1:])
	return first*10 + ones
}

func digits(s string) []int {
	ret := make([]int, 0, len(s))
	for _, c := range s {
		ret = append(ret, int(c-'0'))
	}
	return ret
}

func main() {
	lines := strings.Split(input, "\n")
	//fmt.Println(lines)

	part1 := 0
	for _, line := range lines {
		digitSlice := digits(line)
		part1 += largestJoltage(digitSlice)
		fmt.Println(largestJoltage2(digitSlice, 10))
	}
	fmt.Println("Part1", part1)
}

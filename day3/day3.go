package main

import (
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/baritonehands/aoc-2025-go/utils"
	"slices"
	"strings"
)

//go:embed input.txt
var input string //= "987654321111111\n811111111111119\n234234234234278\n818181911112111"

func largestJoltage(line string, vals []int) int {
	tens := slices.Max(vals[:len(vals)-1])
	tensIdx := strings.Index(line, fmt.Sprint(tens))
	ones := slices.Max([]byte(line)[tensIdx+1:])
	return tens*10 + int(ones-'0')
}

func main() {
	lines := strings.Split(input, "\n")
	//fmt.Println(lines)

	part1 := 0
	for _, line := range lines {
		distinctSeq := utils.PartitionFunc2([]byte(line), utils.Identity)
		distinctSlice := make([]int, 0, it.Len(distinctSeq))
		for seq := range distinctSeq {
			_, vals := it.Collect2(seq)
			distinctSlice = append(distinctSlice, int(vals[0]-'0'))
		}
		//fmt.Println(distinctSlice)
		part1 += largestJoltage(line, distinctSlice)
	}
	fmt.Println("Part1", part1)
}

package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string //= "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"

func parseLine(line string) (string, int) {
	n, _ := strconv.ParseInt(line[1:], 10, 64)
	return string(line[0]), int(n)
}

func main() {
	lines := strings.Split(input, "\n")

	pos := 50
	part1 := 0
	part2 := 0
	for _, l := range lines {
		dir, n := parseLine(l)

		zeros := 0
		var nextNumber int
		if dir == "L" {
			if pos+n >= 100 {
				zeros = n / 100
				if pos+(n%100) >= 100 {
					zeros++
				}
			}
			nextNumber = pos - n
		} else {
			if pos-n > 0 {
				zeros = n / 100
				if pos-(n%100) <= 0 {
					zeros++
				}
			}
			nextNumber = pos + n
		}

		pos = (nextNumber%100 + 100) % 100
		fmt.Println(pos)

		part2 += zeros
		if pos == 0 {
			part1++
		}
	}
	fmt.Println("Part1: ", part1)
	fmt.Println("Part2: ", part2)
}

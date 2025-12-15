package main

import (
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string //= "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  "

const height = 4

type Operation struct {
	operands []int
	operator string
}

func (o Operation) Compute() int {
	init := 0
	op := func(a, b int) int {
		return a + b
	}
	if o.operator == "*" {
		init = 1
		op = func(a, b int) int {
			return a * b
		}
	}
	return it.Fold(slices.Values(o.operands), op, init)
}

func column(lines []string, idx int) (int, *string) {
	if idx < 0 {
		return -1, nil
	}

	ret := make([]byte, 0, height)
	var op *string
	for row := range height {
		if lines[row][idx] != ' ' {
			ret = append(ret, lines[row][idx])
		}
	}
	if lines[height][idx] != ' ' {
		opBytes := string(lines[height][idx])
		op = &opBytes
	}

	if len(ret) == 0 {
		return -1, nil
	}

	retStr := string(ret)
	retInt, _ := strconv.Atoi(retStr)
	return retInt, op
}

func main() {
	lines := strings.Split(input, "\n")
	whitespace := regexp.MustCompile(`\s+`)

	strOps := [][]string{}
	for _, line := range lines {
		args := whitespace.Split(line, -1)
		strOps = append(strOps, args)
	}

	part1 := 0
	for opIdx := range strOps[0] {
		operands := make([]int, height)
		for i := range height {
			operands[i], _ = strconv.Atoi(strOps[i][opIdx])
		}
		operator := strOps[height][opIdx]
		op := Operation{operands: operands, operator: operator}
		part1 += op.Compute()
	}
	fmt.Println("part1", part1)

	part2 := 0
	operands := make([]int, 0, 10)
	var operator *string
	for idx := len(lines[0]) - 1; idx >= -1; idx-- {
		n, op := column(lines, idx)
		if n == -1 {
			operation := Operation{operands: operands, operator: *operator}
			fmt.Println(operation)
			part2 += operation.Compute()
			operator = nil
			operands = operands[:0]
			continue
		}

		if op != nil {
			operator = op
		}

		operands = append(operands, n)
	}
	fmt.Println("part2", part2)
}

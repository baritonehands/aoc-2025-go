package utils

import (
	"fmt"
	"iter"
)

type Pair[T any] struct {
	P1, P2 T
}

func (p Pair[T]) String() string {
	return fmt.Sprintf("{%v, %v}", p.P1, p.P2)
}

func Permutations[T any](input []T) iter.Seq[Pair[T]] {
	return func(yield func(Pair[T]) bool) {
		for i, lhs := range input {
			for j := i + 1; j < len(input); j++ {
				rhs := input[j]

				if !yield(Pair[T]{lhs, rhs}) {
					return
				}
			}
		}
	}
}

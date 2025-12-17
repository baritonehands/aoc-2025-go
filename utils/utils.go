package utils

import (
	"github.com/BooleanCat/go-functional/v2/it"
	"iter"
	"slices"
	"strings"
)

func Split2(s string) (string, string) {
	arr := strings.SplitN(s, " ", 2)
	return arr[0], arr[1]
}

func Frequencies[I iter.Seq[T], T comparable](iter I) map[T]int64 {
	return it.Fold(iter, func(m map[T]int64, t T) map[T]int64 {
		m[t]++
		return m
	}, make(map[T]int64))
}

func FlatMap[V, W any, S iter.Seq[W]](delegate func(func(V) bool), f func(V) S) iter.Seq[W] {
	return func(yield func(W) bool) {
		for innerValue := range delegate {
			for value := range f(innerValue) {
				if !yield(value) {
					return
				}
			}
		}
	}
}

func FlatMap2[V, W, X, Y any, S iter.Seq2[X, Y]](delegate func(func(V, W) bool), f func(V, W) S) iter.Seq2[X, Y] {
	return func(yield func(X, Y) bool) {
		for v, w := range delegate {
			for x, y := range f(v, w) {
				if !yield(x, y) {
					return
				}
			}
		}
	}
}

func Partition[T any, S ~[]T](slice S, n int, step int) iter.Seq[S] {
	if n == step {
		return slices.Chunk(slice, n)
	}

	return func(yield func(S) bool) {
		for i := 0; i < len(slice); i += step {
			innerLen := min(i+n, len(slice))
			inner := slice[i:innerLen:innerLen]
			if !yield(inner) {
				return
			}
		}
	}

}

func PartitionFunc2[V any, U comparable](slice []V, fn func(t V) U) iter.Seq[iter.Seq2[int, V]] {
	if len(slice) == 0 {
		return it.Exhausted[iter.Seq2[int, V]]()
	}

	return func(yield func(iter.Seq2[int, V]) bool) {
		last := fn(slice[0])
		lastI := 0
		partition := slice[0:1:1]

		endPartition := func() bool {
			indexes, values := it.Collect2(it.Map2(slices.All(partition), func(idx int, v V) (int, V) {
				return lastI + idx, v
			}))
			return yield(it.Zip(slices.Values(indexes), slices.Values(values)))
		}

		for i := 1; i < len(slice); i += 1 {
			byValue := fn(slice[i])
			if last == byValue {
				partition = append(partition, slice[i])
			} else {
				if !endPartition() {
					return
				}
				lastI = i
				partition = slice[i : i+1 : i+1]
			}
			last = byValue
		}
		endPartition()
		return
	}
}

func Identity[T any](t T) T {
	return t
}

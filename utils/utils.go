package utils

import (
	"cmp"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"iter"
	"maps"
	"math"
	"slices"
	"strings"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p Point) OrthogonalNeighbors(xMax, yMax int) []Point {
	ret := make([]Point, 0, 4)
	if p.X < xMax {
		ret = append(ret, Point{p.X + 1, p.Y})
	}
	if p.Y < yMax {
		ret = append(ret, Point{p.X, p.Y + 1})
	}
	if p.X > 0 {
		ret = append(ret, Point{p.X - 1, p.Y})
	}
	if p.Y > 0 {
		ret = append(ret, Point{p.X, p.Y - 1})
	}
	return ret
}

func (p Point) AllNeighbors(xMax, yMax int) []Point {
	ret := p.OrthogonalNeighbors(xMax, yMax)
	if p.X < xMax && p.Y < yMax {
		ret = append(ret, Point{p.X + 1, p.Y + 1})
	}
	if p.X > 0 && p.Y < yMax {
		ret = append(ret, Point{p.X - 1, p.Y + 1})
	}
	if p.X > 0 && p.Y > 0 {
		ret = append(ret, Point{p.X - 1, p.Y - 1})
	}
	if p.X < xMax && p.Y > 0 {
		ret = append(ret, Point{p.X + 1, p.Y - 1})
	}
	return ret
}

func PointCompareYX(p1 Point, p2 Point) int {
	y := cmp.Compare(p1.Y, p2.Y)
	if y != 0 {
		return y
	}
	x := cmp.Compare(p1.X, p2.X)
	return x
}

func PointCompareXY(p1 Point, p2 Point) int {
	x := cmp.Compare(p1.X, p2.X)
	if x != 0 {
		return x
	}
	y := cmp.Compare(p1.Y, p2.Y)
	return y
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}

func (p Point3D) Compare(other Point3D) int {
	return cmp.Or(cmp.Compare(p.X, other.X), cmp.Compare(p.Y, other.Y), cmp.Compare(p.Z, other.Z))
}

func (p Point3D) DistanceTo(other Point3D) float64 {
	return math.Sqrt(
		math.Pow(float64(other.X-p.X), 2) +
			math.Pow(float64(other.Y-p.Y), 2) +
			math.Pow(float64(other.Z-p.Z), 2))
}

func Split2(s string) (string, string) {
	arr := strings.SplitN(s, " ", 2)
	return arr[0], arr[1]
}

type Set[T comparable] map[T]bool

func SeqSet[I iter.Seq[T], T comparable](iter I) Set[T] {
	return maps.Collect(it.Zip(iter, it.Repeat(true)))
}

func (s Set[T]) Values() []T {
	return slices.Collect(maps.Keys(s))
}

func (s Set[T]) Contains(v T) bool {
	sVal, ok := s[v]
	return ok && sVal
}

func (lhs Set[T]) Union(rhs Set[T]) Set[T] {
	var ret = make(map[T]bool)
	for c, v := range lhs {
		ret[c] = v
	}
	for c, v := range rhs {
		ret[c] = v
	}
	return ret
}

func (lhs Set[T]) Difference(rhs Set[T]) Set[T] {
	var ret = make(map[T]bool)
	for c, v := range lhs {
		_, present := rhs[c]
		if v && !present {
			ret[c] = true
		}
	}
	return ret
}

func (lhs Set[T]) Intersection(rhs Set[T]) Set[T] {
	var ret = make(map[T]bool)
	for c, v := range lhs {
		_, present := rhs[c]
		if v && present {
			ret[c] = true
		}
	}
	return ret
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

	ret := it.Exhausted[S]()
	for i := 0; i < len(slice); i += step {
		innerLen := min(i+n, len(slice))
		inner := slice[i:innerLen:innerLen]
		ret = it.Chain(ret, it.Once(inner))
	}
	return ret
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

type Pair[T any] struct {
	P1, P2 T
}

func (p Pair[T]) String() string {
	return fmt.Sprintf("{%v, %v}", p.P1, p.P2)
}

func Permutations[T any](input []T) []Pair[T] {
	ret := make([]Pair[T], 0, len(input))
	for i, lhs := range input {
		for j := i + 1; j < len(input); j++ {
			rhs := input[j]

			ret = append(ret, Pair[T]{lhs, rhs})
		}
	}
	return ret
}

package utils

import (
	"github.com/BooleanCat/go-functional/v2/it"
	"iter"
	"maps"
	"slices"
)

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

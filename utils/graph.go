package utils

import (
	"maps"
	"slices"

	"github.com/BooleanCat/go-functional/v2/it"
)

type DirectedGraph[T comparable] struct {
	vertices      map[T][]T
	incomingEdges map[T][]T
}

func NewDirectedGraph[T comparable]() DirectedGraph[T] {
	return DirectedGraph[T]{vertices: make(map[T][]T), incomingEdges: make(map[T][]T)}
}

func (graph *DirectedGraph[T]) AddVertex(label T) {
	if graph.vertices[label] == nil {
		graph.vertices[label] = []T{}
	}
	if graph.incomingEdges[label] == nil {
		graph.incomingEdges[label] = []T{}
	}
}

func (graph *DirectedGraph[T]) AddEdge(label1, label2 T) {
	graph.vertices[label1] = append(graph.vertices[label1], label2)
	graph.incomingEdges[label2] = append(graph.incomingEdges[label2], label1)
}

func (graph *DirectedGraph[T]) RemoveEdge(label1, label2 T) {
	eV1 := graph.vertices[label1]
	eV2 := graph.incomingEdges[label2]
	if eV1 != nil {
		idx := slices.Index(eV1, label2)
		if idx != -1 {
			graph.vertices[label1] = append(eV1[:idx], eV1[idx+1:]...)
		}
	}
	if eV2 != nil {
		idx := slices.Index(eV2, label1)
		if idx != -1 {
			graph.incomingEdges[label2] = append(eV2[:idx], eV2[idx+1:]...)
		}
	}
}

func (graph *DirectedGraph[T]) TopologicalSort() []T {
	ret := []T{}
	s := []T{}

	for k, nodes := range graph.incomingEdges {
		if len(nodes) == 0 {
			s = append(s, k)
		}
	}

	for len(s) > 0 {
		n := s[0]
		s = s[1:]
		ret = append(ret, n)

		edges := slices.Clone(graph.vertices[n])
		for _, m := range edges {
			graph.RemoveEdge(n, m)

			if len(graph.incomingEdges[m]) == 0 {
				s = append(s, m)
			}
		}
	}

	allEmpty := it.All(it.Map(maps.Values(graph.incomingEdges), func(list []T) bool {
		return len(list) == 0
	}))

	if !allEmpty {
		panic("Graph has at least one allEmpty")
	}

	return ret
}

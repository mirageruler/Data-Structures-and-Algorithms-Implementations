package main

import (
	"container/heap"
	"fmt"
	"math"
)

// edge represents an edge in the graph
type edge struct {
	to     int // the id of the destination node
	weight int // the weight of the edge
}

// node represents a node in the graph
type node struct {
	id       int // node's id
	distance int // distance from the starting/source node to this node
}

// dijkstra takes a graph g represented as an adjacency list and a starting/source node as input, and returns an array of distances from the starting node to each node in the graph.
// dijsktra can not handle graph with negative weighted edges and/or negative cycles
// This implementation works for undirected graph.
func dijkstraWithMinHeap(g [][]edge, start int) ([]int, map[int][]int) {
	infinity := math.MaxInt32
	dists := make([]int, len(g))
	sps := map[int][]int{}

	// distance to each node is unknown at the beginning
	for i := range dists {
		dists[i] = infinity
	}
	dists[start] = 0
	sps[start] = []int{start}

	pq := make(priorityQueue, len(g))
	nodes := make([]*node, len(g))

	for i := range g {
		nodes[i] = &node{id: i, distance: infinity}
		if i == start {
			nodes[i].distance = 0
		}
		pq[i] = nodes[i]
	}

	heap.Init(&pq)
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*node)
		if u.distance > dists[u.id] {
			continue
		}

		// relaxing edges from the extracted vertex u
		for _, edge := range g[u.id] {
			if dists[edge.to] > u.distance+edge.weight {
				dists[edge.to] = u.distance + edge.weight
				pq.updateDistance(edge.to, dists[edge.to])
				sps[edge.to] = append(sps[u.id], edge.to)
			}
		}
	}

	return dists, sps
}

func main() {
	// undirected graph (if from x to y is w then from y to x is also w)
	graph := [][]edge{
		{{to: 1, weight: 4}, {to: 2, weight: 1}},
		{{to: 0, weight: 4}, {to: 2, weight: 2}, {to: 3, weight: 5}},
		{{to: 0, weight: 1}, {to: 1, weight: 2}, {to: 3, weight: 1}},
		{{to: 1, weight: 5}, {to: 2, weight: 1}},
	}

	start := 0
	dist, sps := dijkstraWithMinHeap(graph, start)
	for i, d := range dist {
		fmt.Printf("SHORTEST PATH WEIGHT FROM SOURCE %d TO NODE %d is %v\n", start, i, d)
	}
	for node, path := range sps {
		fmt.Printf("SHORTEST WEIGHTED PATH FROM SOURCE %d TO NODE %d is %v\n", start, node, path)
	}
}

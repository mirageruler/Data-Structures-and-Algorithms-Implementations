package main

import "fmt"

type Edge struct {
	from   int
	to     int
	weight int
}

// This implementation assumes that the graph is represented as an adjacency list, where each vertex has a list of its adjacent vertices and their corresponding edge weights.
// The function returns a slice of integers representing the shortest distance from the starting vertex to each other vertex in the graph, or an error if the graph contains a negative cycle.
// This implementation works for undirected graph.
func bellmanFord(edges []Edge, start int, numVertices int) ([]int, map[int][]int, error) {
	distances := make([]int, numVertices)
	for i := range distances {
		distances[i] = 1<<31 - 1 // set initial distances to infinity
	}
	distances[start] = 0 // set distance to starting vertex to 0
	sps := map[int][]int{}
	sps[start] = []int{start}

	for i := 0; i < numVertices-1; i++ {
		for _, edge := range edges {
			// relaxing edges
			if distances[edge.from]+edge.weight < distances[edge.to] {
				distances[edge.to] = distances[edge.from] + edge.weight
				sps[edge.to] = append(sps[edge.from], edge.to)
			}
		}
	}

	// check for negative cycles
	for _, edge := range edges {
		if distances[edge.from]+edge.weight < distances[edge.to] {
			return nil, nil, fmt.Errorf("graph contains a negative cycle")
		}
	}

	return distances, sps, nil
}

func main() {
	// undirected graph (if from x to y is w then from y to x is also w)
	edges := []Edge{
		// 1 vertex
		{0, 1, 5},
		{from: 1, to: 0, weight: 5},
		// 1 vertex
		{0, 2, 4},
		{2, 0, 4},
		// 1 vertex
		{1, 3, 3},
		{3, 1, 3},
		// 1 vertex
		{2, 1, 3},
		{1, 2, 3},
		// 1 vertex
		{2, 3, 2},
		{3, 2, 2},
		// 1 vertex, if we add a new edge for the opposite side, this will be considered as a negative cycle which is bad
		{3, 4, -2},
	}
	start := 0
	numVertices := 5

	distances, sps, err := bellmanFord(edges, start, numVertices)
	if err != nil {
		fmt.Println(err)
	} else {
		for i, d := range distances {
			fmt.Printf("SHORTEST PATH WEIGHT FROM SOURCE %d TO NODE %d is %v\n", start, i, d)
		}
	}
	for node, path := range sps {
		fmt.Printf("SHORTEST WEIGHTED PATH FROM SOURCE %d TO NODE %d is %v\n", start, node, path)
	}
}

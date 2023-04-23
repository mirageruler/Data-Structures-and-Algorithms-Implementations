package main

import (
	"fmt"
	"sync"
)

func floydWarshall(graph [][]float64) [][]float64 {
	n := len(graph)
	dist := make([][]float64, n)
	for i := range dist {
		dist[i] = make([]float64, n)
		copy(dist[i], graph[i])
	}

	var wg sync.WaitGroup
	for k := range graph {
		for i := range graph {
			for j := range graph {
				if i != j && i != k && j != k { // a kind of memoization
					wg.Add(1)
					go func(i, j, k int) { // concurrency
						defer wg.Done()
						if dist[i][k]+dist[k][j] < dist[i][j] {
							dist[i][j] = dist[i][k] + dist[k][j]
						}
					}(i, j, k)
				}
			}
		}
	}

	wg.Wait()

	return dist
}

func main() {
	graph := [][]float64{
		{0, 5, 6, 10},
		{5, 0, 3, 4},
		{6, 3, 0, 1},
		{10, 4, 1, 0},
	}

	dist := floydWarshall(graph)

	fmt.Println("Shortest distances between all pairs of vertices:")
	for _, row := range dist {
		fmt.Println(row)
	}
}

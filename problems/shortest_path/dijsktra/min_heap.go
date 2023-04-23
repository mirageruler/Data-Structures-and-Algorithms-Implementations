package main

import (
	"container/heap"
)

// Priority Queue implemented using a min heap
type priorityQueue []*node

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }

func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*node)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) updateDistance(id int, newDistance int) {
	// Find the index of the node in the priority queue
	for i := 0; i < len(*pq); i++ {
		if (*pq)[i].id == id {
			// Update the distance of the node
			(*pq)[i].distance = newDistance

			// Re-heapify the priority queue
			heap.Fix(pq, i)
			break
		}
	}
}

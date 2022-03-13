package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()

	g.AddVertex("0")
	g.AddVertex("1")
	g.AddVertex("2")
	g.AddVertex("3")
	g.AddVertex("4")
	g.AddVertex("5")
	g.AddVertex("6")
	g.AddEdge("3", "1")
	g.AddEdge("3", "4")
	g.AddEdge("4", "2")
	g.AddEdge("4", "5")
	g.AddEdge("1", "2")
	g.AddEdge("1", "0")
	g.AddEdge("0", "2")
	g.AddEdge("6", "5")

	// bs, _ := json.Marshal(g)
	// fmt.Println("graph: ", string(bs))
	fmt.Printf("graph: %v \n", g.AdjacencyList)

	//t.Error()
}

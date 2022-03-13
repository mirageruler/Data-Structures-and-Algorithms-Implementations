package graph

/*
					[2]--------[0]
				  /    \
				 /      \
				/        \
               [1]-------[3]




*EDGE LIST*: Simply says list out all the edges of the graph with the fact is that each edge connects 2 adjacency nodes.
graph := [[0, 2], [2, 3], [2, 1], [1, 3]]

*ADJACENCY LIST*: Simply says the index is the value of the node itself and the value is a collection of its neighbor nodes' value.
graph := [[2], [2,3], [0,1,3], [1,2]]
		   0     1       2       3

*ADJACENCY MATRIX*: Simply displays the connection between Node X and Node Y while '0' is no and '1' is yes. If the graph is weighted, we can replace number '1' with the weight number.
 graph := {
	0: [0, 0, 1, 0],			At index 0 of this graph array, we're telling that the Node with value '0' only has connection to the Node with value '2', because the number '1' is placed at index 2 of this sub-array while other places are placed with '0'.
	1: [0, 0, 1, 1],			Simlilarly, this tells Node with value '1' has connections to Node with value '2' and '3'
	2: [1, 1, 0, 1],			This tells Node with value '2' has connections to Node with value '0', '1', and '3'
	3: [0, 1, 1, 0]			This tells Node with value '3' has connections to Node with value '1', '2'
 }

NOTE: We can flexibly use arrays/hash tables or linked lists to represent graph, choose the one that perfectly fit in the situation/context.
*/

// Graph defines an undirected, unweighted, cyclic graph using Adjacency List a.k.a Golang map
type Graph struct {
	// Number of current nodes/vertexes of the graph
	Size int
	// The Adjacency List represents the graph
	AdjacencyList map[interface{}][]interface{}
}

func NewGraph() *Graph {
	return &Graph{AdjacencyList: map[interface{}][]interface{}{}}
}

func (g *Graph) AddVertex(node interface{}) interface{} {
	if g == nil {
		return false
	}

	if _, ok := g.AdjacencyList[node]; ok {
		return false
	}

	g.AdjacencyList[node] = []interface{}{}
	g.Size++

	return true
}

func (g *Graph) AddEdge(node1, node2 interface{}) interface{} {
	// undirected graph
	if g == nil {
		return false
	}

	g.AdjacencyList[node1] = append(g.AdjacencyList[node1], node2)
	g.AdjacencyList[node2] = append(g.AdjacencyList[node2], node1)

	return true
}

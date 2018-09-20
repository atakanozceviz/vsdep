package graph

// Graph base struct
type Graph struct {
	Nodes map[string]struct{}
	Edges map[string]map[string]struct{}
}

// NewGraph Create graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]struct{}),
		Edges: make(map[string]map[string]struct{}),
	}
}

// AddNode Add node id to graph, return true if added (string's are unique).
func (g *Graph) AddNode(id string) bool {
	if _, ok := g.Nodes[id]; ok {
		return false
	}
	g.Nodes[id] = struct{}{}
	return true
}

// AddEdge Add an edge from u to v.
func (g *Graph) AddEdge(u, v string) {
	if _, ok := g.Nodes[u]; !ok {
		g.AddNode(u)
	}
	if _, ok := g.Nodes[v]; !ok {
		g.AddNode(v)
	}

	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = make(map[string]struct{})
	}
	g.Edges[u][v] = struct{}{}
}

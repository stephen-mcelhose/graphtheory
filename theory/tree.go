package theory

import (
	"github.com/stephen-mcelhose/graphtheory/graph"
	"github.com/stephen-mcelhose/graphtheory/traversal"
)

func IsTree(g graph.Graph) bool {
	if g.Order() == 0 {
		return true
	}
	return traversal.IsConnected(g) && g.Size() == g.Order()-1
}

func IsForest(g graph.Graph) bool {
	comps := traversal.ConnectedComponents(g)
	return g.Size() == g.Order()-len(comps)
}

package mcts

import "math"

// UCT constant ~= sqrt(2)
// A lower value makes the AI explore more
// A greater value makes the AI focus on the winning path
const C = 1.41421356237

func calculateUCT[M any](node *Node[M]) float64 {
	winRate := node.Wins / float64(node.Visits)

	// The lower the visit count of the child relative to the parent, the higher this number is
	exploration := C * math.Sqrt(math.Log(float64(node.Parent.Visits))/float64(node.Visits))

	return winRate + exploration
}

func getMostVisitedChild[M any](node *Node[M]) *Node[M] {
	var bestChild *Node[M]
	maxVisits := -1

	for _, child := range node.Children {
		if child.Visits > maxVisits {
			maxVisits = child.Visits
			bestChild = child
		}
	}
	return bestChild
}

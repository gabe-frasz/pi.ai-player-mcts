package mcts

import "time"

func Search[M any](initialState State[M], maxDuration time.Duration) M {
	root := NewRootNode(initialState)
	timeout := time.After(maxDuration)

SearchLoop:
	for {
		select {
		case <-timeout:
			break SearchLoop
		default:
			// 1. Go forward in the three (if the root has untried moves, stay in the root)
			node := root
			for len(node.UntriedMoves) == 0 && len(node.Children) > 0 {
				node = node.SelectChild()
			}

			// 2. If the game is not over in this node, expand creating a child
			if len(node.UntriedMoves) > 0 && !node.State.IsTerminal() {
				node = node.Expand()
			}

			// 2. Simulate a game from the new child
			terminalState := node.Rollout()

			// 5. Backpropagate the result
			node.Backpropagate(terminalState)
		}
	}

	// MCTS rule: pick the node with the most Visits
	// This node is more resistant against the opponent's counter-moves.
	return getMostVisitedChild(root).Move
}

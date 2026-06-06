package main

import (
  "github.com/gabe-frasz/pi.ai-player-mcts/game"
  "time"
)

// FindBestMove é a função principal chamada pelo seu servidor/webhook.
// Recebe o estado inicial e quanto tempo ela tem para "pensar" (ex: 1500ms).
func FindBestMove(initialState game.State, maxDuration time.Duration) game.Move {
	root := game.NewRootNode(initialState)
	timeout := time.After(maxDuration)

	// O Loop Implacável do MCTS
SearchLoop:
	for {
		select {
		case <-timeout:
			// Acabou o tempo! Sai do loop.
			break SearchLoop
		default:
			// O Ciclo Padrão do MCTS:
			
			// 1. Desce na árvore (Se a raiz tem UntriedMoves, o node = root)
			node := root
			for len(node.UntriedMoves) == 0 && len(node.Children) > 0 {
				node = node.SelectChild()
			}

			// 2. Se o jogo não acabou nesse nó, expande criando um filho
			if len(node.UntriedMoves) > 0 {
				node = node.Expand()
			}

			// 3. Simula uma partida caótica a partir do novo filho
			result := node.Rollout()

			// 4. Sobe avisando o resultado
			node.Backpropagate(result)
		}
	}

	// O tempo acabou. Qual a melhor jogada?
	// A regra do MCTS diz: não pegue o nó com mais Wins, pegue o nó com mais VISITS.
	// O nó mais visitado provou ser resistente contra as contra-jogadas do inimigo.
	return GetMostVisitedChild(root).Move
}

// Função auxiliar para achar o filho mais visitado
func GetMostVisitedChild(node *game.Node) *game.Node {
	var bestChild *game.Node
	maxVisits := -1

	for _, child := range node.Children {
		if child.Visits > maxVisits {
			maxVisits = child.Visits
			bestChild = child
		}
	}
	return bestChild
}

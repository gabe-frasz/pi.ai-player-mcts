package game

import (
	"math"
	"math/rand"
	"time"
)

// Constante de exploração do UCT (geralmente raiz de 2, ou ~1.41)
// Um valor maior faz a IA explorar jogadas ruins só por curiosidade.
// Um valor menor faz ela focar logo na jogada que tá ganhando.
const explorationParam = 1.414

// O Nó da Árvore
type Node struct {
	State       State    // O tabuleiro neste momento da simulação
	Parent      *Node    // Ponteiro pro pai (pra poder subir atualizando vitórias)
	Children    []*Node  // Ponteiros pros filhos (próximas jogadas)
	Move        Move     // Qual jogada levou a este nó
	
	Visits      int      // N (Quantas vezes a simulação passou aqui)
	Wins        float64  // W (Quantas vitórias passaram por aqui. Float para lidar com empates 0.5 se houver)
	
	UntriedMoves []Move  // Jogadas que AINDA NÃO viraram filhos
}

// Cria um Nó Raiz (o início do turno)
func NewRootNode(state State) *Node {
	return &Node{
		State:        state,
		Parent:       nil,
		Children:     make([]*Node, 0),
		UntriedMoves: state.GetMoves(), // Pede pra física do jogo gerar as opções
	}
}

// Passo 1: SELEÇÃO
// Desce na árvore usando UCT até achar um Nó que tenha UntriedMoves ou seja fim de jogo.
func (n *Node) SelectChild() *Node {
	// A lógica do UCT virá aqui
	return nil
}

// Passo 2: EXPANSÃO
// Pega UMA jogada do UntriedMoves, cria o filho, adiciona na lista e retorna esse filho.
func (n *Node) Expand() *Node {
	// A lógica de criação de filho virá aqui
	return nil
}

// Passo 3: ROLLOUT (A simulação cega)
// A partir do estado deste nó, joga movimentos aleatórios até o jogo acabar.
// Retorna quem ganhou (ResultWin, ResultLoss).
func (n *Node) Rollout() int8 {
	// Loop de sorteio de movimentos virá aqui
	return 0
}

// Passo 4: BACKPROPAGATION
// Recebe o resultado do Rollout e sobe pela árvore atualizando Visits e Wins.
func (n *Node) Backpropagate(result int8) {
	// Loop subindo pelos Parents virá aqui
}

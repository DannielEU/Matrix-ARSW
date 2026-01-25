package entity

import "matrix-game/internal/board"

type Entity interface {
	Icon() rune
	Run(b *board.Board, moveCh chan<- Move)
}

type Move struct {
	OldX, OldY int
	NewX, NewY int
	Icon       rune
}

// Canal de fin de juego
type Done struct {
	Winner string // "Protagonista" o "Agente"
}

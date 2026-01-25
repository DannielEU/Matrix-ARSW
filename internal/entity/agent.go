package entity

import (
	"matrix-game/internal/board"
	"time"
)

type Agent struct {
	X, Y int
	ID   int
}

func (a *Agent) Icon() rune {
	return 'A'
}

func (a *Agent) Run(b *board.Board, moveCh chan<- Move, doneCh chan<- Done, target func() (int, int)) {
	for {
		time.Sleep(500 * time.Millisecond)

		tx, ty := target() // posición del protagonista

		dx, dy := 0, 0
		if a.X < tx {
			dx = 1
		} else if a.X > tx {
			dx = -1
		}
		if a.Y < ty {
			dy = 1
		} else if a.Y > ty {
			dy = -1
		}

		newX, newY := a.X+dx, a.Y+dy

		// Si atrapa al protagonista, pierde el jugador
		if newX == tx && newY == ty {
			moveCh <- Move{OldX: a.X, OldY: a.Y, NewX: newX, NewY: newY, Icon: a.Icon()}
			doneCh <- Done{Winner: "Agente"}
			return
		}

		if b.IsEmpty(newX, newY) {
			moveCh <- Move{OldX: a.X, OldY: a.Y, NewX: newX, NewY: newY, Icon: a.Icon()}
			a.X, a.Y = newX, newY
		}
	}
}

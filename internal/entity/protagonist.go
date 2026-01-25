package entity

import (
	"matrix-game/internal/board"
	"time"
)

type Protagonist struct {
	X, Y int
}

func (p *Protagonist) Icon() rune {
	return 'N'
}

func (p *Protagonist) Run(b *board.Board, moveCh chan<- Move, doneCh chan<- Done) {
	for {
		time.Sleep(300 * time.Millisecond)

		tx, ty := b.NearestPhone(p.X, p.Y)
		if tx == -1 && ty == -1 {
			continue
		}

		dx, dy := 0, 0
		if p.X < tx {
			dx = 1
		} else if p.X > tx {
			dx = -1
		}
		if p.Y < ty {
			dy = 1
		} else if p.Y > ty {
			dy = -1
		}

		newX, newY := p.X+dx, p.Y+dy

		// Si va a un teléfono, gana el juego
		if b.Get(newX, newY) == 'T' {
			moveCh <- Move{OldX: p.X, OldY: p.Y, NewX: newX, NewY: newY, Icon: p.Icon()}
			doneCh <- Done{Winner: "Protagonista"}
			return
		}

		if b.IsEmpty(newX, newY) {
			moveCh <- Move{OldX: p.X, OldY: p.Y, NewX: newX, NewY: newY, Icon: p.Icon()}
			p.X, p.Y = newX, newY
		}
	}
}

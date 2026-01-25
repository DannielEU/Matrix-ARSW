package engine

import (
	"fmt"
	"matrix-game/internal/board"
	"matrix-game/internal/entity"
	"matrix-game/internal/render"
	"time"
)

func Run(b *board.Board, moveCh <-chan entity.Move, doneCh <-chan entity.Done) {
	for {
		select {
		case mv := <-moveCh:
			b.Set(mv.OldX, mv.OldY, ' ')
			b.Set(mv.NewX, mv.NewY, mv.Icon)
		case d := <-doneCh:
			render.Draw(b)
			fmt.Printf("\nJuego terminado! Ganador: %s\n", d.Winner)
			return
		default:
			render.Draw(b)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

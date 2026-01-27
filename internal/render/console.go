package render

import (
	"fmt"
	"matrix-game/internal/board"
)

func Draw(b *board.Board) {
	fmt.Print("\033[H\033[2J") // limpiar consola
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			fmt.Printf("%c ", b.Get(i, j))
		}
		fmt.Println()
		
	}
	fmt.Println("#-----------------------#")
}

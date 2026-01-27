package board

import (
	"math"
	"sync"
)

type Board struct {
	Grid [][]rune
	Size int
	mu   sync.Mutex
}

func New(size int) *Board {
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
		for j := range grid[i] {
			grid[i][j] = '□'
		}
	}
	return &Board{
		Grid: grid,
		Size: size,
	}
}

func (b *Board) Set(x, y int, r rune) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if x >= 0 && x < b.Size && y >= 0 && y < b.Size {
		b.Grid[x][y] = r
	}
}

func (b *Board) Get(x, y int) rune {
	b.mu.Lock()
	defer b.mu.Unlock()
	if x >= 0 && x < b.Size && y >= 0 && y < b.Size {
		return b.Grid[x][y]
	}
	return '#'
}

func (b *Board) IsEmpty(x, y int) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if x < 0 || x >= b.Size || y < 0 || y >= b.Size {
		return false
	}
	return b.Grid[x][y] == '□'
}

func (b *Board) NearestPhone(x, y int) (int, int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	minDist := math.MaxFloat64
	tx, ty := -1, -1
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.Grid[i][j] == 'T' {
				dist := math.Sqrt(float64((x-i)*(x-i) + (y-j)*(y-j)))
				if dist < minDist {
					minDist = dist
					tx, ty = i, j
				}
			}
		}
	}
	return tx, ty
}

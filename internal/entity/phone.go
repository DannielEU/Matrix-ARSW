package entity

type Phone struct {
	X, Y int
}

func (p *Phone) Icon() rune {
	return 'T'
}

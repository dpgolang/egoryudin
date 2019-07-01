package main

type Board struct {
	Length, Width uint
}

func (b *Board) GetBoardSize() uint {
	return b.Length * b.Width
} //ф-ия площади шахматной доски

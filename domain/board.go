package domain

import (
	"errors"
)

var (
	initSquares = [4][4]string{
		[4]string{"", "1", "2", "3"},
		[4]string{"1", "■", "■", "■"},
		[4]string{"2", "■", "■", "■"},
		[4]string{"3", "■", "■", "■"},
	}
)

type board struct {
	squares [4][4]string
}

func newBoard() board {
	return board{initSquares}
}

func (b *board) Put(row int, col int, turn string) error {
	if 1 > row || row > 3 {
		return errors.New("rowの入力は1~3の間でのみ可能です")
	}
	if 1 > col || col > 3 {
		return errors.New("colの入力は1~3の間でのみ可能です")
	}
	if b.squares[row][col] != "■" {
		return errors.New("既に入力されている場所です")
	}
	b.squares[row][col] = turn
	return nil
}

func (b *board) IsFull() bool {
	for r := 1; r <= 3; r++ {
		for c := 1; c <= 3; c++ {
			if b.squares[r][c] == "■" {
				return false
			}
		}
	}
	return true
}

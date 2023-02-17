package domain

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

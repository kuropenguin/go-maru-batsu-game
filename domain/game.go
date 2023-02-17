package domain

type (
	Game struct {
		board   board
		player1 player
		player2 player
		status  status
		turn    turn
	}

	turn   int
	status int

	player struct {
		name string
	}
)

const (
	turn1 turn = iota + 1
	turn2 turn = iota + 1
)

const (
	playing status = iota
	win     status = iota
	draw    status = iota
)

const (
	maru  string = "●"
	batsu string = "✖️"
)

func NewGame() *Game {
	return &Game{
		board:   newBoard(),
		player1: player{name: "●"},
		player2: player{name: "✖️"},
		turn:    turn1,
		status:  playing,
	}
}

func (g *Game) CurrentBoard() [4][4]string {
	return g.board.squares
}

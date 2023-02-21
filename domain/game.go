package domain

type (
	Game struct {
		Board   board
		player1 player
		player2 player
		status  status
		turn    Turn
	}

	Turn   string
	status int

	player struct {
		name string
	}
)

const (
	playing status = iota
	win     status = iota
	draw    status = iota
)

const (
	Maru  Turn = "●"
	Batsu Turn = "✖️"
)

func NewGame() *Game {
	return &Game{
		Board:   newBoard(),
		player1: player{name: "●"},
		player2: player{name: "✖️"},
		turn:    Maru,
		status:  playing,
	}
}

func (g *Game) CurrentBoard() [4][4]string {
	return g.Board.squares
}

func (g *Game) CurrentTurn() Turn {
	return g.turn
}

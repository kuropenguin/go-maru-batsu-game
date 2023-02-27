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
	Playing status = iota
	Win     status = iota
	Draw    status = iota
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
		status:  Playing,
	}
}

func (g *Game) CurrentBoard() [4][4]string {
	return g.Board.squares
}

func (g *Game) CurrentTurn() Turn {
	return g.turn
}

func (g *Game) ChangeTurn() {
	if g.turn == Maru {
		g.turn = Batsu
		return
	}
	g.turn = Maru
}

func (g *Game) Judge() {
	turnStr := string(g.turn)
	for i := 1; i <= 3; i++ {
		if g.Board.squares[i][1] == turnStr && g.Board.squares[i][2] == turnStr && g.Board.squares[i][3] == turnStr {
			g.status = Win
			return
		}
		if g.Board.squares[1][i] == turnStr && g.Board.squares[2][i] == turnStr && g.Board.squares[3][i] == turnStr {
			g.status = Win
			return
		}
	}
	if g.Board.squares[1][1] == turnStr && g.Board.squares[2][2] == turnStr && g.Board.squares[3][3] == turnStr {
		g.status = Win
	}
	if g.Board.squares[3][1] == turnStr && g.Board.squares[2][2] == turnStr && g.Board.squares[1][3] == turnStr {
		g.status = Win
	}

	if g.Board.IsFull() {
		g.status = Draw
	}
}

func (g *Game) Status() status {
	return g.status
}

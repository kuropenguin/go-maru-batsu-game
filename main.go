package main

import (
	"fmt"

	"github.com/kuropenguin/go-maru-batsu-game/controller"
	"github.com/kuropenguin/go-maru-batsu-game/domain"
)

func main() {
	game := domain.NewGame()
	controller := controller.Controller{}
	for {
		for _, row := range game.CurrentBoard() {
			fmt.Println(row)
		}
		// fmt.Printf("%s's turn\n", game.turn)
		fmt.Println("マスを選択してください(row, col). 例: 1,3")

		input, err := controller.SelectSquare()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("もう一度入力してください")
			continue
		}
		err = game.Board.Put(input.Row, input.Col, string(game.CurrentTurn()))
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("もう一度入力してください")
			continue
		}

		// Board[row][col] = turn
		// for i := 1; i <= 3; i++ {
		// 	if Board[i][1] == turn && Board[i][2] == turn && Board[i][3] == turn {
		// 		status = win
		// 		break
		// 	}
		// 	if Board[1][i] == turn && Board[2][i] == turn && Board[3][i] == turn {
		// 		status = win
		// 		break
		// 	}
		// }
		// if Board[1][1] == turn && Board[2][2] == turn && Board[3][3] == turn {
		// 	status = win
		// }
		// if Board[3][1] == turn && Board[2][2] == turn && Board[1][3] == turn {
		// 	status = win
		// }

		// isDraw := true
		// for r := 1; r <= 3; r++ {
		// 	for c := 1; c <= 3; c++ {
		// 		if Board[r][c] == "■" {
		// 			isDraw = false
		// 		}
		// 	}
		// }

		// if isDraw {
		// 	status = draw
		// }

		// if status == win {
		// 	fmt.Printf("%s WIN\n", turn)
		// 	break
		// }
		// if status == draw {
		// 	fmt.Printf("%s DRAW\n", turn)
		// 	break
		// }
		// if turn == "●" {
		// 	turn = "✖️"
		// } else {
		// 	turn = "●"
		// }
	}
	fmt.Println("Game Done")
}

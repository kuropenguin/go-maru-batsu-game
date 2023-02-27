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

		game.Judge()

		if game.Status() == domain.Win {
			fmt.Printf("%s is win!\n", game.CurrentTurn())
			break
		}

		if game.Status() == domain.Draw {
			fmt.Println("Draw!")
			break
		}

		game.ChangeTurn()
	}
	fmt.Println("Game Done")
}

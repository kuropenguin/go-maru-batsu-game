package main

import (
	"fmt"

	"github.com/kuropenguin/go-maru-batsu-game/domain"
)

func main() {
	game := domain.NewGame()
	for {
		for _, row := range game.CurrentBoard() {
			fmt.Println(row)
		}
		// fmt.Printf("%s's turn\n", game.turn)
		// fmt.Println("Please enter the row, col(1~3). For example: 1,3")
		// reader := bufio.NewReader(os.Stdin)
		// input, err := reader.ReadString('\n')
		// if err != nil && err != io.EOF {
		// 	fmt.Println(err)
		// 	fmt.Println("もう一度入力してください")
		// 	continue
		// }

		// input = strings.Trim(input, "\n")
		// inputArr := strings.Split(input, ",")
		// if len(inputArr) > 2 {
		// 	fmt.Println("入力が多すぎます。もう一度入力してください")
		// 	continue
		// }
		// if len(inputArr) < 2 {
		// 	fmt.Println("入力が少なすぎます。もう一度入力してください")
		// 	continue
		// }

		// row, err := strconv.Atoi(inputArr[0])
		// if err != nil {
		// 	fmt.Println("rowの入力で数字以外が使用されました。もう一度入力してください")
		// 	continue
		// }
		// if !(1 <= row && row <= 3) {
		// 	fmt.Println("rowの入力は1~3の間でのみ可能です。もう一度入力してください")
		// 	continue
		// }
		// col, err := strconv.Atoi(inputArr[1])
		// if err != nil {
		// 	fmt.Println("colの入力で数字以外が使用されました。もう一度入力してください")
		// 	continue
		// }
		// if !(1 <= col && col <= 3) {
		// 	fmt.Println("colの入力は1~3の間でのみ可能です。もう一度入力してください")
		// 	continue
		// }

		// if Board[row][col] != "■" {
		// 	fmt.Println("既に入力されている場所です。もう一度入力してください")
		// 	continue
		// }
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

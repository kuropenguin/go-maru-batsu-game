package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	Board = [][]string{
		[]string{"", "1", "2", "3"},
		[]string{"1", "■", "■", "■"},
		[]string{"2", "■", "■", "■"},
		[]string{"3", "■", "■", "■"},
	}

	turn string
)

func main() {
	turn = "●"
	for {
		for _, row := range Board {
			fmt.Println(row)
		}

		fmt.Printf("%s's turn\n", turn)
		fmt.Println("Please enter the row, col(1~3). For example: 1,3")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
			fmt.Println("もう一度入力してください")
			continue
		}

		input = strings.Trim(input, "\n")
		inputArr := strings.Split(input, ",")
		if len(inputArr) > 2 {
			fmt.Println("入力が多すぎます。もう一度入力してください")
			continue
		}
		if len(inputArr) < 2 {
			fmt.Println("入力が少なすぎます。もう一度入力してください")
			continue
		}
		fmt.Println(inputArr)

		row, err := strconv.Atoi(inputArr[0])
		if err != nil {
			fmt.Println("rowの入力で数字以外が使用されました。もう一度入力してください")
			continue
		}
		if !(1 <= row && row <= 3) {
			fmt.Println("rowの入力は1~3の間でのみ可能です。もう一度入力してください")
			continue
		}
		col, err := strconv.Atoi(inputArr[1])
		if err != nil {
			fmt.Println("colの入力で数字以外が使用されました。もう一度入力してください")
			continue
		}
		if !(1 <= col && col <= 3) {
			fmt.Println("colの入力は1~3の間でのみ可能です。もう一度入力してください")
			continue
		}

		//入力
		Board[row][col] = turn
		if turn == "●" {
			turn = "✖️"
		} else {
			turn = "●"
		}
	}
	fmt.Println("Game Done")
}

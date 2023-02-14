package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var Board = [][]string{
	[]string{"", "1", "2", "3"},
	[]string{"1", "■", "■", "■"},
	[]string{"2", "■", "■", "■"},
	[]string{"3", "■", "■", "■"},
}

func main() {
	for {
		for _, row := range Board {
			fmt.Println(row)
		}

		fmt.Println("Please enter the row, col(1~3). For example: 1,3")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
			fmt.Println("もう一度入力してください")
			continue
		}
		fmt.Println(text)
	}
	fmt.Println("Game Done")
}

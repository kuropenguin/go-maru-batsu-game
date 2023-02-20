package controller

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Row int
	Col int
}

type Controller struct {
}

func (c *Controller) SelectSquare() (input *Input, err error) {
	reader := bufio.NewReader(os.Stdin)
	inputLine, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	inputLine = strings.Trim(inputLine, "\n")
	inputArr := strings.Split(inputLine, ",")
	if len(inputArr) > 2 {
		return nil, errors.New("入力が多すぎます")
	}
	if len(inputArr) < 2 {
		return nil, errors.New("入力が少なすぎます")
	}

	row, err := strconv.Atoi(inputArr[0])
	if err != nil {
		return nil, errors.New("rowの入力で数字以外が使用されました。もう一度入力してください")
	}
	col, err := strconv.Atoi(inputArr[1])
	if err != nil {
		return nil, errors.New("colの入力で数字以外が使用されました。もう一度入力してください")
	}
	return &Input{
		Row: row,
		Col: col,
	}, nil
}

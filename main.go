package main

import "fmt"

var Board = [][]string{
	[]string{"■", "■", "■"},
	[]string{"■", "■", "■"},
	[]string{"■", "■", "■"},
}

func main() {
	for _, row := range Board {
		fmt.Println(row)
	}
}

package main

import (
	"fmt"
)

const (
	boardSize = 3
)

var board [boardSize][boardSize]string

func printBoard() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			fmt.Print(board[i][j], ".")
		}
		fmt.Print("\n")
	}
}

func main() {
	printBoard()

}

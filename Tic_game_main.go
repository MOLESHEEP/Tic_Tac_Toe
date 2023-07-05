package main

import (
	"fmt"
)

const (
	boardSize = 3
	emptyCell = "."
	playerX   = "X"
	playerO   = "O"
)

var board [boardSize][boardSize]string

func initialBoard() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[i][j] = emptyCell
		}
	}
}

func printBoard() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			fmt.Print(board[i][j])
		}
		fmt.Print("\n")
	}
}

func isBoardFull() bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == emptyCell {
				return false
			}
		}
	}
	return true
}

func main() {
	initialBoard()
	var row, col int
	for {
		fmt.Print("Input (x,y) ")
		fmt.Scan(&row, &col)

		if row < 0 || row >= boardSize || col < 0 || col >= boardSize {
			fmt.Println("無効な行または列です。もう一度入力してください。")
		}
		board[row][col] = playerO
		printBoard()

		if isBoardFull() {
			break
		}
	}
}

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
var currentPlayer string

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

func checkWin() bool {
	// 横方向と縦方向をチェック
	for i := 0; i < boardSize; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return true
		}
	}

	// 対角線をチェック
	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return true
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return true
	}

	return false
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
	currentPlayer = playerO
	fmt.Println("〇×ゲームを開始します。")

	var row, col int
	for {

		fmt.Print("Input (x,y) ")
		fmt.Scan(&row, &col)

		if row < 0 || row >= boardSize || col < 0 || col >= boardSize {
			fmt.Println("無効な行または列です。もう一度入力してください。")
		}
		// セルにプレーヤーの記号を入れる
		board[row][col] = currentPlayer
		printBoard()

		if checkWin() {
			fmt.Print("\n")
			//printBoard()
			fmt.Printf("プレーヤー %s の勝利です！\n", currentPlayer)
			break
		}
		if isBoardFull() {
			break
		}

		// プレーヤー交代
		if currentPlayer == playerX {
			currentPlayer = playerO
		} else {
			currentPlayer = playerX
		}

	}
}

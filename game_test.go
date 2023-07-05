package main

import (
	"testing"
)

func gametest(t *testing.T) {
	board := main_test(0, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != "O" && board[i][j] != "X" {
				t.Error("!!!!!!!!Error!!!!!!!!!!")
			}
		}
	}

}

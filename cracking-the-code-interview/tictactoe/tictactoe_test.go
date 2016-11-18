package tictactoe

import (
	"testing"
)

func TestIsWinning(t *testing.T) {

	var board Board

	board = Board{
		{"x", "o", "x"},
		{"o", "o", "o"},
		{"x", "o", "x"}}

	if !board.IsWinning() {
		t.Error("should be winning")
	}

	board = Board{
		{"x", "o", "x"},
		{"o", "x", "o"},
		{"x", "o", "x"}}

	if !board.IsWinning() {
		t.Error("should be winning")
	}

	board = Board{
		{"x", "o", "x"},
		{"x", "o", "o"},
		{"o", "x", "x"}}

	if board.IsWinning() {
		t.Error("should not be winning")
	}


}

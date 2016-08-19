package boggle

import (
	"testing"
	"os"
)

func TestReadBoard(t *testing.T) {
	file, error := os.Open("test/test-1-input.txt")
	if (error != nil) {
		t.Error("could not read test/test-1-input.txt")
	}
	board, error := ReadBoard(file)
	if (board[0][0] != 'b') {
		t.Error("didn't have a 'b' as the first char")
	}
}

func TestReadWord(t *testing.T) {
	file, error := os.Open("test/test-1-input.txt")
	if (error != nil) {
		t.Error("could not read test/test-1-input.txt")
	}
	board, error := ReadBoard(file)
	word, error := board.ReadWord(Position{0, 0}, Position{2, 0});
	if (word.String() != "bag") {
		t.Error("the first three letters should be `bag`")
	}
}

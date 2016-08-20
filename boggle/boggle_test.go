package boggle

import (
	"testing"
	"os"
	"fmt"
)

func TestReadBoard(t *testing.T) {
	file, error := os.Open("test/test-1-input.txt")
	if (error != nil) {
		t.Error("could not read test/test-1-input.txt")
	}
	board, error := ReadBoard(file)
	if (board[0][0] != 'a') {
		t.Error("didn't have a 'a' as the first char")
	}
}

func TestReadWordList(t *testing.T) {
	file, error := os.Open("words.txt")
	if (error != nil) {
		t.Error("could not read `words.txt`")
	}
	words, error := ReadWordList(file)
	if (error != nil) {
		t.Error(error)
	}
	if (!words.HasWord(NewWord("ball"))) {
		t.Error("should have ball")
	}
}

func TestFindWords(t *testing.T) {
	file, error := os.Open("words.txt")
	if (error != nil) {
		t.Error("could not read `words.txt`")
	}
	wordList, error := ReadWordList(file)
	inputFile, error := os.Open("test/test-1-input.txt")
	if (error != nil) {
		t.Error("could not read test/test-1-input.txt")
	}
	board, error := ReadBoard(inputFile)
	words := board.FindWords(wordList)
	if (words.Size() != 1) {
		fmt.Printf("%v", words)
		t.Error("should have found 1 word but didn't")
	}
}

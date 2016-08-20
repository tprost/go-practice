package boggle

import "os"
import "bufio"

func ReadBoard(file *os.File) (*Board, error) {
	// initialize board
	var board Board;

	// for each line
	reader := bufio.NewReader(file)
	for y := 0; y < len(board[0]); y++ {
		// read a row of letters
		line, _, error := reader.ReadLine()
		// throw errors if there are errors
		if (error != nil) {
			return &board, nil
		}
		// set a row of the board
		for x := 0; x < len(board); x++ {
			board[x][y] = Letter(line[x])
		}
	}

	// return the board
	return &board, nil
}

func (board *Board) FindWords(wordList *WordList) *WordList {

	words := NewWordList();

	width := len(board)
	height := len(board[0])

	// for each starting position
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {

			start := Position{x, y}

			// horizontal words ltr
			for end := x + 1; end < width; end++ {
				word, _ := board.ReadWord(start, Position{end, y})
				if (wordList.HasWord(word)) {
					words.AddWord(word);
				}
			}

			// horizontal words rtl
			for end := y - 1; end >= 0; end-- {

			}
			// vertical words

		}
	}

	return words
}

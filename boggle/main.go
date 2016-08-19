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

// func (board *Board) FindWords(wordList *WordList) []string {

// 	width := len(board[0])
// 	height := len(board)

// 	// for each starting position
// 	for i := 0; i < height; i++ {
// 		for j := 0; j < width; j++ {


// 			// horizontal words ltr
// 			for end := j + 1; end < width; end++ {

// 			}

// 			// horizontal words rtl
// 			for end := j - 1; end >= 0; end-- {

// 			}
// 			// vertical words


// 			board[i][j] = line[j]
// 		}
// 	}
// }

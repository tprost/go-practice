package boggle

import "errors"

type Boggle struct {
	Board Board
	WordList WordList
}

type Position struct {
	X int
	Y int
}

type Letter byte
type Board [5][5]Letter
type Word []Letter

func NewWord(wordString string) Word {

	var word Word
	for i := 0; i < len(wordString); i++ {
		word = append(word, Letter(wordString[i]))
	}
	return word
}

func (board Board) String() string {
	str := ""
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			str = str + string(board[x][y])
		}
		str = str + "\n"
	}
	return str
}

// func NewWord(letters []*Letter) *Word {

// }

func (word Word) String() string {
	str := ""
	for _, letter := range word {
		str = str + string(letter)
	}
	return str
}

func (board *Board) ReadLetter(position Position) (Letter, error) {
	return board[position.X][position.Y], nil
}

func (board *Board) ReadWord(start Position, end Position) (Word, error) {

	var word Word

	// TODO look for errors
	if (start.X == end.X && start.Y == end.Y) {
		return nil, errors.New("positions are equal")
	}

	if (start.X > end.X && start.Y == end.Y) {
		for i := start.X; i >= end.X; i-- {
			letter, error := board.ReadLetter(Position{i, start.Y})
			if (error != nil) {
				return nil, error
			}
			word = append(word, letter)
		}
	} else if (start.X < end.X && start.Y == end.Y) {
		for i := start.X; i <= end.X; i++ {
			letter, error := board.ReadLetter(Position{i, start.Y})
			if (error != nil) {
				return nil, error
			}
			word = append(word, letter)
		}
	} else if (start.Y > end.Y && start.X == end.X) {
		for i := start.Y; i >= end.Y; i-- {
			letter, error := board.ReadLetter(Position{start.X, i})
			if (error != nil) {
				return nil, error
			}
			word = append(word, letter)
		}
	} else if (start.Y < end.Y && start.X == end.X) {
		for i := start.Y; i <= end.Y; i++ {
			letter, error := board.ReadLetter(Position{start.X, i})
			if (error != nil) {
				return nil, error
			}
			word = append(word, letter)
		}
	} else {
		return nil, errors.New("cannot read word based on those positions")
	}
	return word, nil
}

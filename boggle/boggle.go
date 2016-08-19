package boggle

import "errors"

type Position struct {
	X int
	Y int
}

type Letter byte
type Board [5][5]Letter
type Word []Letter

// func NewWord(letters []*Letter) *Word {

// }

func (word Word) String() string {
	return string([]byte(word))
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

	if (start.X > end.X) {
		for i := start.X; i >= end.X; i-- {
			letter, error := board.ReadLetter(Position{i, start.Y})
			if (error != nil) {
				return nil, error
			}
			word = append(word, letter)
		}
	} else {
		for i := start.X; i <= end.X; i++ {
			letter, error := board.ReadLetter(Position{i, start.Y})
			if (error != nil) {
				return nil, error
			}
			word = append(word, letter)
		}
	}
	if (start.Y > end.Y) {
		for i := start.Y; i >= end.Y; i-- {
			letter, error := board.ReadLetter(Position{start.X, i})
			if (error != nil) {
				return nil, error
			}
			word = append(word, letter)
		}
	} else {
		for i := start.Y; i <= end.Y; i++ {
			letter, error := board.ReadLetter(Position{start.X, i})
			if (error != nil) {
				return nil, error
			}
			word = append(word, letter)
		}
	}
	return word, nil
}

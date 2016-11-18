package tictactoe

type Board [3][3]string
type Coordinate [2]int
type Line [3]Coordinate

func (board Board) IsWinning() bool {



	var generatePossibleLocations = func() []Line {
		possibleLocations := []Line{}
		// generate straight lines
		for i := 0; i < 3; i++ {
			line1 := Line{}
			line2 := Line{}
			for j := 0; j < 3; j++ {
				line1[j] = Coordinate{j, i}
				line2[j] = Coordinate{i, j}
			}
			possibleLocations = append(possibleLocations, line1, line2)
		}
		line1 := Line{}
		line2 := Line{}
		// generate diagonal lines
		for i := 0; i < 3; i++ {
			line1[i] = Coordinate{i, i}
			line2[i] = Coordinate{i, 2-i}
		}
		possibleLocations = append(possibleLocations, line1, line2)
		return possibleLocations
	}

	possibleLocations := generatePossibleLocations()

	var isWinningLine = func(line Line) bool {
		if board[line[0][0]][line[0][1]] != board[line[1][0]][line[1][1]] {
			return false
		}
		if board[line[1][0]][line[0][1]] != board[line[2][0]][line[1][1]] {
			return false
		}
		return true
	}

	for _, line := range possibleLocations {
		if isWinningLine(line) {
			return true
		}
	}
	return false
}

package movegeneration

import "math"

var ValidKnightMoves [64]uint64

func init() {
	initialiseValidKnightMoves()
}

func initialiseValidKnightMoves() {

	for startColumn := 0; startColumn < 8; startColumn++ {
		for startRow := 0; startRow < 8; startRow++ {

			validMoves := uint64(0)

			var column = startColumn + 1
			var row = startRow + 2

			if row >= 0 && row < 8 && column >= 0 && column < 8 {
				validMoves |= calculateSquareValue(column, row)
			}

			column = startColumn + 2
			row = startRow + 1

			if row >= 0 && row < 8 && column >= 0 && column < 8 {
				validMoves |= calculateSquareValue(column, row)
			}

			column = startColumn + 2
			row = startRow - 1

			if row >= 0 && row < 8 && column >= 0 && column < 8 {
				validMoves |= calculateSquareValue(column, row)
			}

			column = startColumn + 1
			row = startRow - 2

			if row >= 0 && row < 8 && column >= 0 && column < 8 {
				validMoves |= calculateSquareValue(column, row)
			}

			column = startColumn - 1
			row = startRow - 2

			if row >= 0 && row < 8 && column >= 0 && column < 8 {
				validMoves |= calculateSquareValue(column, row)
			}

			column = startColumn - 2
			row = startRow - 1

			if row >= 0 && row < 8 && column >= 0 && column < 8 {
				validMoves |= calculateSquareValue(column, row)
			}

			column = startColumn - 2
			row = startRow + 1

			if row >= 0 && row < 8 && column >= 0 && column < 8 {
				validMoves |= calculateSquareValue(column, row)
			}

			column = startColumn - 1
			row = startRow + 2

			if row >= 0 && row < 8 && column >= 0 && column < 8 {
				validMoves |= calculateSquareValue(column, row)
			}

			//move.Moves = validMoves;

			var startSquare = startColumn + (startRow * 8)

			ValidKnightMoves[startSquare] = validMoves
		}
	}
}

func calculateSquareValue(column, row int) uint64 {
	moveIndex := column + (row * 8)
	var moveSquare = uint64(math.Pow(2, float64(moveIndex)))

	return moveSquare
}

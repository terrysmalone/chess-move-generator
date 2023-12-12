package lookuptables

var UpBoard [64]uint64
var RightBoard [64]uint64
var DownBoard [64]uint64
var LeftBoard [64]uint64

var UpRightBoard [64]uint64
var DownRightBoard [64]uint64
var DownLeftBoard [64]uint64
var UpLeftBoard [64]uint64

func calculateDirectionBoards() {
	calculateUpDirectionBoards()
	calculateRightDirectionBoards()
	calculateDownDirectionBoards()
	calculateLeftDirectionBoards()

	//calculateUpRightDirectionBoards()
	//calculateDownRightDirectionBoards()
	//calculateDownLeftDirectionBoards()
	//calculateUpLeftDirectionBoards()
}

func calculateUpDirectionBoards() {
	// Only goes to 7 as row 8 will all be empty since there are no up moves
	for startRow := 0; startRow < 7; startRow++ {
		for startColumn := 0; startColumn < 8; startColumn++ {
			row := startRow + 1

			var upSquares uint64

			for row < 8 {
				upSquares |= BitboardValueFromPosition[startColumn][row]
				row++
			}

			UpBoard[(startRow*8)+startColumn] = upSquares
		}
	}
}

func calculateRightDirectionBoards() {
	for startRow := 0; startRow < 8; startRow++ {
		//Only goes to 7 as column 8 will all be empty since there are no right moves
		for startColumn := 0; startColumn < 7; startColumn++ {
			column := startColumn + 1

			var rightSquares uint64

			for column < 8 {
				rightSquares |= BitboardValueFromPosition[column][startRow]
				column++
			}

			RightBoard[(startRow*8)+startColumn] = rightSquares
		}
	}
}

func calculateDownDirectionBoards() {
	//Starts at 1 as row 0 will all be empty since there are no up moves
	for startRow := 1; startRow < 8; startRow++ {
		for startColumn := 0; startColumn < 8; startColumn++ {
			row := startRow - 1

			var downSquares uint64

			for row >= 0 {
				downSquares |= BitboardValueFromPosition[startColumn][row]
				row--
			}

			DownBoard[(startRow*8)+startColumn] = downSquares
		}
	}
}

func calculateLeftDirectionBoards() {
	for startRow := 0; startRow < 8; startRow++ {
		//Starts at 1 as column 0 will all be empty since there are no left moves
		for startColumn := 1; startColumn < 8; startColumn++ {
			column := startColumn - 1

			var leftSquares uint64

			for column >= 0 {
				leftSquares |= BitboardValueFromPosition[column][startRow]
				column--
			}

			LeftBoard[(startRow*8)+startColumn] = leftSquares
		}
	}
}
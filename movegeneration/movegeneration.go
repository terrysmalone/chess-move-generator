package movegeneration

import (
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
)

type PieceMove struct {
}

func CalculateAllMoves(gameBoard *boardrepresentation.GameBoard) {
	allMoves := calculateAllPseudoLegalMoves(gameBoard)

	removeSelfCheckingMoves(gameBoard, &allMoves)

	removeSelfCheckingCastlingMoves(gameBoard, &allMoves)
}

func calculateAllPseudoLegalMoves(gameBoard *boardrepresentation.GameBoard) []PieceMove {
	// TODO
	return []PieceMove{}
}

// We want to change the slice in place since passing the slice by value and then returning a new slice will
// be far too inefficient for the amount of times this is called. So, we're passing in the address of the slice
func removeSelfCheckingMoves(gameBoard *boardrepresentation.GameBoard, pieceMoves *[]PieceMove) {
	// TODO
}

// We want to change the slice in place since passing the slice by value and then returning a new slice will
// be far too inefficient for the amount of times this is called. So, we're passing in the address of the slice
func removeSelfCheckingCastlingMoves(gameBoard *boardrepresentation.GameBoard, pieceMoves *[]PieceMove) {
	// TODO
}

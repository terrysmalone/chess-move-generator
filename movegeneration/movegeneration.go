package movegeneration

import (
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
	"github.com/terrysmalone/chess-move-generator/boardrepresentation/bitboardoperations"
	"github.com/terrysmalone/chess-move-generator/boardsearching/lookuptables"
)

type PieceMove struct {
}

func CalculateAllMoves(gameBoard *boardrepresentation.GameBoard) {
	allMoves := calculateAllPseudoLegalMoves(gameBoard)

	removeSelfCheckingMoves(gameBoard, &allMoves)

	removeSelfCheckingCastlingMoves(gameBoard, &allMoves)
}

// calculateAllPseudoLegalMoves returns all possible pseudo legal moves from the given postion.
// This will include self-checking moves
func calculateAllPseudoLegalMoves(gameBoard *boardrepresentation.GameBoard) []PieceMove {

	allMoves := []PieceMove{}

	if gameBoard.WhiteToMove {
		calculateWhiteKnigthMoves(gameBoard.Board.WhiteKnights)
		// calculateBishopMoves
		// calculateRookMoves
		// calculateQueenMoves
		// calculateWhitePawnMoves
		// calculateWhiteKingMoves

		// calculateWhiteCastlingMoves
	} else {
		// calculateBlackKnigthMoves(gameBoard.Board.WhiteKnightMoves)
		// calculateBishopMoves
		// calculateRookMoves
		// calculateQueenMoves
		// calculateBlackPawnMoves
		// calculateBlackKingMoves
		// calculateBlackCastlingMoves
	}

	return allMoves
}

func calculateWhiteKnigthMoves(whiteKnightPositions uint64) ([]PieceMove, error) {
	// Get whiteKnightPositions from bitboard
	whiteKnightIndexes := bitboardoperations.GetSquareIndexesFromBitboard(whiteKnightPositions)

	// We need to iterate through the positions backwards
	index := len(whiteKnightIndexes) - 1

	for index >= 0 {
		currentPosition := whiteKnightIndexes[index]

		_ = lookuptables.BitboardValueFromIndex[currentPosition]
		_ = boardrepresentation.KnightPieceType

		// Calculate possible moves bitboard
		// Split moves

		// for each move
		//    if move & all black occupied squares > 0
		//       Add as a capture move
		//
		//   if move & Empty squares > 0
		//      Add as normal move

		index--
	}

	return []PieceMove{}, nil
}

// Note: We want to change the slice in place since passing the slice by value and then returning a new slice will
// be far too inefficient for the amount of times this is called. So, we're passing in the address of the slice
func removeSelfCheckingMoves(gameBoard *boardrepresentation.GameBoard, pieceMoves *[]PieceMove) {
	// TODO
}

// Note: We want to change the slice in place since passing the slice by value and then returning a new slice will
// be far too inefficient for the amount of times this is called. So, we're passing in the address of the slice
func removeSelfCheckingCastlingMoves(gameBoard *boardrepresentation.GameBoard, pieceMoves *[]PieceMove) {
	// TODO
}

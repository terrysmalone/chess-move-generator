package movegeneration

import (
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
	"github.com/terrysmalone/chess-move-generator/boardrepresentation/bitboardoperations"
	"github.com/terrysmalone/chess-move-generator/boardsearching/lookuptables"
)

func CalculateAllMoves(gameBoard *boardrepresentation.GameBoard) {
	gameBoard.CalculateUsefulBitboards()

	allMoves := []PieceMove{}
	calculateAllPseudoLegalMoves(gameBoard, &allMoves)

	removeSelfCheckingMoves(gameBoard, &allMoves)

	removeSelfCheckingCastlingMoves(gameBoard, &allMoves)
}

// calculateAllPseudoLegalMoves returns all possible pseudo legal moves from the given postion.
// This will include self-checking moves
func calculateAllPseudoLegalMoves(gameBoard *boardrepresentation.GameBoard, pieceMoves *[]PieceMove) []PieceMove {

	allMoves := []PieceMove{}

	if gameBoard.WhiteToMove {
		calculateKnightMoves(
			&allMoves,
			gameBoard.Board.WhiteKnights,
			gameBoard.UsefulBitboards.AllBlackOccupiedSquares,
			gameBoard.UsefulBitboards.EmptySquares)
		// calculateBishopMoves
		// calculateRookMoves
		// calculateQueenMoves
		// calculateWhitePawnMoves
		// calculateWhiteKingMoves

		// calculateWhiteCastlingMoves
	} else {
		calculateKnightMoves(
			&allMoves,
			gameBoard.Board.BlackKnights,
			gameBoard.UsefulBitboards.AllWhiteOccupiedSquares,
			gameBoard.UsefulBitboards.EmptySquares)
		// calculateBishopMoves
		// calculateRookMoves
		// calculateQueenMoves
		// calculateBlackPawnMoves
		// calculateBlackKingMoves
		// calculateBlackCastlingMoves
	}

	return allMoves
}

func calculateKnightMoves(pieceMoves *[]PieceMove, knights, enemyOccupied, emptySquares uint64) {
	knightIndexes := bitboardoperations.GetSquareIndexesFromBitboard(knights)

	// We need to iterate through the positions backwards
	index := len(knightIndexes) - 1

	for index >= 0 {
		currentPosition := knightIndexes[index]

		pieceBitboard := lookuptables.BitboardValueFromIndex[currentPosition]
		pieceType := boardrepresentation.KnightPieceType

		possibleMoves := ValidKnightMoves[currentPosition]
		splitMoves := bitboardoperations.SplitBitboard(possibleMoves)

		for _, move := range splitMoves {
			if (move & enemyOccupied) > 0 {
				*pieceMoves = append(*pieceMoves, PieceMove{
					PositionBitboard: pieceBitboard,
					MoveBitboard:     move,
					PieceType:        pieceType,
					MoveType:         boardrepresentation.CaptureMoveType,
				})
			}

			if (move & emptySquares) > 0 {
				*pieceMoves = append(*pieceMoves, PieceMove{
					PositionBitboard: pieceBitboard,
					MoveBitboard:     move,
					PieceType:        pieceType,
					MoveType:         boardrepresentation.NormalMoveType,
				})
			}
		}

		index--
	}
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

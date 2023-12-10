package movegeneration

import (
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
	"github.com/terrysmalone/chess-move-generator/boardrepresentation/bitboardoperations"
	"github.com/terrysmalone/chess-move-generator/boardsearching/lookuptables"
)

func CalculateAllMoves(gameBoard *boardrepresentation.GameBoard) {
	gameBoard.CalculateUsefulBitboards()

	allMoves := calculateAllPseudoLegalMoves(gameBoard)

	removeSelfCheckingMoves(gameBoard, &allMoves)

	removeSelfCheckingCastlingMoves(gameBoard, &allMoves)
}

// calculateAllPseudoLegalMoves returns all possible pseudo legal moves from the given postion.
// This will include self-checking moves
func calculateAllPseudoLegalMoves(gameBoard *boardrepresentation.GameBoard) []PieceMove {

	allMoves := []PieceMove{}

	if gameBoard.WhiteToMove {
		calculateWhiteKnightMoves(gameBoard)
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

func calculateWhiteKnightMoves(gameBoard *boardrepresentation.GameBoard) ([]PieceMove, error) {
	moves := []PieceMove{}
	// Get whiteKnightPositions from bitboard
	whiteKnightIndexes := bitboardoperations.GetSquareIndexesFromBitboard(gameBoard.Board.WhiteKnights)

	// We need to iterate through the positions backwards
	index := len(whiteKnightIndexes) - 1

	for index >= 0 {
		currentPosition := whiteKnightIndexes[index]

		pieceBitboard := lookuptables.BitboardValueFromIndex[currentPosition]
		pieceType := boardrepresentation.KnightPieceType

		possibleMoves := ValidKnightMoves[currentPosition]
		splitMoves := bitboardoperations.SplitBitboard(possibleMoves)

		for _, move := range splitMoves {
			if (move & gameBoard.UsefulBitboards.AllBlackOccupiedSquares) > 0 {
				moves = append(moves, PieceMove{
					PositionBitboard: pieceBitboard,
					MoveBitboard:     move,
					PieceType:        pieceType,
					MoveType:         boardrepresentation.CaptureMoveType,
				})
			}

			if (move & gameBoard.UsefulBitboards.EmptySquares) > 0 {
				moves = append(moves, PieceMove{
					PositionBitboard: pieceBitboard,
					MoveBitboard:     move,
					PieceType:        pieceType,
					MoveType:         boardrepresentation.NormalMoveType,
				})
			}
		}

		index--
	}

	return moves, nil
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

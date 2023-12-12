package movegeneration

import (
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
	"github.com/terrysmalone/chess-move-generator/boardrepresentation/bitboardoperations"
	"github.com/terrysmalone/chess-move-generator/boardsearching/lookuptables"
	"github.com/terrysmalone/chess-move-generator/boardsearching/piecechecking"
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
		calculateBishopMoves(
			&allMoves,
			gameBoard.Board.WhiteBishops,
			&gameBoard.UsefulBitboards,
			true)

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
		calculateBishopMoves(
			&allMoves,
			gameBoard.Board.WhiteBishops,
			&gameBoard.UsefulBitboards,
			false)
		// calculateRookMoves
		// calculateQueenMoves
		// calculateBlackPawnMoves
		// calculateBlackKingMoves
		// calculateBlackCastlingMoves
	}

	return allMoves
}

func calculateKnightMoves(pieceMoves *[]PieceMove, knights, enemyOccupied, emptySquares uint64) {
	knightsIndexes := bitboardoperations.GetSquareIndexesFromBitboard(knights)

	index := len(knightsIndexes) - 1

	for index >= 0 {
		currentPosition := knightsIndexes[index]

		pieceBitboard := lookuptables.BitboardValueFromIndex[currentPosition]

		possibleMoves := ValidKnightMoves[currentPosition]
		splitMoves := bitboardoperations.SplitBitboard(possibleMoves)

		for _, move := range splitMoves {
			if (move & enemyOccupied) > 0 {
				*pieceMoves = append(*pieceMoves, PieceMove{
					PositionBitboard: pieceBitboard,
					MoveBitboard:     move,
					PieceType:        boardrepresentation.KnightPieceType,
					MoveType:         boardrepresentation.CaptureMoveType,
				})
			}

			if (move & emptySquares) > 0 {
				*pieceMoves = append(*pieceMoves, PieceMove{
					PositionBitboard: pieceBitboard,
					MoveBitboard:     move,
					PieceType:        boardrepresentation.KnightPieceType,
					MoveType:         boardrepresentation.NormalMoveType,
				})
			}
		}

		index--
	}
}

func calculateBishopMoves(pieceMoves *[]PieceMove, bishops uint64, usefulBitboards *boardrepresentation.UsefulBitboards, whiteToMove bool) {
	bishopsIndexes := bitboardoperations.GetSquareIndexesFromBitboard(bishops)

	index := len(bishopsIndexes) - 1

	for index >= 0 {
		currentPosition := bishopsIndexes[index]
		bishopPosition := lookuptables.BitboardValueFromIndex[currentPosition]

		allowedMoves := piecechecking.CalculateAllowedBishopMoves(usefulBitboards, currentPosition, whiteToMove)

		// Positions in allowed moves and non-emprty squares are captures
		// (we've already excluded own pieces in CalculateAllowedBishopMoves)
		captureMoves := allowedMoves & ^usefulBitboards.EmptySquares
		splitAndAddMoves(pieceMoves, captureMoves, bishopPosition, boardrepresentation.BishopPieceType, boardrepresentation.CaptureMoveType)

		// Positions in allowed moves and emprty squares are non-capture moves
		normalMoves := allowedMoves & usefulBitboards.EmptySquares
		splitAndAddMoves(pieceMoves, normalMoves, bishopPosition, boardrepresentation.BishopPieceType, boardrepresentation.NormalMoveType)

		index--
	}
}

func splitAndAddMoves(pieceMoves *[]PieceMove, moves uint64, position uint64, pieceType boardrepresentation.PieceType, moveType boardrepresentation.MoveType) {
	splitMoves := bitboardoperations.SplitBitboard(moves)

	for _, move := range splitMoves {
		if move > 0 {
			*pieceMoves = append(
				*pieceMoves,
				PieceMove{
					PositionBitboard: position,
					MoveBitboard:     move,
					PieceType:        pieceType,
					MoveType:         moveType,
				})
		}
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

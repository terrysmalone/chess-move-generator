package piecechecking

import (
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
)

func CalculateAllowedBishopMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	return calculateAllowedUpRightMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedDownRightMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedDownLeftMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedUpLeftMoves(usefulBitboards, pieceIndex, whiteToMove)
}

func calculateAllowedUpRightMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	//upRightBoard := lookuptables.UpRightBoard[pieceIndex]
	// CalculateAllowedUpRightMovesFromBoard(usefulBitboards, upRightBoard, whiteToMove)
	return 0
}

func calculateAllowedDownRightMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	return 0
}

func calculateAllowedDownLeftMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	return 0
}

func calculateAllowedUpLeftMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	return 0
}

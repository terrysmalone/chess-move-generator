package piecechecking

import (
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
	"github.com/terrysmalone/chess-move-generator/boardsearching/lookuptables"
)

func CalculateAllowedBishopMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	return calculateAllowedUpRightMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedDownRightMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedDownLeftMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedUpLeftMoves(usefulBitboards, pieceIndex, whiteToMove)
}

func CalculateAllowedRookMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	return calculateAllowedUpMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedRightMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedDownMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedLeftMoves(usefulBitboards, pieceIndex, whiteToMove)

}

func CalculateAllowedQueenMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	return calculateAllowedUpRightMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedDownRightMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedDownLeftMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedUpLeftMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedUpMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedRightMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedDownMoves(usefulBitboards, pieceIndex, whiteToMove) |
		calculateAllowedLeftMoves(usefulBitboards, pieceIndex, whiteToMove)

}

func calculateAllowedUpRightMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	upRightBoard := lookuptables.UpRightBoard[pieceIndex]
	return calculateAllowedUpRightMovesFromBoard(usefulBitboards, upRightBoard, whiteToMove)
}

func calculateAllowedUpRightMovesFromBoard(usefulBitboards *boardrepresentation.UsefulBitboards, upRightBoard uint64, whiteToMove bool) uint64 {
	// Find first hit square
	upRightMoves := upRightBoard & usefulBitboards.AllOccupiedSquares

	// Fill all squares up and right by performing left shifts
	upRightMoves = (upRightMoves << 9) | (upRightMoves << 18) | (upRightMoves << 27) | (upRightMoves << 36) | (upRightMoves << 45) | (upRightMoves << 54)

	// Remove overflow
	upRightMoves &= upRightBoard

	// Get just the allowed squares using XOR
	upRightMoves ^= upRightBoard

	// Remove the blocking piece if it can't be captured (i.e. It is a friendly piece)
	if whiteToMove {
		upRightMoves &= usefulBitboards.BlackOrEmpty
	} else {
		upRightMoves &= usefulBitboards.WhiteOrEmpty
	}

	return upRightMoves
}

func calculateAllowedDownRightMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	downRightBoard := lookuptables.DownRightBoard[pieceIndex]
	return calculateAllowedDownRightMovesFromBoard(usefulBitboards, downRightBoard, whiteToMove)
}

func calculateAllowedDownRightMovesFromBoard(usefulBitboards *boardrepresentation.UsefulBitboards, downRightBoard uint64, whiteToMove bool) uint64 {
	// Find first hit square
	downRightMoves := downRightBoard & usefulBitboards.AllOccupiedSquares

	// Fill all squares down and right by performing right shifts
	downRightMoves = (downRightMoves >> 7) | (downRightMoves >> 14) | (downRightMoves >> 21) | (downRightMoves >> 28) | (downRightMoves >> 35) | (downRightMoves >> 42)

	// Remove overflow
	downRightMoves &= downRightBoard

	// Get just the allowed squares using XOR
	downRightMoves ^= downRightBoard

	// Remove the blocking piece if it can't be captured (i.e. It is a friendly piece)
	if whiteToMove {
		downRightMoves &= usefulBitboards.BlackOrEmpty
	} else {
		downRightMoves &= usefulBitboards.WhiteOrEmpty
	}

	return downRightMoves
}

func calculateAllowedDownLeftMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	downLeftBoard := lookuptables.DownLeftBoard[pieceIndex]
	return calculateAllowedDownLeftMovesFromBoard(usefulBitboards, downLeftBoard, whiteToMove)
}

func calculateAllowedDownLeftMovesFromBoard(usefulBitboards *boardrepresentation.UsefulBitboards, downLeftBoard uint64, whiteToMove bool) uint64 {
	// Find first hit square
	downLeftMoves := downLeftBoard & usefulBitboards.AllOccupiedSquares

	// Fill all squares down and left by performing right shifts
	downLeftMoves = (downLeftMoves >> 9) | (downLeftMoves >> 18) | (downLeftMoves >> 27) | (downLeftMoves >> 36) | (downLeftMoves >> 45) | (downLeftMoves >> 54)

	// Remove overflow
	downLeftMoves &= downLeftBoard

	// Get just the allowed squares using XOR
	downLeftMoves ^= downLeftBoard

	// Remove the blocking piece if it can't be captured (i.e. It is a friendly piece)
	if whiteToMove {
		downLeftMoves &= usefulBitboards.BlackOrEmpty
	} else {
		downLeftMoves &= usefulBitboards.WhiteOrEmpty
	}

	return downLeftMoves
}

func calculateAllowedUpLeftMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	upLeftBoard := lookuptables.UpLeftBoard[pieceIndex]
	return calculateAllowedUpLeftMovesFromBoard(usefulBitboards, upLeftBoard, whiteToMove)
}

func calculateAllowedUpLeftMovesFromBoard(usefulBitboards *boardrepresentation.UsefulBitboards, upLeftBoard uint64, whiteToMove bool) uint64 {
	// Find first hit square
	upLeftMoves := upLeftBoard & usefulBitboards.AllOccupiedSquares

	// Fill all squares up and left by performing right shifts
	upLeftMoves = (upLeftMoves << 7) | (upLeftMoves << 14) | (upLeftMoves << 21) | (upLeftMoves << 28) | (upLeftMoves << 35) | (upLeftMoves << 42)

	// Remove overflow
	upLeftMoves &= upLeftBoard

	// Get just the allowed squares using XOR
	upLeftMoves ^= upLeftBoard

	// Remove the blocking piece if it can't be captured (i.e. It is a friendly piece)
	if whiteToMove {
		upLeftMoves &= usefulBitboards.BlackOrEmpty
	} else {
		upLeftMoves &= usefulBitboards.WhiteOrEmpty
	}

	return upLeftMoves
}

func calculateAllowedUpMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	upBoard := lookuptables.UpBoard[pieceIndex]
	return calculateAllowedUpMovesFromBoard(usefulBitboards, upBoard, whiteToMove)
}

func calculateAllowedUpMovesFromBoard(usefulBitboards *boardrepresentation.UsefulBitboards, upBoard uint64, whiteToMove bool) uint64 {
	// Find first hit square
	upMoves := upBoard & usefulBitboards.AllOccupiedSquares

	// Fill all squares up by performing left shifts
	upMoves = (upMoves << 8) | (upMoves << 16) | (upMoves << 24) | (upMoves << 32) | (upMoves << 40) | (upMoves << 48)

	// Get just the allowed squares using XOR
	upMoves ^= upBoard

	// Remove the blocking piece if it can't be captured (i.e. It is a friendly piece)
	if whiteToMove {
		upMoves &= usefulBitboards.BlackOrEmpty
	} else {
		upMoves &= usefulBitboards.WhiteOrEmpty
	}

	return upMoves
}

func calculateAllowedRightMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	rightBoard := lookuptables.RightBoard[pieceIndex]
	return calculateAllowedRightMovesFromBoard(usefulBitboards, rightBoard, whiteToMove)
}

func calculateAllowedRightMovesFromBoard(usefulBitboards *boardrepresentation.UsefulBitboards, rightBoard uint64, whiteToMove bool) uint64 {
	// Find first hit square
	rightMoves := rightBoard & usefulBitboards.AllOccupiedSquares

	// Fill all right squares up by performing left shifts
	rightMoves = (rightMoves << 1) | (rightMoves << 2) | (rightMoves << 3) | (rightMoves << 4) | (rightMoves << 5) | (rightMoves << 6)

	// Remove overflow
	rightMoves &= rightBoard

	// Get just the allowed squares using XOR
	rightMoves ^= rightBoard

	// Remove the blocking piece if it can't be captured (i.e. It is a friendly piece)
	if whiteToMove {
		rightMoves &= usefulBitboards.BlackOrEmpty
	} else {
		rightMoves &= usefulBitboards.WhiteOrEmpty
	}

	return rightMoves
}

func calculateAllowedDownMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	downBoard := lookuptables.DownBoard[pieceIndex]
	return calculateAllowedDownMovesFromBoard(usefulBitboards, downBoard, whiteToMove)
}

func calculateAllowedDownMovesFromBoard(usefulBitboards *boardrepresentation.UsefulBitboards, downBoard uint64, whiteToMove bool) uint64 {
	// Find first hit square
	downMoves := downBoard & usefulBitboards.AllOccupiedSquares

	// Fill all down squares up by performing right shifts
	downMoves = (downMoves >> 8) | (downMoves >> 16) | (downMoves >> 24) | (downMoves >> 32) | (downMoves >> 40) | (downMoves >> 48)

	// Get just the allowed squares using XOR
	downMoves ^= downBoard

	// Remove the blocking piece if it can't be captured (i.e. It is a friendly piece)
	if whiteToMove {
		downMoves &= usefulBitboards.BlackOrEmpty
	} else {
		downMoves &= usefulBitboards.WhiteOrEmpty
	}

	return downMoves
}

func calculateAllowedLeftMoves(usefulBitboards *boardrepresentation.UsefulBitboards, pieceIndex byte, whiteToMove bool) uint64 {
	leftBoard := lookuptables.LeftBoard[pieceIndex]
	return calculateAllowedLeftMovesFromBoard(usefulBitboards, leftBoard, whiteToMove)
}

func calculateAllowedLeftMovesFromBoard(usefulBitboards *boardrepresentation.UsefulBitboards, leftBoard uint64, whiteToMove bool) uint64 {
	// Find first hit square
	leftMoves := leftBoard & usefulBitboards.AllOccupiedSquares

	// Fill all left squares up by performing right shifts
	leftMoves = (leftMoves >> 1) | (leftMoves >> 2) | (leftMoves >> 3) | (leftMoves >> 4) | (leftMoves >> 5) | (leftMoves >> 6)

	// Remove overflow
	leftMoves &= leftBoard

	// Get just the allowed squares using XOR
	leftMoves ^= leftBoard

	// Remove the blocking piece if it can't be captured (i.e. It is a friendly piece)
	if whiteToMove {
		leftMoves &= usefulBitboards.BlackOrEmpty
	} else {
		leftMoves &= usefulBitboards.WhiteOrEmpty
	}

	return leftMoves
}

package fentranslator

import (
	"strings"

	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
)

// Fen notation - "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
func toGameBoard(fenString string, gameBoard *boardrepresentation.GameBoard) {
	parts := strings.Fields(fenString)

	// Split fen string
	// [0] board

	// Side to move
	if parts[1] == "w" {
		gameBoard.WhiteToMove = true
	} else {
		gameBoard.WhiteToMove = false
	}

	// [2] Castling status
	toCastlingStatus(parts[2], gameBoard)

	// [3] half move clock
	// [4] Full move clock
}

func toCastlingStatus(castlingFen string, gameBoard *boardrepresentation.GameBoard) {
	gameBoard.WhiteCanCastleKingside = false
	gameBoard.BlackCanCastleKingside = false
	gameBoard.WhiteCanCastleQueenside = false
	gameBoard.BlackCanCastleQueenside = false

	if strings.Contains(castlingFen, "K") {
		gameBoard.WhiteCanCastleKingside = true
	}

	if strings.Contains(castlingFen, "k") {
		gameBoard.BlackCanCastleKingside = true
	}

	if strings.Contains(castlingFen, "Q") {
		gameBoard.WhiteCanCastleQueenside = true
	}

	if strings.Contains(castlingFen, "q") {
		gameBoard.BlackCanCastleQueenside = true
	}
}

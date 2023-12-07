package fentranslator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
	"github.com/terrysmalone/chess-move-generator/boardsearching/lookuptables"
)

// Fen notation example - "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
func toGameBoard(fenString string, gameBoard *boardrepresentation.GameBoard) error {
	parts := strings.Fields(fenString)

	// Split fen string
	// [0] board

	// Side to move
	if parts[1] == "w" {
		gameBoard.WhiteToMove = true
	} else {
		gameBoard.WhiteToMove = false
	}

	toCastlingStatus(parts[2], gameBoard)

	err := toEnPassantPosition(parts[3], gameBoard)
	if err != nil {
		return fmt.Errorf("error parsing En passant position: %w", err)
	}

	// Half move clock
	halfMove, err := strconv.Atoi(parts[4])
	if err != nil {
		return fmt.Errorf("error parsing half move clock:%w", err)
	}
	gameBoard.HalfMoveClock = halfMove

	// Full move clock
	fullMove, err := strconv.Atoi(parts[5])
	if err != nil {
		return fmt.Errorf("error parsing full move clock:%w", err)
	}
	gameBoard.FullMoveClock = fullMove

	return nil
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

func toEnPassantPosition(enPassantFen string, gameBoard *boardrepresentation.GameBoard) error {
	if enPassantFen != "-" {
		if len(enPassantFen) != 2 {
			return fmt.Errorf("there should only be two characters in en passant position. There are %d", len(enPassantFen))
		}

		chars := []rune(enPassantFen)

		column := chars[0]
		var columnNum int

		switch column {
		case 'a':
			columnNum = 0
		case 'b':
			columnNum = 1
		case 'c':
			columnNum = 2
		case 'd':
			columnNum = 3
		case 'e':
			columnNum = 4
		case 'f':
			columnNum = 5
		case 'g':
			columnNum = 6
		case 'h':
			columnNum = 7
		default:
			return fmt.Errorf("en passant column %d not recognised", column)
		}

		rowNum := int(chars[1]-'0') - 1

		gameBoard.EnPassantPosition = lookuptables.BitboardValueFromPosition[columnNum][rowNum]
	} else {
		gameBoard.EnPassantPosition = 0
	}

	return nil
}

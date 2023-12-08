package fentranslator

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
	"github.com/terrysmalone/chess-move-generator/boardsearching/lookuptables"
)

// Fen notation example - "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
func toGameBoard(fenString string, gameBoard *boardrepresentation.GameBoard) error {
	parts := strings.Fields(fenString)

	if len(parts) != 6 {
		return fmt.Errorf("there should be 6 parts to a fen notation string. There are %d", len(parts))
	}

	if err := toBoard(parts[0], gameBoard); err != nil {
		return fmt.Errorf("error parsing fen string to board: %w", err)
	}

	// Side to move
	if parts[1] == "w" {
		gameBoard.WhiteToMove = true
	} else {
		gameBoard.WhiteToMove = false
	}

	toCastlingStatus(parts[2], gameBoard)

	if err := toEnPassantPosition(parts[3], gameBoard); err != nil {
		return fmt.Errorf("error parsing En passant position: %w", err)
	}

	// Half move clock
	halfMove, err := strconv.Atoi(parts[4])
	if err != nil {
		return fmt.Errorf("error parsing half move clock: %w", err)
	}
	gameBoard.HalfMoveClock = halfMove

	// Full move clock
	fullMove, err := strconv.Atoi(parts[5])
	if err != nil {
		return fmt.Errorf("error parsing full move clock: %w", err)
	}
	gameBoard.FullMoveClock = fullMove

	return nil
}

func toBoard(boardFEN string, gameBoard *boardrepresentation.GameBoard) error {
	rows := strings.Split(boardFEN, "/")

	if len(rows) != 8 {
		return fmt.Errorf("expected 8 rows. There are %d", len(rows))
	}

	for r := 0; r < 8; r++ {
		pieces := []rune(rows[r])

		currentRow := 7 - r
		currentColumn := 0
		for _, piece := range pieces {
			if unicode.IsLetter(piece) {
				// Get piece from letter
				switch piece {
				case 'P':
					gameBoard.Board.WhitePawns |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				case 'N':
					gameBoard.Board.WhiteKnights |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				case 'B':
					gameBoard.Board.WhiteBishops |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				case 'R':
					gameBoard.Board.WhiteRooks |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				case 'Q':
					gameBoard.Board.WhiteQueens |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				case 'K':
					gameBoard.Board.WhiteKing |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				case 'p':
					gameBoard.Board.BlackPawns |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				case 'n':
					gameBoard.Board.BlackKnights |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				case 'b':
					gameBoard.Board.BlackBishops |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				case 'r':
					gameBoard.Board.BlackRooks |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				case 'q':
					gameBoard.Board.BlackQueens |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				case 'k':
					gameBoard.Board.BlackKing |= lookuptables.BitboardValueFromPosition[currentColumn][currentRow]
				default:
					return fmt.Errorf("piece %s not recognised", strconv.QuoteRune(piece))
				}
				currentColumn += 1
			} else if unicode.IsNumber(piece) {
				currentColumn += int(piece - '0')
			} else {
				return fmt.Errorf("piece %s should be a letter or number", strconv.QuoteRune(piece))
			}
		}
	}

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
		if !unicode.IsLetter(chars[0]) {
			return fmt.Errorf("en passant column must be a letter. It is %s", strconv.QuoteRune(column))
		}
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
			return fmt.Errorf("en passant column %s not recognised", strconv.QuoteRune(column))
		}

		if !unicode.IsNumber(chars[1]) {
			return fmt.Errorf("en passant row must be a number. It is %s", strconv.QuoteRune(chars[1]))
		}

		rowNum := int(chars[1]-'0') - 1

		if rowNum > 7 {
			return fmt.Errorf("en passant row %s not recognised", strconv.QuoteRune(chars[1]))
		}

		gameBoard.EnPassantPosition = lookuptables.BitboardValueFromPosition[columnNum][rowNum]
	} else {
		gameBoard.EnPassantPosition = 0
	}

	return nil
}

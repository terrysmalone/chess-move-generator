package boardrepresentation

import "math"

// Minimum representation of a board in a game of chess
type GameBoard struct {
	Board           Board
	UsefulBitboards UsefulBitboards

	WhiteCanCastleKingside  bool
	WhiteCanCastleQueenside bool

	BlackCanCastleKingside  bool
	BlackCanCastleQueenside bool

	EnPassantPosition uint64

	WhiteToMove bool

	HalfMoveClock int
	FullMoveClock int
}

func (g *GameBoard) initialiseStartingGamePosition() {
	g.Board.initialiseStartingPosition()

	g.WhiteCanCastleKingside = true
	g.WhiteCanCastleQueenside = true

	g.BlackCanCastleKingside = true
	g.BlackCanCastleQueenside = true

	g.EnPassantPosition = 0

	g.WhiteToMove = true

	g.HalfMoveClock = 0
	g.FullMoveClock = 0
}

func (g *GameBoard) CalculateUsefulBitboards() {
	g.UsefulBitboards.AllWhiteOccupiedSquares = g.Board.WhitePawns | g.Board.WhiteKnights | g.Board.WhiteBishops | g.Board.WhiteRooks | g.Board.WhiteQueens | g.Board.WhiteKing
	g.UsefulBitboards.AllBlackOccupiedSquares = g.Board.BlackPawns | g.Board.BlackKnights | g.Board.BlackBishops | g.Board.BlackRooks | g.Board.BlackQueens | g.Board.BlackKing
	g.UsefulBitboards.AllOccupiedSquares = g.UsefulBitboards.AllWhiteOccupiedSquares | g.UsefulBitboards.AllBlackOccupiedSquares
	g.UsefulBitboards.EmptySquares = g.UsefulBitboards.AllOccupiedSquares ^ math.MaxUint64

	g.UsefulBitboards.WhiteOrEmpty = g.UsefulBitboards.AllWhiteOccupiedSquares | g.UsefulBitboards.EmptySquares
	g.UsefulBitboards.BlackOrEmpty = g.UsefulBitboards.AllBlackOccupiedSquares | g.UsefulBitboards.EmptySquares
}

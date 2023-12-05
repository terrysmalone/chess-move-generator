package boardrepresentation

// Minimum representation of a board in a game of chess
type GameBoard struct {
	Board Board

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

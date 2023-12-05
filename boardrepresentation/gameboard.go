package boardrepresentation

// Minimum representation of a board in a game of chess
type GameBoard struct {
	board Board

	whiteCanCastleKingside  bool
	whiteCanCastleQueenside bool

	blackCanCastleKingside  bool
	blackCanCastleQueenside bool

	enPassantPosition uint64

	whiteToMove bool
}

func (g *GameBoard) initialiseStartingGamePosition() {
	g.board.initialiseStartingPosition()

	g.whiteCanCastleKingside = true
	g.whiteCanCastleQueenside = true

	g.blackCanCastleKingside = true
	g.blackCanCastleQueenside = true

	g.enPassantPosition = 0

	g.whiteToMove = true
}

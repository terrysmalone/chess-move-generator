package boardrepresentation

// Minimum representation of all of the game pieces on a board
type Board struct {
	WhitePawns   uint64
	WhiteKnights uint64
	WhiteBishops uint64
	WhiteRooks   uint64
	WhiteQueens  uint64
	WhiteKing    uint64

	BlackPawns   uint64
	BlackKnights uint64
	BlackBishops uint64
	BlackRooks   uint64
	BlackQueens  uint64
	BlackKing    uint64
}

func (b *Board) initialiseStartingPosition() {
	b.WhitePawns = 65280
	b.WhiteKnights = 66
	b.WhiteBishops = 36
	b.WhiteRooks = 129
	b.WhiteQueens = 8
	b.WhiteKing = 16

	b.BlackPawns = 71776119061217280
	b.BlackKnights = 4755801206503243776
	b.BlackBishops = 2594073385365405696
	b.BlackRooks = 9295429630892703744
	b.BlackQueens = 576460752303423488
	b.BlackKing = 1152921504606846976
}

func (b *Board) clearBoard() {
	b.WhitePawns = 0
	b.WhiteKnights = 0
	b.WhiteBishops = 0
	b.WhiteRooks = 0
	b.WhiteQueens = 0
	b.WhiteKing = 0

	b.BlackPawns = 0
	b.BlackKnights = 0
	b.BlackBishops = 0
	b.BlackRooks = 0
	b.BlackQueens = 0
	b.BlackKing = 0
}

package boardrepresentation

// Minimum representation of all of the game pieces on a board
type Board struct {
	whitePawns   uint64
	whiteKnights uint64
	whiteBishops uint64
	whiteRooks   uint64
	whiteQueens  uint64
	whiteKing    uint64

	blackPawns   uint64
	blackKnights uint64
	blackBishops uint64
	blackRooks   uint64
	blackQueens  uint64
	blackKing    uint64
}

func (b *Board) initialiseStartingPosition() {
	b.whitePawns = 65280
	b.whiteKnights = 66
	b.whiteBishops = 36
	b.whiteRooks = 129
	b.whiteQueens = 8
	b.whiteKing = 16

	b.blackPawns = 71776119061217280
	b.blackKnights = 4755801206503243776
	b.blackBishops = 2594073385365405696
	b.blackRooks = 9295429630892703744
	b.blackQueens = 576460752303423488
	b.blackKing = 1152921504606846976
}

func (b *Board) clearBoard() {
	b.whitePawns = 0
	b.whiteKnights = 0
	b.whiteBishops = 0
	b.whiteRooks = 0
	b.whiteQueens = 0
	b.whiteKing = 0

	b.blackPawns = 0
	b.blackKnights = 0
	b.blackBishops = 0
	b.blackRooks = 0
	b.blackQueens = 0
	b.blackKing = 0
}

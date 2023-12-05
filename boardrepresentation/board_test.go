package boardrepresentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialiseStartingPosition(t *testing.T) {
	board := &Board{}

	board.initialiseStartingPosition()

	assert.EqualValues(t, 65280, board.WhitePawns)
	assert.EqualValues(t, 66, board.WhiteKnights)
	assert.EqualValues(t, 36, board.WhiteBishops)
	assert.EqualValues(t, 129, board.WhiteRooks)
	assert.EqualValues(t, 8, board.WhiteQueens)
	assert.EqualValues(t, 16, board.WhiteKing)

	assert.EqualValues(t, 71776119061217280, board.BlackPawns)
	assert.EqualValues(t, 4755801206503243776, board.BlackKnights)
	assert.EqualValues(t, 2594073385365405696, board.BlackBishops)
	assert.EqualValues(t, uint64(9295429630892703744), board.BlackRooks)
	assert.EqualValues(t, 576460752303423488, board.BlackQueens)
	assert.EqualValues(t, 1152921504606846976, board.BlackKing)
}

func TestClearBoard(t *testing.T) {
	board := &Board{}

	board.initialiseStartingPosition()
	// Just check it's not clear
	assert.EqualValues(t, 71776119061217280, board.BlackPawns)

	board.clearBoard()

	assert.EqualValues(t, 0, board.WhitePawns)
	assert.EqualValues(t, 0, board.WhiteKnights)
	assert.EqualValues(t, 0, board.WhiteBishops)
	assert.EqualValues(t, 0, board.WhiteRooks)
	assert.EqualValues(t, 0, board.WhiteQueens)
	assert.EqualValues(t, 0, board.WhiteKing)

	assert.EqualValues(t, 0, board.BlackPawns)
	assert.EqualValues(t, 0, board.BlackKnights)
	assert.EqualValues(t, 0, board.BlackBishops)
	assert.EqualValues(t, 0, board.BlackRooks)
	assert.EqualValues(t, 0, board.BlackQueens)
	assert.EqualValues(t, 0, board.BlackKing)
}

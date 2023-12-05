package boardrepresentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialiseStartingPosition(t *testing.T) {
	board := &Board{}

	board.initialiseStartingPosition()

	assert.EqualValues(t, 65280, board.whitePawns)
	assert.EqualValues(t, 66, board.whiteKnights)
	assert.EqualValues(t, 36, board.whiteBishops)
	assert.EqualValues(t, 129, board.whiteRooks)
	assert.EqualValues(t, 8, board.whiteQueens)
	assert.EqualValues(t, 16, board.whiteKing)

	assert.EqualValues(t, 71776119061217280, board.blackPawns)
	assert.EqualValues(t, 4755801206503243776, board.blackKnights)
	assert.EqualValues(t, 2594073385365405696, board.blackBishops)
	assert.EqualValues(t, uint64(9295429630892703744), board.blackRooks)
	assert.EqualValues(t, 576460752303423488, board.blackQueens)
	assert.EqualValues(t, 1152921504606846976, board.blackKing)
}

func TestClearBoard(t *testing.T) {
	board := &Board{}

	board.initialiseStartingPosition()
	// Just check it's not clear
	assert.EqualValues(t, 71776119061217280, board.blackPawns)

	board.clearBoard()

	assert.EqualValues(t, 0, board.whitePawns)
	assert.EqualValues(t, 0, board.whiteKnights)
	assert.EqualValues(t, 0, board.whiteBishops)
	assert.EqualValues(t, 0, board.whiteRooks)
	assert.EqualValues(t, 0, board.whiteQueens)
	assert.EqualValues(t, 0, board.whiteKing)

	assert.EqualValues(t, 0, board.blackPawns)
	assert.EqualValues(t, 0, board.blackKnights)
	assert.EqualValues(t, 0, board.blackBishops)
	assert.EqualValues(t, 0, board.blackRooks)
	assert.EqualValues(t, 0, board.blackQueens)
	assert.EqualValues(t, 0, board.blackKing)
}

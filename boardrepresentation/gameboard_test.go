package boardrepresentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialiseStartingGamePosition(t *testing.T) {
	gameBoard := &GameBoard{}

	gameBoard.initialiseStartingGamePosition()

	assert.EqualValues(t, 65280, gameBoard.Board.WhitePawns)
	assert.EqualValues(t, 576460752303423488, gameBoard.Board.BlackQueens)

	assert.Equal(t, true, gameBoard.WhiteCanCastleKingside)
	assert.Equal(t, true, gameBoard.WhiteCanCastleQueenside)
	assert.Equal(t, true, gameBoard.BlackCanCastleKingside)
	assert.Equal(t, true, gameBoard.BlackCanCastleQueenside)

	assert.EqualValues(t, 0, gameBoard.EnPassantPosition)

	assert.Equal(t, true, gameBoard.WhiteToMove)

	assert.Equal(t, 0, gameBoard.HalfMoveClock)
	assert.Equal(t, 0, gameBoard.FullMoveClock)
}

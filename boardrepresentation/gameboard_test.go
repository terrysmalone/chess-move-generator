package boardrepresentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialiseStartingGamePosition(t *testing.T) {
	gameBoard := &GameBoard{}

	gameBoard.initialiseStartingGamePosition()

	assert.EqualValues(t, 65280, gameBoard.board.whitePawns)
	assert.EqualValues(t, 576460752303423488, gameBoard.board.blackQueens)

	assert.Equal(t, true, gameBoard.whiteCanCastleKingside)
	assert.Equal(t, true, gameBoard.whiteCanCastleQueenside)
	assert.Equal(t, true, gameBoard.blackCanCastleKingside)
	assert.Equal(t, true, gameBoard.blackCanCastleQueenside)

	assert.EqualValues(t, 0, gameBoard.enPassantPosition)

	assert.Equal(t, true, gameBoard.whiteToMove)
}

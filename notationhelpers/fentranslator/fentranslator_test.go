package fentranslator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
)

func TestToGameBoard_WhiteToMove(t *testing.T) {
	gameBoard := &boardrepresentation.GameBoard{}

	toGameBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 0", gameBoard)
	assert.Equal(t, false, gameBoard.WhiteToMove)

	toGameBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 0", gameBoard)
	assert.Equal(t, true, gameBoard.WhiteToMove)
}

func TestToGameBoard_CastlingStatus(t *testing.T) {
	tests := []struct {
		name                   string
		fenNotation            string
		expectedWhiteKingside  bool
		expectedWhiteQueenside bool
		expectedBlackKingside  bool
		expectedBlackQueenside bool
	}{
		{
			name:                   "All can castle",
			fenNotation:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 0",
			expectedWhiteKingside:  true,
			expectedWhiteQueenside: true,
			expectedBlackKingside:  true,
			expectedBlackQueenside: true,
		},
		{
			name:                   "None can castle",
			fenNotation:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b - - 0 0",
			expectedWhiteKingside:  false,
			expectedWhiteQueenside: false,
			expectedBlackKingside:  false,
			expectedBlackQueenside: false,
		},
		{
			name:                   "Only white can castle",
			fenNotation:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQ - 0 0",
			expectedWhiteKingside:  true,
			expectedWhiteQueenside: true,
			expectedBlackKingside:  false,
			expectedBlackQueenside: false,
		},
		{
			name:                   "Only black can castle",
			fenNotation:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b kq - 0 0",
			expectedWhiteKingside:  false,
			expectedWhiteQueenside: false,
			expectedBlackKingside:  true,
			expectedBlackQueenside: true,
		},
		{
			name:                   "Only kingside can castle",
			fenNotation:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b Kk - 0 0",
			expectedWhiteKingside:  true,
			expectedWhiteQueenside: false,
			expectedBlackKingside:  true,
			expectedBlackQueenside: false,
		},
		{
			name:                   "Only queenside can castle",
			fenNotation:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b Qq - 0 0",
			expectedWhiteKingside:  false,
			expectedWhiteQueenside: true,
			expectedBlackKingside:  false,
			expectedBlackQueenside: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameBoard := &boardrepresentation.GameBoard{}

			toGameBoard(tt.fenNotation, gameBoard)
			assert.Equal(t, tt.expectedWhiteKingside, gameBoard.WhiteCanCastleKingside)
			assert.Equal(t, tt.expectedWhiteQueenside, gameBoard.WhiteCanCastleQueenside)
			assert.Equal(t, tt.expectedBlackKingside, gameBoard.BlackCanCastleKingside)
			assert.Equal(t, tt.expectedBlackKingside, gameBoard.BlackCanCastleKingside)
		})
	}

}

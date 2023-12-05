package fentranslator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
)

func TestToGameBoard_WhiteToMove(t *testing.T) {
	gameBoard := &boardrepresentation.GameBoard{}

	err := toGameBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 0", gameBoard)

	require.NoError(t, err)
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

			err := toGameBoard(tt.fenNotation, gameBoard)
			require.NoError(t, err)

			assert.Equal(t, tt.expectedWhiteKingside, gameBoard.WhiteCanCastleKingside)
			assert.Equal(t, tt.expectedWhiteQueenside, gameBoard.WhiteCanCastleQueenside)
			assert.Equal(t, tt.expectedBlackKingside, gameBoard.BlackCanCastleKingside)
			assert.Equal(t, tt.expectedBlackKingside, gameBoard.BlackCanCastleKingside)
		})
	}
}

func TestToGameBoard_MoveClock(t *testing.T) {
	tests := []struct {
		name                  string
		fenNotation           string
		expectedHalfMoveClock int
		expectedFullMoveClock int
		expectedError         error
	}{
		{
			name:                  "Moves are right",
			fenNotation:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 3 6",
			expectedHalfMoveClock: 3,
			expectedFullMoveClock: 6,
		},
		{
			name:                  "Half clock error",
			fenNotation:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - ds 6",
			expectedHalfMoveClock: 0,
			expectedFullMoveClock: 0,
			expectedError:         fmt.Errorf("strconv.Atoi: parsing \"ds\": invalid syntax"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameBoard := &boardrepresentation.GameBoard{}

			err := toGameBoard(tt.fenNotation, gameBoard)

			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.expectedHalfMoveClock, gameBoard.HalfMoveClock)
			assert.Equal(t, tt.expectedFullMoveClock, gameBoard.FullMoveClock)
		})
	}
}

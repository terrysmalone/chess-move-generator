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

func TestToCastlingStatus(t *testing.T) {
	tests := []struct {
		name                   string
		fenCastlingStatus      string
		expectedWhiteKingside  bool
		expectedWhiteQueenside bool
		expectedBlackKingside  bool
		expectedBlackQueenside bool
	}{
		{
			name:                   "All can castle",
			fenCastlingStatus:      "KQkq",
			expectedWhiteKingside:  true,
			expectedWhiteQueenside: true,
			expectedBlackKingside:  true,
			expectedBlackQueenside: true,
		},
		{
			name:                   "None can castle",
			fenCastlingStatus:      "-",
			expectedWhiteKingside:  false,
			expectedWhiteQueenside: false,
			expectedBlackKingside:  false,
			expectedBlackQueenside: false,
		},
		{
			name:                   "Only white can castle",
			fenCastlingStatus:      "KQ",
			expectedWhiteKingside:  true,
			expectedWhiteQueenside: true,
			expectedBlackKingside:  false,
			expectedBlackQueenside: false,
		},
		{
			name:                   "Only black can castle",
			fenCastlingStatus:      "kq",
			expectedWhiteKingside:  false,
			expectedWhiteQueenside: false,
			expectedBlackKingside:  true,
			expectedBlackQueenside: true,
		},
		{
			name:                   "Only kingside can castle",
			fenCastlingStatus:      "Kk",
			expectedWhiteKingside:  true,
			expectedWhiteQueenside: false,
			expectedBlackKingside:  true,
			expectedBlackQueenside: false,
		},
		{
			name:                   "Only queenside can castle",
			fenCastlingStatus:      "Qq",
			expectedWhiteKingside:  false,
			expectedWhiteQueenside: true,
			expectedBlackKingside:  false,
			expectedBlackQueenside: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameBoard := &boardrepresentation.GameBoard{}

			toCastlingStatus(tt.fenCastlingStatus, gameBoard)

			assert.Equal(t, tt.expectedWhiteKingside, gameBoard.WhiteCanCastleKingside)
			assert.Equal(t, tt.expectedWhiteQueenside, gameBoard.WhiteCanCastleQueenside)
			assert.Equal(t, tt.expectedBlackKingside, gameBoard.BlackCanCastleKingside)
			assert.Equal(t, tt.expectedBlackKingside, gameBoard.BlackCanCastleKingside)
		})
	}
}

func TestToEnPassantPosition(t *testing.T) {
	tests := []struct {
		name                      string
		fenEnPassantPosition      string
		expectedEnPassantPosition uint64
		expectedError             error
	}{
		{
			name:                      "Invalid text",
			fenEnPassantPosition:      "invalid",
			expectedEnPassantPosition: 0,
			expectedError:             fmt.Errorf("there should only be two characters in en passant position. There are 7"),
		},
		{
			name:                      "No en passant position",
			fenEnPassantPosition:      "-",
			expectedEnPassantPosition: 0,
		},
		{
			name:                      "Set a position",
			fenEnPassantPosition:      "e3",
			expectedEnPassantPosition: 1048576,
		},
		{
			name:                      "Set a high uint position",
			fenEnPassantPosition:      "h6",
			expectedEnPassantPosition: 140737488355328,
		},
		{
			name:                      "Set to a8",
			fenEnPassantPosition:      "a8",
			expectedEnPassantPosition: 72057594037927936,
		},
		{
			name:                      "Set to a1",
			fenEnPassantPosition:      "a1",
			expectedEnPassantPosition: 1,
		},
		{
			name:                      "Set to h8",
			fenEnPassantPosition:      "h8",
			expectedEnPassantPosition: 9223372036854775808,
		},
		{
			name:                      "Set to h1",
			fenEnPassantPosition:      "h1",
			expectedEnPassantPosition: 128,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameBoard := &boardrepresentation.GameBoard{}

			// Set it to something so we can be sure it changes
			gameBoard.EnPassantPosition = 70368744177664

			err := toEnPassantPosition(tt.fenEnPassantPosition, gameBoard)

			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)

				assert.Equal(t, tt.expectedEnPassantPosition, gameBoard.EnPassantPosition)
			}
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
			name:                  "Half move clock error",
			fenNotation:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - ds 6",
			expectedHalfMoveClock: 0,
			expectedFullMoveClock: 0,
			expectedError:         fmt.Errorf("error parsing half move clock:strconv.Atoi: parsing \"ds\": invalid syntax"),
		},
		{
			name:                  "Full move clock error",
			fenNotation:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 2 fev",
			expectedHalfMoveClock: 2,
			expectedFullMoveClock: 0,
			expectedError:         fmt.Errorf("error parsing full move clock:strconv.Atoi: parsing \"fev\": invalid syntax"),
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

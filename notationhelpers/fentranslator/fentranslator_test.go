package fentranslator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
)

func TestToGameBoardErrors(t *testing.T) {
	tests := []struct {
		name          string
		fenNotation   string
		expectedError error
	}{
		{
			name:          "Too few parts",
			fenNotation:   "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - 0 1",
			expectedError: fmt.Errorf("there should be 6 parts to a fen notation string. There are 5"),
		},
		{
			name:          "Too many parts",
			fenNotation:   "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq EXTRA - 0 1",
			expectedError: fmt.Errorf("there should be 6 parts to a fen notation string. There are 7"),
		},
		{
			name:          "To board error",
			fenNotation:   "rnbqkbnr/pBQKBNR w KQkq - 0 1",
			expectedError: fmt.Errorf("error parsing fen string to board: expected 8 rows. There are 2"),
		},
		{
			name:          "En Passant error",
			fenNotation:   "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq WRONG 0 1",
			expectedError: fmt.Errorf("error parsing En passant position: there should only be two characters in en passant position. There are 5"),
		},
		{
			name:          "To board error",
			fenNotation:   "rnbqkXnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq WRONG 0 1",
			expectedError: fmt.Errorf("error parsing fen string to board: piece 'X' not recognised"),
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
		})
	}
}

func TestToGameBoard_WhiteToMove(t *testing.T) {
	gameBoard := &boardrepresentation.GameBoard{}

	err := toGameBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 0", gameBoard)

	require.NoError(t, err)
	assert.Equal(t, false, gameBoard.WhiteToMove)

	toGameBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 0", gameBoard)
	assert.Equal(t, true, gameBoard.WhiteToMove)
}

func TestToBoard(t *testing.T) {
	tests := []struct {
		name          string
		fenBoard      string
		expectedBoard boardrepresentation.Board
		expectedError error
	}{
		{
			name:     "White pawns only",
			fenBoard: "8/8/8/8/8/8/PPPPPPPP/8",
			expectedBoard: boardrepresentation.Board{
				WhitePawns: uint64(65280),
			},
			expectedError: nil,
		},
		{
			name:     "Black pawns only",
			fenBoard: "8/pppppppp/8/8/8/8/8/8",
			expectedBoard: boardrepresentation.Board{
				BlackPawns: uint64(71776119061217280),
			},
			expectedError: nil,
		},
		{
			name:     "White knights only",
			fenBoard: "8/8/8/8/8/8/8/1N4N1",
			expectedBoard: boardrepresentation.Board{
				WhiteKnights: uint64(66),
			},
			expectedError: nil,
		},
		{
			name:     "Black knights only",
			fenBoard: "1n4n1/8/8/8/8/8/8/8",
			expectedBoard: boardrepresentation.Board{
				BlackKnights: uint64(4755801206503243776),
			},
			expectedError: nil,
		},
		{
			name:     "White bishops only",
			fenBoard: "8/8/8/8/8/8/8/2B2B2",
			expectedBoard: boardrepresentation.Board{
				WhiteBishops: uint64(36),
			},
			expectedError: nil,
		},
		{
			name:     "Black bishops only",
			fenBoard: "2b2b2/8/8/8/8/8/8/8",
			expectedBoard: boardrepresentation.Board{
				BlackBishops: uint64(2594073385365405696),
			},
			expectedError: nil,
		},
		{
			name:     "White rooks only",
			fenBoard: "8/8/8/8/8/8/8/R6R",
			expectedBoard: boardrepresentation.Board{
				WhiteRooks: uint64(129),
			},
			expectedError: nil,
		},
		{
			name:     "Black rooks only",
			fenBoard: "r6r/8/8/8/8/8/8/8",
			expectedBoard: boardrepresentation.Board{
				BlackRooks: uint64(9295429630892703744),
			},
			expectedError: nil,
		},
		{
			name:     "White queens only",
			fenBoard: "8/8/8/8/8/8/8/3Q4",
			expectedBoard: boardrepresentation.Board{
				WhiteQueens: uint64(8),
			},
			expectedError: nil,
		},
		{
			name:     "Black queens only",
			fenBoard: "3q4/8/8/8/8/8/8/8",
			expectedBoard: boardrepresentation.Board{
				BlackQueens: uint64(576460752303423488),
			},
			expectedError: nil,
		},
		{
			name:     "White king only",
			fenBoard: "8/8/8/8/8/8/8/4K3",
			expectedBoard: boardrepresentation.Board{
				WhiteKing: uint64(16),
			},
			expectedError: nil,
		},
		{
			name:     "Black king only",
			fenBoard: "4k3/8/8/8/8/8/8/8",
			expectedBoard: boardrepresentation.Board{
				BlackKing: uint64(1152921504606846976),
			},
			expectedError: nil,
		},
		{
			name:     "All pieces starting position",
			fenBoard: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			expectedBoard: boardrepresentation.Board{
				WhitePawns:   uint64(65280),
				WhiteKnights: uint64(66),
				WhiteBishops: uint64(36),
				WhiteRooks:   uint64(129),
				WhiteQueens:  uint64(8),
				WhiteKing:    uint64(16),

				BlackPawns:   uint64(71776119061217280),
				BlackKnights: uint64(4755801206503243776),
				BlackBishops: uint64(2594073385365405696),
				BlackRooks:   uint64(9295429630892703744),
				BlackQueens:  uint64(576460752303423488),
				BlackKing:    uint64(1152921504606846976),
			},
			expectedError: nil,
		},
		{
			name:          "Unrecognised Piece",
			fenBoard:      "D7/8/8/8/8/8/8/8",
			expectedError: fmt.Errorf("piece 'D' not recognised"),
		},
		{
			name:          "Piece not letter or number",
			fenBoard:      "8/8/8/8/7£/8/8/8",
			expectedError: fmt.Errorf("piece '£' should be a letter or number"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameBoard := &boardrepresentation.GameBoard{}

			err := toBoard(tt.fenBoard, gameBoard)

			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.expectedBoard, gameBoard.Board)
		})
	}
}

func TestToEnPassantPosition(t *testing.T) {
	tests := []struct {
		name                      string
		fenEnPassantPosition      string
		expectedEnPassantPosition uint64
		expectedFullMoveClock     int
		expectedError             error
	}{
		{
			name:                      "No en passant position",
			fenEnPassantPosition:      "-",
			expectedEnPassantPosition: uint64(0),
		},
		{
			name:                      "a3",
			fenEnPassantPosition:      "a3",
			expectedEnPassantPosition: uint64(65536),
		},
		{
			name:                      "b6",
			fenEnPassantPosition:      "b6",
			expectedEnPassantPosition: uint64(2199023255552),
		},
		{
			name:                      "c3",
			fenEnPassantPosition:      "c3",
			expectedEnPassantPosition: uint64(262144),
		},
		{
			name:                      "d6",
			fenEnPassantPosition:      "d6",
			expectedEnPassantPosition: uint64(8796093022208),
		},
		{
			name:                      "e3",
			fenEnPassantPosition:      "e3",
			expectedEnPassantPosition: uint64(1048576),
		},
		{
			name:                      "f3",
			fenEnPassantPosition:      "f3",
			expectedEnPassantPosition: uint64(2097152),
		},
		{
			name:                      "g3",
			fenEnPassantPosition:      "g3",
			expectedEnPassantPosition: uint64(4194304),
		},
		{
			name:                      "h6",
			fenEnPassantPosition:      "h6",
			expectedEnPassantPosition: uint64(140737488355328),
		},

		{
			name:                 "Text too long",
			fenEnPassantPosition: "gtg",
			expectedError:        fmt.Errorf("there should only be two characters in en passant position. There are 3"),
		},
		{
			name:                 "Column text not a letter",
			fenEnPassantPosition: "34",
			expectedError:        fmt.Errorf("en passant column must be a letter. It is '3'"),
		},
		{
			name:                 "Column text invalid",
			fenEnPassantPosition: "j4",
			expectedError:        fmt.Errorf("en passant column 'j' not recognised"),
		},
		{
			name:                 "Row text not a number",
			fenEnPassantPosition: "ab",
			expectedError:        fmt.Errorf("en passant row must be a number. It is 'b'"),
		},
		{
			name:                 "Row text invalid",
			fenEnPassantPosition: "a9",
			expectedError:        fmt.Errorf("en passant row '9' not recognised"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameBoard := &boardrepresentation.GameBoard{}

			err := toEnPassantPosition(tt.fenEnPassantPosition, gameBoard)

			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.expectedEnPassantPosition, gameBoard.EnPassantPosition)
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
			expectedError:         fmt.Errorf("error parsing half move clock: strconv.Atoi: parsing \"ds\": invalid syntax"),
		},
		{
			name:                  "Full move clock error",
			fenNotation:           "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 2 fev",
			expectedHalfMoveClock: 2,
			expectedFullMoveClock: 0,
			expectedError:         fmt.Errorf("error parsing full move clock: strconv.Atoi: parsing \"fev\": invalid syntax"),
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

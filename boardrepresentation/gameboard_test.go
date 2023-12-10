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

func TestCalculateUsefulBitboards(t *testing.T) {
	tests := []struct {
		name                    string
		board                   Board
		expectedUsefulBitboards UsefulBitboards
	}{
		{
			name:  "Empty board",
			board: Board{},
			expectedUsefulBitboards: UsefulBitboards{
				AllWhiteOccupiedSquares: uint64(0),
				AllBlackOccupiedSquares: uint64(0),
				AllOccupiedSquares:      uint64(0),
				EmptySquares:            uint64(18446744073709551615),
			},
		},
		{
			name:  "Starting position",
			board: getStartingBoard(),
			expectedUsefulBitboards: UsefulBitboards{
				AllWhiteOccupiedSquares: uint64(65535),
				AllBlackOccupiedSquares: uint64(18446462598732840960),
				AllOccupiedSquares:      uint64(18446462598732906495),
				EmptySquares:            uint64(281474976645120),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameBoard := GameBoard{Board: tt.board}
			gameBoard.CalculateUsefulBitboards()

			assert.Equal(t, tt.expectedUsefulBitboards, gameBoard.UsefulBitboards)
		})
	}
}

func getStartingBoard() Board {
	return Board{
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
	}
}

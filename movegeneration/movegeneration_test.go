package movegeneration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
)

func TestCalculateWhiteKnightMoves(t *testing.T) {
	tests := []struct {
		name          string
		board         boardrepresentation.Board
		expectedMoves []PieceMove
		expectedError error
	}{
		{
			name: "One in middle of board (d5)",
			board: boardrepresentation.Board{
				WhiteKnights: uint64(34359738368),
			},
			expectedMoves: []PieceMove{
				getNormalKnightMove(uint64(34359738368), uint64(1125899906842624)), // d5-c7
				getNormalKnightMove(uint64(34359738368), uint64(2199023255552)),    // d5-b6
				getNormalKnightMove(uint64(34359738368), uint64(33554432)),         // d5-c7
				getNormalKnightMove(uint64(34359738368), uint64(262144)),           // d5-c3
				getNormalKnightMove(uint64(34359738368), uint64(1048576)),          // d5-e3
				getNormalKnightMove(uint64(34359738368), uint64(536870912)),        // d5-f4
				getNormalKnightMove(uint64(34359738368), uint64(35184372088832)),   // d5-f6
				getNormalKnightMove(uint64(34359738368), uint64(4503599627370496)), // d5-e7
			},
		},
		{
			name: "One on edge of board (a7)",
			board: boardrepresentation.Board{
				WhiteKnights: uint64(281474976710656),
			},
			expectedMoves: []PieceMove{
				getNormalKnightMove(uint64(281474976710656), uint64(288230376151711744)), // a7-c8
				getNormalKnightMove(uint64(281474976710656), uint64(4398046511104)),      // a7-c6
				getNormalKnightMove(uint64(281474976710656), uint64(8589934592)),         // a7-b5

			},
		},
		{
			name: "Two pieces with overlap (a3, c7)",
			board: boardrepresentation.Board{
				WhiteKnights: uint64(1125899906908160),
			},
			expectedMoves: []PieceMove{
				getNormalKnightMove(uint64(65536), uint64(8589934592)),                     // a3-b5
				getNormalKnightMove(uint64(65536), uint64(67108864)),                       // a3-c4
				getNormalKnightMove(uint64(65536), uint64(1024)),                           // a3-c2
				getNormalKnightMove(uint64(65536), uint64(2)),                              // a3-b1
				getNormalKnightMove(uint64(1125899906842624), uint64(72057594037927936)),   // c7-a8
				getNormalKnightMove(uint64(1125899906842624), uint64(1099511627776)),       // c7-a6
				getNormalKnightMove(uint64(1125899906842624), uint64(8589934592)),          // c7-b5
				getNormalKnightMove(uint64(1125899906842624), uint64(34359738368)),         // c7-d5
				getNormalKnightMove(uint64(1125899906842624), uint64(17592186044416)),      // c7-e6
				getNormalKnightMove(uint64(1125899906842624), uint64(1152921504606846976)), // c7-e8
			},
		},
		// TODO: One capturing Pieces
		// TODO: Multi capturing Pieces
		// TODO: One blocked by own piece
		// TODO: multipl blocked by own piece
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameBoard := &boardrepresentation.GameBoard{
				Board: tt.board,
			}
			gameBoard.CalculateUsefulBitboards()

			moves, err := calculateWhiteKnightMoves(gameBoard)

			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			assert.ElementsMatch(t, tt.expectedMoves, moves)
		})
	}
}

func getNormalKnightMove(from, to uint64) PieceMove {
	return PieceMove{
		PositionBitboard: from,
		MoveBitboard:     to,
		PieceType:        boardrepresentation.KnightPieceType,
		MoveType:         boardrepresentation.NormalMoveType,
	}
}

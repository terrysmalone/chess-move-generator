package movegeneration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terrysmalone/chess-move-generator/boardrepresentation"
)

func TestCalculateWhiteKnightMoves(t *testing.T) {
	tests := []struct {
		name          string
		board         boardrepresentation.Board
		expectedMoves []PieceMove
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
		{
			name: "One piece capturing and blocked (h4)",
			board: boardrepresentation.Board{
				WhiteKnights: uint64(2147483648),     // h4
				WhitePawns:   uint64(2117632),        // e2, f3(blocking), g2(blocking)
				BlackBishops: uint64(70368878395392), // d4, g6(capturable)

			},
			expectedMoves: []PieceMove{
				getCapturingKnightMove(uint64(2147483648), uint64(70368744177664)), // h4-g6 (capture)
				getNormalKnightMove(uint64(2147483648), uint64(137438953472)),      // a3-f5

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameBoard := &boardrepresentation.GameBoard{
				Board: tt.board,
			}
			gameBoard.CalculateUsefulBitboards()

			moves := &[]PieceMove{}
			calculateKnightMoves(moves,
				gameBoard.Board.WhiteKnights,
				gameBoard.UsefulBitboards.AllBlackOccupiedSquares,
				gameBoard.UsefulBitboards.EmptySquares)

			assert.ElementsMatch(t, tt.expectedMoves, *moves)
		})
	}
}

func TestCalculateBlackKnightMoves(t *testing.T) {
	tests := []struct {
		name          string
		board         boardrepresentation.Board
		expectedMoves []PieceMove
	}{
		{
			name: "One in middle of board (c4)",
			board: boardrepresentation.Board{
				BlackKnights: uint64(67108864),
			},
			expectedMoves: []PieceMove{
				getNormalKnightMove(uint64(67108864), uint64(65536)),         // c4-a3
				getNormalKnightMove(uint64(67108864), uint64(4294967296)),    // c4-a5
				getNormalKnightMove(uint64(67108864), uint64(512)),           // c4-b2
				getNormalKnightMove(uint64(67108864), uint64(2199023255552)), // c4-b6
				getNormalKnightMove(uint64(67108864), uint64(8796093022208)), // c4-d6
				getNormalKnightMove(uint64(67108864), uint64(2048)),          // c4-d2
				getNormalKnightMove(uint64(67108864), uint64(1048576)),       // c4-e3
				getNormalKnightMove(uint64(67108864), uint64(68719476736)),   // c4-e5
			},
		},
		{
			name: "One on edge of board (h1)",
			board: boardrepresentation.Board{
				BlackKnights: uint64(128),
			},
			expectedMoves: []PieceMove{
				getNormalKnightMove(uint64(128), uint64(4194304)), // h1-g3
				getNormalKnightMove(uint64(128), uint64(8192)),    // h1-f2
			},
		},
		{
			name: "Two pieces with overlap (a3, c7)",
			board: boardrepresentation.Board{
				BlackKnights: uint64(1125899906908160),
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameBoard := &boardrepresentation.GameBoard{
				Board: tt.board,
			}
			gameBoard.CalculateUsefulBitboards()

			moves := &[]PieceMove{}
			calculateKnightMoves(moves,
				gameBoard.Board.BlackKnights,
				gameBoard.UsefulBitboards.AllWhiteOccupiedSquares,
				gameBoard.UsefulBitboards.EmptySquares)

			assert.ElementsMatch(t, tt.expectedMoves, *moves)
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

func getCapturingKnightMove(from, to uint64) PieceMove {
	return PieceMove{
		PositionBitboard: from,
		MoveBitboard:     to,
		PieceType:        boardrepresentation.KnightPieceType,
		MoveType:         boardrepresentation.CaptureMoveType,
	}
}

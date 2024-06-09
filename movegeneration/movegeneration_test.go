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
			expectedMoves: getMoves(
				boardrepresentation.KnightPieceType,
				uint64(34359738368),
				[]uint64{
					uint64(1125899906842624),
					uint64(2199023255552),
					uint64(33554432),
					uint64(262144),
					uint64(1048576),
					uint64(536870912),
					uint64(35184372088832),
					uint64(4503599627370496),
				},
				[]uint64{}),
		},
		{
			name: "One on edge of board (a7)",
			board: boardrepresentation.Board{
				WhiteKnights: uint64(281474976710656),
			},
			expectedMoves: getMoves(
				boardrepresentation.KnightPieceType,
				uint64(281474976710656),
				[]uint64{
					uint64(288230376151711744),
					uint64(4398046511104),
					uint64(8589934592),
				},
				[]uint64{}),
		},
		{
			name: "Two pieces with overlap (a3, c7)",
			board: boardrepresentation.Board{
				WhiteKnights: uint64(1125899906908160),
			},
			expectedMoves: append(
				getMoves(
					boardrepresentation.KnightPieceType,
					uint64(65536),
					[]uint64{
						uint64(8589934592),
						uint64(67108864),
						uint64(1024),
						uint64(2),
					},
					[]uint64{}),
				getMoves(
					boardrepresentation.KnightPieceType,
					uint64(1125899906842624),
					[]uint64{
						uint64(72057594037927936),
						uint64(1099511627776),
						uint64(8589934592),
						uint64(34359738368),
						uint64(17592186044416),
						uint64(1152921504606846976),
					},
					[]uint64{})...),
		},
		{
			name: "One piece capturing and blocked (h4)",
			board: boardrepresentation.Board{
				WhiteKnights: uint64(2147483648),     // h4
				WhitePawns:   uint64(2117632),        // e2, f3(blocking), g2(blocking)
				BlackBishops: uint64(70368878395392), // d4, g6(capturable)

			},
			expectedMoves: getMoves(
				boardrepresentation.KnightPieceType,
				uint64(2147483648),
				[]uint64{
					uint64(137438953472),
				},
				[]uint64{
					uint64(70368744177664),
				}),
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
			expectedMoves: getMoves(
				boardrepresentation.KnightPieceType,
				uint64(67108864),
				[]uint64{
					uint64(65536),
					uint64(4294967296),
					uint64(512),
					uint64(2199023255552),
					uint64(8796093022208),
					uint64(2048),
					uint64(1048576),
					uint64(68719476736),
				},
				[]uint64{}),
		},
		{
			name: "One on edge of board (h1)",
			board: boardrepresentation.Board{
				BlackKnights: uint64(128),
			},
			expectedMoves: getMoves(
				boardrepresentation.KnightPieceType,
				uint64(128),
				[]uint64{
					uint64(4194304),
					uint64(8192),
				},
				[]uint64{}),
		},
		{
			name: "Two pieces with overlap (a3, c7)",
			board: boardrepresentation.Board{
				BlackKnights: uint64(1125899906908160),
			},
			expectedMoves: append(
				getMoves(
					boardrepresentation.KnightPieceType,
					uint64(65536),
					[]uint64{
						uint64(8589934592),
						uint64(67108864),
						uint64(1024),
						uint64(2),
					},
					[]uint64{}),
				getMoves(
					boardrepresentation.KnightPieceType,
					uint64(1125899906842624),
					[]uint64{
						uint64(72057594037927936),
						uint64(1099511627776),
						uint64(8589934592),
						uint64(34359738368),
						uint64(17592186044416),
						uint64(1152921504606846976),
					},
					[]uint64{})...),
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

func TestCalculateWhiteBishopMoves(t *testing.T) {
	tests := []struct {
		name          string
		board         boardrepresentation.Board
		expectedMoves []PieceMove
	}{
		{
			name: "One in middle of board e4)",
			board: boardrepresentation.Board{
				WhiteBishops: uint64(268435456),
			},
			expectedMoves: getMoves(
				boardrepresentation.BishopPieceType,
				uint64(268435456),
				[]uint64{
					34359738368,
					4398046511104,
					562949953421312,
					72057594037927936,
					2097152,
					16384,
					128,
					524288,
					1024,
					2,
					137438953472,
					70368744177664,
					36028797018963968},
				[]uint64{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameBoard := &boardrepresentation.GameBoard{
				Board: tt.board,
			}
			gameBoard.CalculateUsefulBitboards()

			moves := &[]PieceMove{}
			calculateBishopMoves(moves,
				gameBoard.Board.WhiteBishops,
				&gameBoard.UsefulBitboards,
				true)

			assert.ElementsMatch(t, tt.expectedMoves, *moves)
		})
	}
}

func getMoves(pieceType boardrepresentation.PieceType, from uint64, normalToMoves, attackToMoves []uint64) []PieceMove {
	pieceMoves := []PieceMove{}

	for _, to := range normalToMoves {
		pieceMoves = append(pieceMoves, PieceMove{
			PositionBitboard: from,
			MoveBitboard:     to,
			PieceType:        pieceType,
			MoveType:         boardrepresentation.NormalMoveType,
		})
	}

	for _, to := range attackToMoves {
		pieceMoves = append(pieceMoves, PieceMove{
			PositionBitboard: from,
			MoveBitboard:     to,
			PieceType:        pieceType,
			MoveType:         boardrepresentation.CaptureMoveType,
		})
	}

	return pieceMoves
}

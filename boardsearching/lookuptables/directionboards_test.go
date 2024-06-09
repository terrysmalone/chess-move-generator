package lookuptables

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateUpBoards(t *testing.T) {
	tests := []struct {
		name             string
		positionIndex    int
		expectedBitboard uint64
	}{
		{
			name:             "a1",
			positionIndex:    0,
			expectedBitboard: uint64(72340172838076672),
		},
		{
			name:             "c3",
			positionIndex:    18,
			expectedBitboard: uint64(289360691352043520),
		},
		{
			name:             "f8 (top row)",
			positionIndex:    61,
			expectedBitboard: uint64(0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := UpBoard[tt.positionIndex]

			assert.Equal(t, tt.expectedBitboard, board, fmt.Errorf("Expected %d got %d", tt.expectedBitboard, board))
		})
	}
}

func TestCalculateDownBoards(t *testing.T) {
	tests := []struct {
		name             string
		positionIndex    int
		expectedBitboard uint64
	}{
		{
			name:             "d1 (bottom row)",
			positionIndex:    3,
			expectedBitboard: uint64(0),
		},

		{
			name:             "h4",
			positionIndex:    31,
			expectedBitboard: uint64(8421504),
		},
		{
			name:             "a8",
			positionIndex:    56,
			expectedBitboard: uint64(282578800148737),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := DownBoard[tt.positionIndex]

			assert.Equal(t, tt.expectedBitboard, board, fmt.Errorf("Expected %d got %d", tt.expectedBitboard, board))
		})
	}
}

func TestCalculateLeftBoards(t *testing.T) {
	tests := []struct {
		name             string
		positionIndex    int
		expectedBitboard uint64
	}{
		{
			name:             "a4 (left column)",
			positionIndex:    24,
			expectedBitboard: uint64(0),
		},

		{
			name:             "h8",
			positionIndex:    63,
			expectedBitboard: uint64(9151314442816847872),
		},
		{
			name:             "c3",
			positionIndex:    18,
			expectedBitboard: uint64(196608),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := LeftBoard[tt.positionIndex]

			assert.Equal(t, tt.expectedBitboard, board, fmt.Errorf("Expected %d got %d", tt.expectedBitboard, board))
		})
	}
}

func TestCalculateRightBoards(t *testing.T) {
	tests := []struct {
		name             string
		positionIndex    int
		expectedBitboard uint64
	}{
		{
			name:             "h5 (right column)",
			positionIndex:    39,
			expectedBitboard: uint64(0),
		},
		{
			name:             "a8",
			positionIndex:    56,
			expectedBitboard: uint64(18302628885633695744),
		},
		{
			name:             "c3",
			positionIndex:    18,
			expectedBitboard: uint64(16252928),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := RightBoard[tt.positionIndex]

			assert.Equal(t, tt.expectedBitboard, board, fmt.Errorf("Expected %d got %d", tt.expectedBitboard, board))
		})
	}
}

func TestCalculateUpRightBoards(t *testing.T) {
	tests := []struct {
		name             string
		positionIndex    int
		expectedBitboard uint64
	}{
		{
			name:             "h5 (right column)",
			positionIndex:    39,
			expectedBitboard: uint64(0),
		},
		{
			name:             "a8",
			positionIndex:    56,
			expectedBitboard: uint64(0),
		},
		{
			name:             "a1",
			positionIndex:    0,
			expectedBitboard: uint64(9241421688590303744),
		},
		{
			name:             "b5",
			positionIndex:    33,
			expectedBitboard: uint64(1155177702467043328),
		},
		{
			name:             "g2",
			positionIndex:    14,
			expectedBitboard: uint64(8388608),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := UpRightBoard[tt.positionIndex]

			assert.Equal(t, tt.expectedBitboard, board, fmt.Errorf("Expected %d got %d", tt.expectedBitboard, board))
		})
	}
}

func TestCalculateDownRightBoards(t *testing.T) {
	tests := []struct {
		name             string
		positionIndex    int
		expectedBitboard uint64
	}{
		{
			name:             "h5 (right column)",
			positionIndex:    39,
			expectedBitboard: uint64(0),
		},
		{
			name:             "a8",
			positionIndex:    56,
			expectedBitboard: uint64(567382630219904),
		},
		{
			name:             "a1",
			positionIndex:    0,
			expectedBitboard: uint64(0),
		},
		{
			name:             "b5",
			positionIndex:    33,
			expectedBitboard: uint64(67637280),
		},
		{
			name:             "g2",
			positionIndex:    14,
			expectedBitboard: uint64(128),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := DownRightBoard[tt.positionIndex]

			assert.Equal(t, tt.expectedBitboard, board, fmt.Errorf("Expected %d got %d", tt.expectedBitboard, board))
		})
	}
}

func TestCalculateDownLeftBoards(t *testing.T) {
	tests := []struct {
		name             string
		positionIndex    int
		expectedBitboard uint64
	}{
		{
			name:             "h5 (right column)",
			positionIndex:    39,
			expectedBitboard: uint64(1075843080),
		},
		{
			name:             "a8",
			positionIndex:    56,
			expectedBitboard: uint64(0),
		},
		{
			name:             "a1",
			positionIndex:    0,
			expectedBitboard: uint64(0),
		},
		{
			name:             "b5",
			positionIndex:    33,
			expectedBitboard: uint64(16777216),
		},
		{
			name:             "g2",
			positionIndex:    14,
			expectedBitboard: uint64(32),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := DownLeftBoard[tt.positionIndex]

			assert.Equal(t, tt.expectedBitboard, board, fmt.Errorf("Expected %d got %d", tt.expectedBitboard, board))
		})
	}
}

package bitboardoperations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSquareIndexFromBitboard(t *testing.T) {
	tests := []struct {
		name          string
		bitboard      uint64
		expectedIndex byte
	}{
		{
			name:          "First (a1)",
			bitboard:      uint64(1),
			expectedIndex: 0,
		},
		{
			name:          "Last (h8)",
			bitboard:      uint64(9223372036854775808),
			expectedIndex: 63,
		},
		{
			name:          "e4",
			bitboard:      uint64(268435456),
			expectedIndex: 28,
		},
		{
			name:          "b6",
			bitboard:      uint64(2199023255552),
			expectedIndex: 41,
		},
		{
			name:          "h6",
			bitboard:      uint64(140737488355328),
			expectedIndex: 47,
		},
		{
			name:          "a8",
			bitboard:      uint64(72057594037927936),
			expectedIndex: 56,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := getSquareIndexFromBitboard(tt.bitboard)

			assert.Equal(t, tt.expectedIndex, index, fmt.Errorf("Expected %d got %d", tt.expectedIndex, index))
		})
	}
}

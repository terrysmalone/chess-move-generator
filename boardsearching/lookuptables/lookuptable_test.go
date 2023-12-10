package lookuptables

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// They're hardcoded values so just make sure I put them in right.
// Since they double each time use that to verify
func TestConsts(t *testing.T) {

	assert.EqualValues(t, 1, A1)
	assert.EqualValues(t, A1*uint64(2), B1)
	assert.EqualValues(t, B1*uint64(2), C1)
	assert.EqualValues(t, C1*uint64(2), D1)
	assert.EqualValues(t, D1*uint64(2), E1)
	assert.EqualValues(t, E1*uint64(2), F1)
	assert.EqualValues(t, F1*uint64(2), G1)
	assert.EqualValues(t, G1*uint64(2), H1)

	assert.EqualValues(t, H1*uint64(2), A2)
	assert.EqualValues(t, A2*uint64(2), B2)
	assert.EqualValues(t, B2*uint64(2), C2)
	assert.EqualValues(t, C2*uint64(2), D2)
	assert.EqualValues(t, D2*uint64(2), E2)
	assert.EqualValues(t, E2*uint64(2), F2)
	assert.EqualValues(t, F2*uint64(2), G2)
	assert.EqualValues(t, G2*uint64(2), H2)

	assert.EqualValues(t, H2*uint64(2), A3)
	assert.EqualValues(t, A3*uint64(2), B3)
	assert.EqualValues(t, B3*uint64(2), C3)
	assert.EqualValues(t, C3*uint64(2), D3)
	assert.EqualValues(t, D3*uint64(2), E3)
	assert.EqualValues(t, E3*uint64(2), F3)
	assert.EqualValues(t, F3*uint64(2), G3)
	assert.EqualValues(t, G3*uint64(2), H3)

	assert.EqualValues(t, H3*uint64(2), A4)
	assert.EqualValues(t, A4*uint64(2), B4)
	assert.EqualValues(t, B4*uint64(2), C4)
	assert.EqualValues(t, C4*uint64(2), D4)
	assert.EqualValues(t, D4*uint64(2), E4)
	assert.EqualValues(t, E4*uint64(2), F4)
	assert.EqualValues(t, F4*uint64(2), G4)
	assert.EqualValues(t, G4*uint64(2), H4)

	assert.EqualValues(t, H4*uint64(2), A5)
	assert.EqualValues(t, A5*uint64(2), B5)
	assert.EqualValues(t, B5*uint64(2), C5)
	assert.EqualValues(t, C5*uint64(2), D5)
	assert.EqualValues(t, D5*uint64(2), E5)
	assert.EqualValues(t, E5*uint64(2), F5)
	assert.EqualValues(t, F5*uint64(2), G5)
	assert.EqualValues(t, G5*uint64(2), H5)

	assert.EqualValues(t, H5*uint64(2), A6)
	assert.EqualValues(t, A6*uint64(2), B6)
	assert.EqualValues(t, B6*uint64(2), C6)
	assert.EqualValues(t, C6*uint64(2), D6)
	assert.EqualValues(t, D6*uint64(2), E6)
	assert.EqualValues(t, E6*uint64(2), F6)
	assert.EqualValues(t, F6*uint64(2), G6)
	assert.EqualValues(t, G6*uint64(2), H6)

	assert.EqualValues(t, H6*uint64(2), A7)
	assert.EqualValues(t, A7*uint64(2), B7)
	assert.EqualValues(t, B7*uint64(2), C7)
	assert.EqualValues(t, C7*uint64(2), D7)
	assert.EqualValues(t, D7*uint64(2), E7)
	assert.EqualValues(t, E7*uint64(2), F7)
	assert.EqualValues(t, F7*uint64(2), G7)
	assert.EqualValues(t, G7*uint64(2), H7)

	assert.EqualValues(t, H7*uint64(2), A8)
	assert.EqualValues(t, A8*uint64(2), B8)
	assert.EqualValues(t, B8*uint64(2), C8)
	assert.EqualValues(t, C8*uint64(2), D8)
	assert.EqualValues(t, D8*uint64(2), E8)
	assert.EqualValues(t, E8*uint64(2), F8)
	assert.EqualValues(t, F8*uint64(2), G8)
	assert.EqualValues(t, G8*uint64(2), H8)
}

func TestBitboardValueFromPosition(t *testing.T) {
	tests := []struct {
		name             string
		column           int
		row              int
		expectedBitboard uint64
	}{
		{
			name:             "First (a1)",
			column:           0,
			row:              0,
			expectedBitboard: uint64(1),
		},
		{
			name:             "Last (h8)",
			column:           7,
			row:              7,
			expectedBitboard: uint64(9223372036854775808),
		},
		{
			name:             "d3",
			column:           3,
			row:              2,
			expectedBitboard: uint64(524288),
		},
		{
			name:             "a5",
			column:           0,
			row:              4,
			expectedBitboard: uint64(4294967296),
		},
		{
			name:             "b8",
			column:           1,
			row:              7,
			expectedBitboard: uint64(144115188075855872),
		},
		{
			name:             "h2",
			column:           7,
			row:              1,
			expectedBitboard: uint64(32768),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bitboard := BitboardValueFromPosition[tt.column][tt.row]

			assert.Equal(t, tt.expectedBitboard, bitboard, fmt.Errorf("Expected %d got %d", tt.expectedBitboard, bitboard))
		})
	}
}

func TestBitboardValueFromIndex(t *testing.T) {
	tests := []struct {
		name             string
		index            int
		row              int
		expectedBitboard uint64
	}{
		{
			name:             "First (a1)",
			index:            0,
			expectedBitboard: uint64(1),
		},
		{
			name:             "Last (h8)",
			index:            63,
			expectedBitboard: uint64(9223372036854775808),
		},
		{
			name:             "f2",
			index:            13,
			expectedBitboard: uint64(8192),
		},
		{
			name:             "d3",
			index:            19,
			expectedBitboard: uint64(524288),
		},
		{
			name:             "a5",
			index:            32,
			expectedBitboard: uint64(4294967296),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bitboard := BitboardValueFromIndex[tt.index]

			assert.Equal(t, tt.expectedBitboard, bitboard, fmt.Errorf("Expected %d got %d", tt.expectedBitboard, bitboard))
		})
	}
}

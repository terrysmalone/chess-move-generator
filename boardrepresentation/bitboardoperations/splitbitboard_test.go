package bitboardoperations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitBitboard(t *testing.T) {
	tests := []struct {
		name                   string
		bitboard               uint64
		expectedSplitBitboards []uint64
	}{
		{
			name:                   "Only 1 (a1)",
			bitboard:               uint64(1),
			expectedSplitBitboards: []uint64{uint64(1)},
		},
		{
			name:                   "Only 1 (h8)",
			bitboard:               uint64(9223372036854775808),
			expectedSplitBitboards: []uint64{uint64(9223372036854775808)},
		},
		{
			name:                   "Only 1 (c4)",
			bitboard:               uint64(67108864),
			expectedSplitBitboards: []uint64{uint64(67108864)},
		},
		{
			name:                   "Split 2 - a1 and h8",
			bitboard:               uint64(9223372036854775809),
			expectedSplitBitboards: []uint64{uint64(1), uint64(9223372036854775808)},
		},
		{
			name:                   "Split 2 - d4 and e4",
			bitboard:               uint64(402653184),
			expectedSplitBitboards: []uint64{uint64(134217728), uint64(268435456)},
		},
		{
			name:                   "Split 3 - f8, g8 and h8",
			bitboard:               uint64(16140901064495857664),
			expectedSplitBitboards: []uint64{uint64(2305843009213693952), uint64(4611686018427387904), uint64(9223372036854775808)},
		},
		{
			name:                   "Split 3 - b2, e5 and b8",
			bitboard:               uint64(144115256795333120),
			expectedSplitBitboards: []uint64{uint64(512), uint64(68719476736), uint64(144115188075855872)},
		},
		{
			name:     "Split a lot - b2, d2, h2, c3, g3, b6, e6, d7",
			bitboard: uint64(2271591027476992),
			expectedSplitBitboards: []uint64{
				uint64(512),
				uint64(2048),
				uint64(32768),
				uint64(262144),
				uint64(4194304),
				uint64(2199023255552),
				uint64(17592186044416),
				uint64(2251799813685248)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitBitboards := SplitBitboard(tt.bitboard)

			assert.Equal(t, tt.expectedSplitBitboards, splitBitboards, fmt.Errorf("Expected %d got %d", tt.expectedSplitBitboards, splitBitboards))
		})
	}
}

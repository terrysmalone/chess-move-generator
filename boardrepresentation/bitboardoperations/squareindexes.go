package bitboardoperations

var index64 [64]byte = [64]byte{
	63, 0, 58, 1, 59, 47, 53, 2,
	60, 39, 48, 27, 54, 33, 42, 3,
	61, 51, 37, 40, 49, 18, 28, 20,
	55, 30, 34, 11, 43, 14, 22, 4,
	62, 57, 46, 52, 38, 26, 32, 41,
	50, 36, 17, 19, 29, 10, 13, 21,
	56, 45, 25, 31, 35, 16, 9, 12,
	44, 24, 15, 8, 23, 7, 6, 5,
}

const debruijn64 = uint64(0x07EDD5E59A4E28C2)

// getSquareIndexFromBitboard uses a de Bruijn sequence to quickly
// determine the index (0-63) from a bitboard value
func getSquareIndexFromBitboard(bitboard uint64) byte {
	return index64[((bitboard&(^bitboard+1))*debruijn64)>>58]
}

// GetSquareIndexesFromBitboard gets multiple indexes from
// a bitboard with more than 1 square set
func GetSquareIndexesFromBitboard(bitboard uint64) []byte {
	values := []byte{}

	for bitboard != 0 {
		first := getSquareIndexFromBitboard(bitboard)
		values = append(values, first)

		bitboard &= ^(uint64(1) << first)
	}

	return values
}

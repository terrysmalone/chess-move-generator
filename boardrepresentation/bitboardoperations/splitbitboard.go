package bitboardoperations

func SplitBitboard(bitboardToSplit uint64) []uint64 {
	size := getPopCount(bitboardToSplit)

	// I really want this to be an array for speed but we can't set
	// an array size from a variable :/
	bitboards := make([]uint64, size)

	var reducingBitboard = bitboardToSplit

	reducingBitboard &= reducingBitboard - 1

	for i := 0; i < size; i++ {
		bitboards[i] = bitboardToSplit & ^reducingBitboard

		bitboardToSplit &= bitboardToSplit - 1

		reducingBitboard &= reducingBitboard - 1
	}

	return bitboards
}

func getPopCount(bitboardValue uint64) int {
	count := 0

	for bitboardValue != 0 {
		count++

		bitboardValue &= (bitboardValue - 1)
	}

	return count
}

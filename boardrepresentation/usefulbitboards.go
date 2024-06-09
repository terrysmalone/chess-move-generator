package boardrepresentation

type UsefulBitboards struct {
	AllWhiteOccupiedSquares uint64
	AllBlackOccupiedSquares uint64
	AllOccupiedSquares      uint64
	EmptySquares            uint64

	WhiteOrEmpty uint64
	BlackOrEmpty uint64
}

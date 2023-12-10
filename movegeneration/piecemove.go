package movegeneration

import "github.com/terrysmalone/chess-move-generator/boardrepresentation"

type PieceMove struct {
	PositionBitboard uint64
	MovesBitboard    uint64
	PieceType        boardrepresentation.PieceType
	MoveType         boardrepresentation.MoveType
}

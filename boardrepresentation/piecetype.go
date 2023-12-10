package boardrepresentation

type PieceType int

const (
	PawnPieceType   PieceType = iota
	KnightPieceType PieceType = iota
	BishopPieceType PieceType = iota
	RookPieceType   PieceType = iota
	QueenPieceType  PieceType = iota
	KingPieceType   PieceType = iota
)

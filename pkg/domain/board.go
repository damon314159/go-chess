package domain

type File uint8
type Rank uint8

type Square interface {
	File() File
	Rank() Rank
}

type Board interface {
	ValidateMove(Move) bool
	MovePiece(Move)
	RemovePiece(Square)
	Checks(Colour)
}

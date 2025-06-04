package domain

type Value uint8
type Colour string

const (
	White Colour = "White"
	Black Colour = "Black"
)

type Piece interface {
	Moves(Square) []Move
	Value() Value
	Colour() Colour
}

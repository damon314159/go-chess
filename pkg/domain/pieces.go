package domain

// The numeric relative value of the piece (https://en.wikipedia.org/wiki/Chess_piece_relative_value)
type Value uint8

// The colour of the piece, either white or black
type Colour string

const (
	White Colour = "White"
	Black Colour = "Black"
)

// Represents any piece on the board
type Piece interface {
	// Finds all moves that the Piece could make given only its position, without validating them
	Moves(Square) []Move
	// Getter for this Piece's relative value
	Value() Value
	// Getter for this Piece's colour
	Colour() Colour
}

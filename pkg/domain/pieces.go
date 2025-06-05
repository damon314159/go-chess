package domain

// The numeric relative value of the piece (https://en.wikipedia.org/wiki/Chess_piece_relative_value)
type Value uint8

// The colour of the piece, either white or black
// You should not use this type directly, but instead use the White or Black constants
type Colour struct {
	colour int8
}

// Implements fmt.Stringer to provide a nice display of the colour
func (c Colour) String() string {
	if c.colour == 1 {
		return "white"
	}
	return "black"
}

// One of the two possible colours of a piece
var (
	White Colour = Colour{1}
	Black Colour = Colour{0}
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

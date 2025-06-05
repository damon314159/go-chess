package domain

// The vertical files on the board, labelled A-H
type File int8

// The horizontal ranks on the board, labelled 1-8
type Rank int8

// Represents a single one of the 64 Squares on the board
type Square interface {
	// Getter for this Square's File
	File() File
	// Getter for this Square's Rank
	Rank() Rank

	// Constructor for a Square translated arbitrarily from this one
	Translate(int8, int8) (Square, error)
	// Constructor for a Square translated up from this one
	Up(int8) (Square, error)
	// Constructor for a Square translated down from this one
	Down(int8) (Square, error)
	// Constructor for a Square translated left from this one
	Left(int8) (Square, error)
	// Constructor for a Square translated right from this one
	Right(int8) (Square, error)

	// Implements fmt.Stringer to provide a nice display of it's coordinates
	String() string
}

// Represents the board on which the game is played
type Board interface {
	// Take in a move and determine whether it is valid, based on the rest of the Board state
	ValidateMove(Move) bool
	// Perform a move on the Board that has been pre-validated
	MovePiece(Move) error
	// Take a piece off the Board
	RemovePiece(Square) error
	// Determine whether any checks (threats) are made against the king of the Colour
	// TODO: determine a return type for this
	Checks(Colour)
}

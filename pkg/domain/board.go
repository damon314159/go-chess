package domain

// The vertical files on the board, labelled A-H
type File uint8

// The horizontal ranks on the board, labelled 1-8
type Rank uint8

// Represents a single one of the 64 Squares on the board
type Square interface {
	// Getter for this Square's File
	File() File
	// Getter for this Square's Rank
	Rank() Rank

	// Constructor for a Square translated arbitrarily from this one
	Translate(uint8, uint8) (Square, error)
	// Constructor for a Square translated up from this one
	Up(uint8) (Square, error)
	// Constructor for a Square translated down from this one
	Down(uint8) (Square, error)
	// Constructor for a Square translated left from this one
	Left(uint8) (Square, error)
	// Constructor for a Square translated right from this one
	Right(uint8) (Square, error)
}

// Represents the board on which the game is played
type Board interface {
	// Take in a move and determine whether it is valid, based on the rest of the Board state
	ValidateMove(Move) bool
	// Perform a move on the Board that has been pre-validated
	MovePiece(Move)
	// Take a piece off the Board
	RemovePiece(Square)
	// Determine whether any checks (threats) are made against the king of the Colour
	Checks(Colour)
}

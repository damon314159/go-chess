package domain

// Represents a move of one piece across the board
type Move interface {
	// Where the move is heading to
	To() Square
	// Where the move is coming from
	From() Square

	// There must be nothing in these squares for the move to be valid
	NothingBlocking() []Square
	// Whether a capture must be made for this move to be valid
	CaptureRequired() bool
}

package domain

// Represents a move of one piece across the board
type Move interface {
	// Where the move is heading to
	To() Square
	// Where the move is coming from
	From() Square

	// There must be nothing in these squares for the move to be valid
	NothingBlocking() []Square
	// There must be something in these squares for the move to be valid
	SomethingBlocking() []Square
}

package pieces

import (
	"github.com/damon314159/go-chess/pkg/board"
	"github.com/damon314159/go-chess/pkg/domain"
)

type Bishop struct {
	colour domain.Colour
}

var _ domain.Piece = (*Bishop)(nil)

func NewBishop(colour domain.Colour) Bishop {
	return Bishop{colour}
}

const BishopValue = 3

type diagonalDirection int8

const (
	upRight diagonalDirection = iota
	upLeft
	downRight
	downLeft
)

func (b Bishop) Moves(from domain.Square) []domain.Move {
	moves := make([]domain.Move, 0, board.MaxDiagDistance*2)

	moves = exploreDiagonal(moves, from, upRight)
	moves = exploreDiagonal(moves, from, upLeft)
	moves = exploreDiagonal(moves, from, downRight)
	moves = exploreDiagonal(moves, from, downLeft)

	return moves
}

func exploreDiagonal(
	moves []domain.Move,
	from domain.Square,
	direction diagonalDirection,
) []domain.Move {
	getTranslation := func(distance int8) (int8, int8) {
		switch direction {
		case upRight:
			return distance, distance
		case upLeft:
			return -distance, distance
		case downRight:
			return distance, -distance
		case downLeft:
			return -distance, -distance
		default:
			// Default case should never happen as this type isn't exported
			// and the only values are defined above, which are exhaustively matched
			return 0, 0
		}
	}

	diag := make([]domain.Square, board.MaxDiagDistance)

	for distanceZeroIdx := range board.MaxDiagDistance {
		distance := distanceZeroIdx + 1

		nextSquare, err := from.Translate(getTranslation(distance))
		if err != nil {
			// Left the board, so stop checking this diagonal
			break
		}

		// Create a copy of the explored diagonal to avoid mutation issues
		exploredDiag := make([]domain.Square, len(diag))
		copy(exploredDiag, diag)

		moves = append(
			moves,
			board.NewMove(
				board.MoveConfig{From: from, To: nextSquare, NothingBlocking: exploredDiag},
			),
		)

		// Add the current square to the explored diagonal before the loop continues
		diag = append(diag, nextSquare)
	}

	return moves
}

func (Bishop) Value() domain.Value {
	return BishopValue
}

func (b Bishop) Colour() domain.Colour {
	return b.colour
}

package pieces

import (
	"github.com/damon314159/go-chess/pkg/board"
	"github.com/damon314159/go-chess/pkg/domain"
)

type Knight struct {
	colour domain.Colour
}

var _ domain.Piece = (*Knight)(nil)

func NewKnight(colour domain.Colour) Knight {
	return Knight{colour}
}

const KnightValue = 3

func (k Knight) Moves(from domain.Square) []domain.Move {
	moves := make([]domain.Move, 0, 8)

	translations := []struct {
		right int8
		up    int8
	}{
		{2, 1},
		{2, -1},
		{-2, 1},
		{-2, -1},
		{1, 2},
		{1, -2},
		{-1, 2},
		{-1, -2},
	}

	for _, translation := range translations {
		to, err := from.Translate(translation.right, translation.up)
		if err != nil {
			continue
		}

		moves = append(moves, board.NewMove(board.MoveConfig{From: from, To: to}))
	}

	return moves
}

func (Knight) Value() domain.Value {
	return KnightValue
}

func (k Knight) Colour() domain.Colour {
	return k.colour
}

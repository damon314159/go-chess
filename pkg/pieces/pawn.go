package pieces

import (
	"github.com/damon314159/go-chess/pkg/board"
	"github.com/damon314159/go-chess/pkg/domain"
)

type Pawn struct {
	colour domain.Colour
}

var _ domain.Piece = (*Pawn)(nil)

const (
	whiteHomeRank = 2
	blackHomeRank = 7
)

func (p Pawn) Moves(from domain.Square) []domain.Move {
	var advance func(spaces uint8) (domain.Square, error)
	var isHome bool

	switch p.colour {
	case domain.White:
		advance = from.Up
		isHome = from.Rank() == whiteHomeRank
	case domain.Black:
		advance = from.Down
		isHome = from.Rank() == blackHomeRank
	}

	moves := make([]domain.Move, 0, 4)

	if advanced1, err := advance(1); err == nil {
		moves = append(moves, board.NewMove(from, advanced1, []domain.Square{advanced1}, false))

		if isHome {
			if advanced2, err := advance(2); err == nil {
				moves = append(
					moves,
					board.NewMove(from, advanced2, []domain.Square{advanced1, advanced2}, false),
				)
			}
		}

		if diagLeft, err := advanced1.Left(1); err == nil {
			moves = append(moves, board.NewMove(from, diagLeft, nil, true))
		}

		if diagRight, err := advanced1.Right(1); err == nil {
			moves = append(moves, board.NewMove(from, diagRight, nil, true))
		}
	}

	return moves

}

func (Pawn) Value() domain.Value {
	return 1
}

func (p Pawn) Colour() domain.Colour {
	return p.colour
}

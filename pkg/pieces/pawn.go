package pieces

import (
	"github.com/damon314159/go-chess/pkg/board"
	"github.com/damon314159/go-chess/pkg/domain"
)

type Pawn struct {
	colour domain.Colour
}

var _ domain.Piece = (*Pawn)(nil)

func NewPawn(colour domain.Colour) Pawn {
	return Pawn{colour}
}

const PawnValue = 1

func (p Pawn) Moves(from domain.Square) []domain.Move {
	var advance func(spaces uint8) (domain.Square, error)
	var isHome bool

	switch p.colour {
	case domain.White:
		advance = from.Up
		isHome = from.Rank() == board.WhiteHomeRank
	case domain.Black:
		advance = from.Down
		isHome = from.Rank() == board.BlackHomeRank
	}

	moves := make([]domain.Move, 0, 4)

	if advanced1, err := advance(1); err == nil {
		moves = append(
			moves,
			board.NewMove(
				board.MoveConfig{
					From:            from,
					To:              advanced1,
					NothingBlocking: []domain.Square{advanced1},
				},
			),
		)

		if isHome {
			if advanced2, err := advance(2); err == nil {
				moves = append(
					moves,
					board.NewMove(
						board.MoveConfig{
							From:            from,
							To:              advanced2,
							NothingBlocking: []domain.Square{advanced1, advanced2},
						},
					),
				)
			}
		}

		if diagLeft, err := advanced1.Left(1); err == nil {
			moves = append(
				moves,
				board.NewMove(board.MoveConfig{From: from, To: diagLeft, CaptureRequired: true}),
			)
		}

		if diagRight, err := advanced1.Right(1); err == nil {
			moves = append(
				moves,
				board.NewMove(board.MoveConfig{From: from, To: diagRight, CaptureRequired: true}),
			)
		}
	}

	return moves

}

func (Pawn) Value() domain.Value {
	return PawnValue
}

func (p Pawn) Colour() domain.Colour {
	return p.colour
}

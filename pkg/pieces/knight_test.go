package pieces_test

import (
	"testing"

	"github.com/damon314159/go-chess/pkg/board"
	"github.com/damon314159/go-chess/pkg/domain"
	"github.com/damon314159/go-chess/pkg/pieces"
)

func TestKnight_Value(t *testing.T) {
	t.Run("returns the correct value", func(t *testing.T) {
		knight := pieces.NewKnight(domain.White)

		value := knight.Value()

		if value != pieces.KnightValue {
			t.Errorf("expected value to be %d, got %d", pieces.KnightValue, value)
		}
	})
}

func TestKnight_Colour(t *testing.T) {
	t.Run("returns the correct colour when white", func(t *testing.T) {
		knight := pieces.NewKnight(domain.White)

		assertWhite(t, knight)
	})

	t.Run("returns the correct colour when black", func(t *testing.T) {
		knight := pieces.NewKnight(domain.Black)

		assertBlack(t, knight)
	})
}

func TestKnight_Moves(t *testing.T) {
	t.Run("moves should start on the correct square", func(t *testing.T) {
		knight := pieces.NewKnight(domain.White)

		moves := knight.Moves(board.E5)

		for _, move := range moves {
			if move.From() != board.E5 {
				t.Errorf("expected move to start on %s, got %s", board.E5, move.From())
			}
		}
	})

	t.Run("cannot move to current square", func(t *testing.T) {
		knight := pieces.NewKnight(domain.Black)

		moves := knight.Moves(board.E5)

		assertCannotMoveTo(t, moves, board.E5)
	})

	t.Run("should move to all 8 L shapes when in the centre", func(t *testing.T) {
		knight := pieces.NewKnight(domain.White)

		moves := knight.Moves(board.E5)

		assertMoveCount(t, moves, 8)
		assertCanMoveTo(t, moves, board.G6)
		assertCanMoveTo(t, moves, board.G4)
		assertCanMoveTo(t, moves, board.C6)
		assertCanMoveTo(t, moves, board.C4)
		assertCanMoveTo(t, moves, board.F3)
		assertCanMoveTo(t, moves, board.D3)
		assertCanMoveTo(t, moves, board.F7)
		assertCanMoveTo(t, moves, board.D7)
	})

	t.Run("should only move to 2 L shapes when in the corner", func(t *testing.T) {
		knight := pieces.NewKnight(domain.White)

		moves := knight.Moves(board.A1)

		assertMoveCount(t, moves, 2)
		assertCanMoveTo(t, moves, board.C2)
		assertCanMoveTo(t, moves, board.B3)
	})
}

package pieces_test

import (
	"testing"

	"github.com/damon314159/go-chess/pkg/board"
	"github.com/damon314159/go-chess/pkg/domain"
	"github.com/damon314159/go-chess/pkg/pieces"
)

func TestBishop_Value(t *testing.T) {
	t.Run("returns the correct value", func(t *testing.T) {
		bishop := pieces.NewBishop(domain.White)

		value := bishop.Value()

		if value != pieces.BishopValue {
			t.Errorf("expected value to be %d, got %d", pieces.BishopValue, value)
		}
	})
}

func TestBishop_Colour(t *testing.T) {
	t.Run("returns the correct colour when white", func(t *testing.T) {
		bishop := pieces.NewBishop(domain.White)

		assertWhite(t, bishop)
	})

	t.Run("returns the correct colour when black", func(t *testing.T) {
		bishop := pieces.NewBishop(domain.Black)

		assertBlack(t, bishop)
	})
}

func TestBishop_Moves(t *testing.T) {
	t.Run("moves should start on the correct square", func(t *testing.T) {
		bishop := pieces.NewBishop(domain.White)

		moves := bishop.Moves(board.E5)

		for _, move := range moves {
			if move.From() != board.E5 {
				t.Errorf("expected move to start on %s, got %s", board.E5, move.From())
			}
		}
	})

	t.Run("cannot move to current square", func(t *testing.T) {
		bishop := pieces.NewBishop(domain.Black)

		moves := bishop.Moves(board.E5)

		assertCannotMoveTo(t, moves, board.E5)
	})

	// TODO: Add tests for possible moves, including checking which squares must have nothing blocking
}

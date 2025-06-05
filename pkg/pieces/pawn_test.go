package pieces_test

import (
	"testing"

	"github.com/damon314159/go-chess/pkg/board"
	"github.com/damon314159/go-chess/pkg/domain"
	"github.com/damon314159/go-chess/pkg/pieces"
)

func TestPawn_Value(t *testing.T) {
	t.Run("returns the correct value", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.White)

		value := pawn.Value()

		if value != pieces.PawnValue {
			t.Errorf("expected value to be %d, got %d", pieces.PawnValue, value)
		}
	})
}

func TestPawn_Colour(t *testing.T) {
	t.Run("returns the correct colour when white", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.White)

		assertWhite(t, pawn)
	})

	t.Run("returns the correct colour when black", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.Black)

		assertBlack(t, pawn)
	})
}

func TestPawn_Moves(t *testing.T) {
	t.Run("moves should start on the correct square", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.White)

		moves := pawn.Moves(board.A3)

		for _, move := range moves {
			if move.From() != board.A3 {
				t.Errorf("expected move to start on %s, got %s", board.A3, move.From())
			}
		}
	})

	t.Run("cannot move to current square or backwards", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.Black)
		square, _ := board.NewSquare(domain.File(4), board.BlackHomeRank-1)

		moves := pawn.Moves(square)

		backwards, _ := square.Up(1)
		assertCannotMoveTo(t, moves, square)
		assertCannotMoveTo(t, moves, backwards)
	})

	t.Run("moves forwards one square when white and not on home rank", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.White)
		square, _ := board.NewSquare(domain.File(4), board.WhiteHomeRank+1)

		moves := pawn.Moves(square)

		up1, _ := square.Up(1)
		up2, _ := square.Up(2)
		assertCanMoveTo(t, moves, up1)
		assertCannotMoveTo(t, moves, up2)
	})

	t.Run("moves forwards one or two squares when white on home rank", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.White)
		square, _ := board.NewSquare(domain.File(4), board.WhiteHomeRank)

		moves := pawn.Moves(square)

		up1, _ := square.Up(1)
		up2, _ := square.Up(2)
		assertCanMoveTo(t, moves, up1)
		assertCanMoveTo(t, moves, up2)
	})

	t.Run("moves forwards one square when black and not on home rank", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.Black)
		square, _ := board.NewSquare(domain.File(4), board.BlackHomeRank-1)

		moves := pawn.Moves(square)

		down1, _ := square.Down(1)
		down2, _ := square.Down(2)
		assertCanMoveTo(t, moves, down1)
		assertCannotMoveTo(t, moves, down2)
	})

	t.Run("moves forwards one or two squares when black on home rank", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.Black)
		square, _ := board.NewSquare(domain.File(4), board.BlackHomeRank)

		moves := pawn.Moves(square)

		down1, _ := square.Down(1)
		down2, _ := square.Down(2)
		assertCanMoveTo(t, moves, down1)
		assertCanMoveTo(t, moves, down2)
	})

	t.Run("cannot move when on back rank", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.White)
		square, _ := board.NewSquare(domain.File(4), board.MaxRank)

		moves := pawn.Moves(square)

		assertMoveCount(t, moves, 0)
	})

	t.Run("can move diagonally but must capture", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.White)
		square, _ := board.NewSquare(domain.File(4), board.WhiteHomeRank+1)

		moves := pawn.Moves(square)

		diagLeft, _ := square.Translate(-1, 1)
		diagRight, _ := square.Translate(1, 1)
		assertCanMoveTo(t, moves, diagLeft)
		assertMustCapture(t, moves, diagLeft)
		assertCanMoveTo(t, moves, diagRight)
		assertMustCapture(t, moves, diagRight)
	})

	t.Run("cannot move diagonally off the board", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.White)
		square, _ := board.NewSquare(domain.File(1), board.WhiteHomeRank+1)

		moves := pawn.Moves(square)

		up1, _ := square.Up(1)
		diagRight, _ := up1.Right(1)
		assertMoveCount(t, moves, 2) // Only forwards and the other diagonal
		assertCanMoveTo(t, moves, diagRight)
		assertCanMoveTo(t, moves, up1)
	})
}

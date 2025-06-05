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
		colour := pawn.Colour()

		if colour != domain.White {
			t.Errorf("expected colour to be %s, got %s", domain.White, colour)
		}
	})

	t.Run("returns the correct colour when black", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.Black)
		colour := pawn.Colour()

		if colour != domain.Black {
			t.Errorf("expected colour to be %s, got %s", domain.Black, colour)
		}
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

	t.Run("returns forwards one square when white and not on home rank", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.White)
		square, _ := board.NewSquare(domain.File(4), board.WhiteHomeRank+1)

		moves := pawn.Moves(square)

		up1, _ := square.Up(1)
		up2, _ := square.Up(2)
		assertCanMoveTo(t, moves, up1)
		assertCannotMoveTo(t, moves, up2)
	})

	t.Run("returns forwards one or two squares when white on home rank", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.White)
		square, _ := board.NewSquare(domain.File(4), board.WhiteHomeRank)

		moves := pawn.Moves(square)

		up1, _ := square.Up(1)
		up2, _ := square.Up(2)
		assertCanMoveTo(t, moves, up1)
		assertCanMoveTo(t, moves, up2)
	})

	t.Run("returns forwards one square when black and not on home rank", func(t *testing.T) {
		pawn := pieces.NewPawn(domain.Black)
		square, _ := board.NewSquare(domain.File(4), board.BlackHomeRank-1)

		moves := pawn.Moves(square)

		down1, _ := square.Down(1)
		down2, _ := square.Down(2)
		assertCanMoveTo(t, moves, down1)
		assertCannotMoveTo(t, moves, down2)
	})

	t.Run("returns forwards one or two squares when black on home rank", func(t *testing.T) {
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

		if len(moves) != 0 {
			t.Errorf("expected no moves, got %d", len(moves))
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
}

func assertCanMoveTo(t *testing.T, moves []domain.Move, to domain.Square) {
	t.Helper()

	for _, move := range moves {
		if move.To() == to {
			return
		}
	}

	allToSquares := []domain.Square{}
	for _, move := range moves {
		allToSquares = append(allToSquares, move.To())
	}

	t.Errorf("expected move to %s, got %s", to, allToSquares)
}

func assertCannotMoveTo(t *testing.T, moves []domain.Move, to domain.Square) {
	t.Helper()

	for _, move := range moves {
		if move.To() == to {
			t.Errorf("expected move to not be able to move to %s", to)
		}
	}
}

func assertMustCapture(t *testing.T, moves []domain.Move, to domain.Square) {
	t.Helper()

	for _, move := range moves {
		if move.To() == to && move.CaptureRequired() {
			return
		}
	}

	allMustCaptureSquares := []domain.Square{}
	for _, move := range moves {
		if move.CaptureRequired() {
			allMustCaptureSquares = append(allMustCaptureSquares, move.To())
		}
	}

	t.Errorf(
		"expected must capture move to %s, got must capture moves to %s",
		to,
		allMustCaptureSquares,
	)
}

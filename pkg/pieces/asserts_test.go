package pieces_test

import (
	"testing"

	"github.com/damon314159/go-chess/pkg/domain"
)

func assertWhite(t *testing.T, piece domain.Piece) {
	t.Helper()

	if piece.Colour() != domain.White {
		t.Errorf("expected piece to be white, got %s", piece.Colour())
	}
}

func assertBlack(t *testing.T, piece domain.Piece) {
	t.Helper()

	if piece.Colour() != domain.Black {
		t.Errorf("expected piece to be black, got %s", piece.Colour())
	}
}

func assertMoveCount(t *testing.T, moves []domain.Move, count int) {
	t.Helper()

	if len(moves) != count {
		t.Errorf("expected %d moves, got %d", count, len(moves))
	}
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

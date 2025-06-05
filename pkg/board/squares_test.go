package board_test

import (
	"testing"

	"github.com/damon314159/go-chess/pkg/board"
	"github.com/damon314159/go-chess/pkg/domain"
)

func TestNewSquare(t *testing.T) {
	t.Run("creates a square", func(t *testing.T) {
		s, err := board.NewSquare(domain.File(1), domain.Rank(1))

		assertNoErr(t, err)
		assertFileAndRank(t, s, domain.File(1), domain.Rank(1))
	})

	t.Run("returns an error if the file is too small", func(t *testing.T) {
		_, err := board.NewSquare(domain.File(0), domain.Rank(1))

		assertErr(t, err)
	})

	t.Run("returns an error if the file is too large", func(t *testing.T) {
		_, err := board.NewSquare(domain.File(9), domain.Rank(1))

		assertErr(t, err)
	})

	t.Run("returns an error if the rank is too small", func(t *testing.T) {
		_, err := board.NewSquare(domain.File(1), domain.Rank(0))

		assertErr(t, err)
	})

	t.Run("returns an error if the rank is too large", func(t *testing.T) {
		_, err := board.NewSquare(domain.File(1), domain.Rank(9))

		assertErr(t, err)
	})
}

func assertNoErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func assertErr(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func assertFileAndRank(t *testing.T, square board.Square, file domain.File, rank domain.Rank) {
	t.Helper()

	if square.File() != file {
		t.Errorf("expected file %d, got %d", file, square.File())
	}

	if square.Rank() != rank {
		t.Errorf("expected rank %d, got %d", rank, square.Rank())
	}
}

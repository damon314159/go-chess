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

func TestSquare_Translate(t *testing.T) {
	t.Run("returns a new square with new file and rank", func(t *testing.T) {
		from, _ := board.NewSquare(domain.File(2), domain.Rank(2))
		to, err := from.Translate(1, 1)

		assertNoErr(t, err)
		assertFileAndRank(t, to, domain.File(3), domain.Rank(3))
	})

	t.Run("returns an error if the file is out of bounds", func(t *testing.T) {
		from, _ := board.NewSquare(domain.File(2), domain.Rank(2))
		_, err := from.Translate(10, 1)

		assertErr(t, err)
	})

	t.Run("returns an error if the rank is out of bounds", func(t *testing.T) {
		from, _ := board.NewSquare(domain.File(2), domain.Rank(2))
		_, err := from.Translate(1, 10)

		assertErr(t, err)
	})
}

func TestSquare_String(t *testing.T) {
	t.Run("returns a string representation of the square", func(t *testing.T) {
		s, _ := board.NewSquare(domain.File(3), domain.Rank(6))

		str := s.String()

		if str != "c6" {
			t.Errorf("expected string to be 'c6', got '%s'", str)
		}
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

func assertFileAndRank(t *testing.T, square domain.Square, file domain.File, rank domain.Rank) {
	t.Helper()

	if square.File() != file {
		t.Errorf("expected file %d, got %d", file, square.File())
	}

	if square.Rank() != rank {
		t.Errorf("expected rank %d, got %d", rank, square.Rank())
	}
}

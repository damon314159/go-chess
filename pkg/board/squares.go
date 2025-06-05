package board

import (
	"fmt"

	"github.com/damon314159/go-chess/pkg/domain"
)

type Square struct {
	file domain.File
	rank domain.Rank
}

var _ domain.Square = (*Square)(nil)

func NewSquare(file domain.File, rank domain.Rank) (Square, error) {
	if file < MinFile || file > MaxFile {
		return Square{}, fmt.Errorf("Square constructed with File out of bounds (%d)", file)
	}
	if rank < MinRank || rank > MaxRank {
		return Square{}, fmt.Errorf("Square constructed with Rank out of bounds (%d)", rank)
	}

	return Square{file, rank}, nil
}

func (s Square) File() domain.File {
	return s.file
}

func (s Square) Rank() domain.Rank {
	return s.rank
}

func (s Square) Translate(right uint8, up uint8) (domain.Square, error) {
	return NewSquare(s.file+domain.File(right), s.rank+domain.Rank(up))
}

func (s Square) Up(spaces uint8) (domain.Square, error) {
	return s.Translate(0, spaces)
}

func (s Square) Down(spaces uint8) (domain.Square, error) {
	return s.Translate(0, -spaces)
}

func (s Square) Left(spaces uint8) (domain.Square, error) {
	return s.Translate(-spaces, 0)
}

func (s Square) Right(spaces uint8) (domain.Square, error) {
	return s.Translate(spaces, 0)
}

func (s Square) String() string {
	files := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return fmt.Sprintf("%s%d", files[s.file-1], s.rank)
}

// Unexported Files used to construct all Squares below
const (
	fA domain.File = iota + 1
	fB
	fC
	fD
	fE
	fF
	fG
	fH
)

// Unexported Ranks used to construct all Squares below
const (
	r1 domain.Rank = iota + 1
	r2
	r3
	r4
	r5
	r6
	r7
	r8
)

// Exported Squares to refer to specific known squares.
// E.g. Start a game with the king on `board.E1`
var (
	A1 Square = Square{fA, r1}
	A2 Square = Square{fA, r2}
	A3 Square = Square{fA, r3}
	A4 Square = Square{fA, r4}
	A5 Square = Square{fA, r5}
	A6 Square = Square{fA, r6}
	A7 Square = Square{fA, r7}
	A8 Square = Square{fA, r8}

	B1 Square = Square{fB, r1}
	B2 Square = Square{fB, r2}
	B3 Square = Square{fB, r3}
	B4 Square = Square{fB, r4}
	B5 Square = Square{fB, r5}
	B6 Square = Square{fB, r6}
	B7 Square = Square{fB, r7}
	B8 Square = Square{fB, r8}

	C1 Square = Square{fC, r1}
	C2 Square = Square{fC, r2}
	C3 Square = Square{fC, r3}
	C4 Square = Square{fC, r4}
	C5 Square = Square{fC, r5}
	C6 Square = Square{fC, r6}
	C7 Square = Square{fC, r7}
	C8 Square = Square{fC, r8}

	D1 Square = Square{fD, r1}
	D2 Square = Square{fD, r2}
	D3 Square = Square{fD, r3}
	D4 Square = Square{fD, r4}
	D5 Square = Square{fD, r5}
	D6 Square = Square{fD, r6}
	D7 Square = Square{fD, r7}
	D8 Square = Square{fD, r8}

	E1 Square = Square{fE, r1}
	E2 Square = Square{fE, r2}
	E3 Square = Square{fE, r3}
	E4 Square = Square{fE, r4}
	E5 Square = Square{fE, r5}
	E6 Square = Square{fE, r6}
	E7 Square = Square{fE, r7}
	E8 Square = Square{fE, r8}

	F1 Square = Square{fF, r1}
	F2 Square = Square{fF, r2}
	F3 Square = Square{fF, r3}
	F4 Square = Square{fF, r4}
	F5 Square = Square{fF, r5}
	F6 Square = Square{fF, r6}
	F7 Square = Square{fF, r7}
	F8 Square = Square{fF, r8}

	G1 Square = Square{fG, r1}
	G2 Square = Square{fG, r2}
	G3 Square = Square{fG, r3}
	G4 Square = Square{fG, r4}
	G5 Square = Square{fG, r5}
	G6 Square = Square{fG, r6}
	G7 Square = Square{fG, r7}
	G8 Square = Square{fG, r8}

	H1 Square = Square{fH, r1}
	H2 Square = Square{fH, r2}
	H3 Square = Square{fH, r3}
	H4 Square = Square{fH, r4}
	H5 Square = Square{fH, r5}
	H6 Square = Square{fH, r6}
	H7 Square = Square{fH, r7}
	H8 Square = Square{fH, r8}
)

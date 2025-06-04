package board

import (
	"github.com/damon314159/go-chess/pkg/domain"
)

type Move struct {
	to   domain.Square
	from domain.Square

	nothingBlocking   []domain.Square
	somethingBlocking []domain.Square
}

var _ domain.Move = (*Move)(nil)

func NewMove(
	to domain.Square,
	from domain.Square,
	nothingBlocking []domain.Square,
	somethingBlocking []domain.Square,
) domain.Move {
	return Move{to, from, nothingBlocking, somethingBlocking}
}

func (m Move) To() domain.Square {
	return m.to
}

func (m Move) From() domain.Square {
	return m.from
}

func (m Move) NothingBlocking() []domain.Square {
	// Ensure nil slices are never returned
	if m.nothingBlocking == nil {
		return []domain.Square{}
	}
	return m.nothingBlocking
}

func (m Move) SomethingBlocking() []domain.Square {
	// Ensure nil slices are never returned
	if m.somethingBlocking == nil {
		return []domain.Square{}
	}
	return m.somethingBlocking
}

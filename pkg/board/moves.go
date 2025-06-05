package board

import (
	"github.com/damon314159/go-chess/pkg/domain"
)

type Move struct {
	from            domain.Square
	to              domain.Square
	nothingBlocking []domain.Square
	captureRequired bool
}

var _ domain.Move = (*Move)(nil)

func NewMove(cfg MoveConfig) domain.Move {
	return Move{cfg.From, cfg.To, cfg.NothingBlocking, cfg.CaptureRequired}
}

type MoveConfig struct {
	From            domain.Square
	To              domain.Square
	NothingBlocking []domain.Square
	CaptureRequired bool
}

func (m Move) From() domain.Square {
	return m.from
}

func (m Move) To() domain.Square {
	return m.to
}

func (m Move) NothingBlocking() []domain.Square {
	// Ensure nil slices are never returned
	if m.nothingBlocking == nil {
		return []domain.Square{}
	}
	// Create a new copy to avoid returning the original slice
	result := make([]domain.Square, len(m.nothingBlocking))
	copy(result, m.nothingBlocking)
	return result
}

func (m Move) CaptureRequired() bool {
	return m.captureRequired
}

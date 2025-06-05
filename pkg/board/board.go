package board

import "github.com/damon314159/go-chess/pkg/domain"

const (
	MinFile domain.File = 1
	MaxFile domain.File = 8
	MinRank domain.Rank = 1
	MaxRank domain.Rank = 8
)

const (
	MaxDiagDistance     int8 = 7
	MaxStraightDistance int8 = 7
)

const (
	WhiteHomeRank domain.Rank = 2
	BlackHomeRank domain.Rank = 7
)

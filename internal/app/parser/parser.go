package parser

import "lr1/internal/app/repository"

type Forward interface {
	Parse(checkedSymptoms []string) (string, []string)
}

type Backward interface {
	Parse(checkedSymptoms []string, diagnose string) (bool, []string)
}

type ChainParser struct {
	Forward
	Backward
	Repository *repository.Repository
}

func NewChainParser(r *repository.Repository) *ChainParser {
	return &ChainParser{
		Repository: r,
		Forward:    NewForwardParser(r),
		Backward:   NewBackwardParser(r),
	}
}

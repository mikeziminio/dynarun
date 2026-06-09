package domain

import (
	"time"
)

type Model struct {
	ID               ID
	Name             string
	RepoID           string
	Filename         string
	InputTokenPrice  int64
	OutputTokenPrice int64
	UpdatedAt        time.Time
}

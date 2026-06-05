package domain

import (
	"time"
)

type Model struct {
	Id               ID
	Name             string
	RepoId           string
	Filename         string
	InputTokenPrice  int64
	OutputTokenPrice int64
	UpdatedAt        time.Time
}

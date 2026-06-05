package domain

import "github.com/google/uuid"

type ID = uuid.UUID

func ParseID(s string) (ID, error) {
	return uuid.Parse(s)
}

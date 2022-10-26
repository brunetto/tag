package tag

import "github.com/google/uuid"

func NewFakeFixedUUID() uuid.UUID {
	return uuid.UUID([16]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
}

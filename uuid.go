package tag

import (
	"github.com/rs/xid"
)

type UUID string

func NewFakeFixedUUID() UUID {
	return UUID("fake-uuid")
}

func NewUUID() string {
	return xid.New().String()
}

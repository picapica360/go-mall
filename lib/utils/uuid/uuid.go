package uuid

import (
	"github.com/google/uuid"
)

// NewUUID create a new UUIDv4, eg: c01d7cf6-ec3f-47f0-9556-a5d6e9009a43
func NewUUID() string {
	return uuid.New().String()
}

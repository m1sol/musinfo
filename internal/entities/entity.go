package entities

import "github.com/google/uuid"

type Entity interface {
	GetID() uuid.UUID
}

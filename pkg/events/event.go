package events

import "github.com/google/uuid"

type DomainEvent interface {
	Name() string
	EventID() uuid.UUID
}

type EventHandler interface {
	Notify(event DomainEvent)
}

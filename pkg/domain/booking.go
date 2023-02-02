package domain

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID          uuid.UUID
	date        time.Time
	Supplier    string
	Description string
	Category    int
	Amount      float32
}

type BookingSvc interface {
	Get(id uuid.UUID) (*Booking, error)
	List(category string) ([]*Booking, error)
	Create(b *Booking) error
	Delete(id uuid.UUID) error
}

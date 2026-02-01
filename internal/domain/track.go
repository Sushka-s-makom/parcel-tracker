package domain

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusNew       Status = "new"
	StatusInTransit Status = "in_transit"
	StatusArrived   Status = "arrived"
	StatusFailed    Status = "failed"
)

type Track struct {
	ID           uuid.UUID `json:"id"`
	TrackNumber  string    `json:"track_number"`
	Carrier      string    `json:"carrier"`
	Status       Status    `json:"status"`
	LastUpdateAt time.Time `json:"last_update_at"`
	CreatedAt    time.Time `json:"created_at"`
}

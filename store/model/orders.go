package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id          int64
	CustomerId  uuid.UUID
	LineItems   []LineItem
	CreatedAt   time.Time
	ShippedAt   *time.Time
	CompletedAt *time.Time
}

type LineItem struct {
	ItemId   uuid.UUID
	Price    int
	Quantity int
}

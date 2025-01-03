package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id          int64      `json:"id"`
	CustomerId  uuid.UUID  `json:"customer_id"`
	LineItems   []LineItem `json:"line_items"`
	CreatedAt   time.Time  `json:"created_at"`
	ShippedAt   *time.Time `json:"shipped_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type LineItem struct {
	ItemId   uuid.UUID `json:"item_id"`
	Price    int       `json:"price"`
	Quantity int       `json:"quantity"`
}

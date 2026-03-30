package model

import (
	"time"
)

// Stock is a row in stock (one current quantity per product).
type Stock struct {
	ID        int64     `db:"id"`
	ProductID int64     `db:"product_id"`
	Quantity  float64   `db:"quantity"`
	UpdatedAt time.Time `db:"updated_at"`
}

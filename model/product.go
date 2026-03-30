package model

import "time"

type Product struct {
	ID          int64     `db:"id"`
	SKU         string    `db:"sku"`
	Name        string    `db:"name"`
	Description *string   `db:"description"`
	Unit        string    `db:"unit"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

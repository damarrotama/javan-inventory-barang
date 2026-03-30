package model

import (
	"javan-inventory-barang/config"
	"time"
)

type Product struct {
	ID          int64     `db:"id"`
	SKU         string    `db:"sku"`
	Name        string    `db:"name"`
	Description *string   `db:"description"`
	Unit        string    `db:"unit"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func InsertProduct(p *Product) error {
	query := `
		INSERT INTO products (sku, name, description, unit)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	return config.DB.QueryRow(
		query, p.SKU, p.Name, p.Description, p.Unit,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

package model

import "time"

// Product maps to the products table.
type Product struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	SKU         string    `gorm:"type:varchar(64);uniqueIndex;not null" json:"sku"`
	Name        string    `gorm:"type:varchar(255);not null;index" json:"name"`
	Description *string   `gorm:"type:text" json:"description,omitempty"`
	Unit        string    `gorm:"type:varchar(32);not null;default:pcs" json:"unit"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName returns the GORM table name.
func (Product) TableName() string {
	return "products"
}

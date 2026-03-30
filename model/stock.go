package model

import "github.com/google/uuid"

type Stock struct {
	Base
	ProductID *uuid.UUID `gorm:"type:varchar(36);not null;uniqueIndex" json:"product_id"`
	Quantity  float64    `gorm:"not null;default:0" json:"quantity"`
}

// TableName returns the GORM table name.
func (Stock) TableName() string {
	return "stocks"
}

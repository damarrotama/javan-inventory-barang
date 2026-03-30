package model

import "github.com/google/uuid"

type StockHistory struct {
	Base
	ProductID     *uuid.UUID        `gorm:"not null;index" json:"product_id"`
	StockID       *uuid.UUID        `gorm:"index" json:"stock_id,omitempty"`
	MovementType  StockMovementType `gorm:"type:varchar(32);not null;index" json:"movement_type"`
	QuantityDelta float64           `gorm:"not null" json:"quantity_delta"`
	QuantityAfter float64           `gorm:"not null" json:"quantity_after"`
	Reference     *string           `gorm:"type:varchar(255)" json:"reference,omitempty"`
	Note          *string           `gorm:"type:text" json:"note,omitempty"`
}

type StockMovementType string

const (
	StockMovementIn         StockMovementType = "IN"
	StockMovementOut        StockMovementType = "OUT"
	StockMovementAdjustment StockMovementType = "ADJUSTMENT"
)

// TableName returns the GORM table name.
func (StockHistory) TableName() string {
	return "stock_histories"
}

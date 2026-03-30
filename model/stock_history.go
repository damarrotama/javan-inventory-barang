package model

import (
	"time"
)

type StockHistory struct {
	ID            int64             `db:"id"`
	ProductID     int64             `db:"product_id"`
	StockID       *int64            `db:"stock_id"`
	MovementType  StockMovementType `db:"movement_type"`
	QuantityDelta float64           `db:"quantity_delta"`
	QuantityAfter float64           `db:"quantity_after"`
	Reference     *string           `db:"reference"`
	Note          *string           `db:"note"`
	CreatedAt     time.Time         `db:"created_at"`
}

type StockMovementType string

const (
	StockMovementIn         StockMovementType = "IN"
	StockMovementOut        StockMovementType = "OUT"
	StockMovementAdjustment StockMovementType = "ADJUSTMENT"
)

package repository

import (
	"context"
	"javan-inventory-barang/model"
)

type StockHistoryRepository interface {
	FindAll(ctx context.Context) ([]model.StockHistory, error)
}

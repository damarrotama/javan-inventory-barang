package repository

import (
	"context"
	"javan-inventory-barang/model"
)

type StockRepository interface {
	FindAll(ctx context.Context) ([]model.Stock, error)
}

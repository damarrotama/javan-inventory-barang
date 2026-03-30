package repository

import (
	"context"
	"javan-inventory-barang/model"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]model.Product, error)
}

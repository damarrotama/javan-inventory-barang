package repository

import (
	"context"
	"javan-inventory-barang/model"
	"javan-inventory-barang/transaction"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StockRepository handles data access for the stocks table.
type StockRepository interface {
	FindAll(ctx context.Context) ([]model.Stock, error)
	FindByID(ctx context.Context, id *uuid.UUID) (*model.Stock, error)
	FindByProductID(ctx context.Context, productID *uuid.UUID) (*model.Stock, error)
	Create(ctx context.Context, stock *model.Stock) error
	Update(ctx context.Context, stock *model.Stock) error

	WithTx(tx transaction.Conn) StockRepository
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) WithTx(tx transaction.Conn) StockRepository {
	return &stockRepository{db: tx.Tx}
}

func (r *stockRepository) FindAll(ctx context.Context) ([]model.Stock, error) {
	var stocks []model.Stock
	if err := r.db.WithContext(ctx).Find(&stocks).Error; err != nil {
		return nil, err
	}
	return stocks, nil
}

func (r *stockRepository) FindByID(ctx context.Context, id *uuid.UUID) (*model.Stock, error) {
	if id == nil {
		return nil, gorm.ErrRecordNotFound
	}
	var stock model.Stock
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&stock).Error; err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *stockRepository) FindByProductID(ctx context.Context, productID *uuid.UUID) (*model.Stock, error) {
	if productID == nil {
		return nil, gorm.ErrRecordNotFound
	}
	var stock model.Stock
	if err := r.db.WithContext(ctx).Where("product_id = ?", productID).First(&stock).Error; err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *stockRepository) Create(ctx context.Context, stock *model.Stock) error {
	return r.db.WithContext(ctx).Create(stock).Error
}

func (r *stockRepository) Update(ctx context.Context, stock *model.Stock) error {
	return r.db.WithContext(ctx).Save(stock).Error
}

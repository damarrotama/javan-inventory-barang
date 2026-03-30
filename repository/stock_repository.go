package repository

import (
	"context"
	"javan-inventory-barang/model"
	"javan-inventory-barang/transaction"
	"javan-inventory-barang/utils/logger"

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
	logger logger.Logger
	db     *gorm.DB
}

func NewStockRepository(logger logger.Logger, db *gorm.DB) StockRepository {
	return &stockRepository{logger: logger, db: db}
}

func (r *stockRepository) WithTx(tx transaction.Conn) StockRepository {
	return &stockRepository{logger: r.logger, db: tx.Tx}
}

func (r *stockRepository) FindAll(ctx context.Context) ([]model.Stock, error) {
	var stocks []model.Stock
	if err := r.db.WithContext(ctx).Find(&stocks).Error; err != nil {
		r.logger.Error(ctx, "failed to find all stocks", "error", err)
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
		r.logger.Error(ctx, "failed to find stock by id", "error", err)
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
		r.logger.Error(ctx, "failed to find stock by product id", "error", err)
		return nil, err
	}
	return &stock, nil
}

func (r *stockRepository) Create(ctx context.Context, stock *model.Stock) error {
	if err := r.db.WithContext(ctx).Create(stock).Error; err != nil {
		r.logger.Error(ctx, "failed to create stock", "error", err)
		return err
	}
	return nil
}

func (r *stockRepository) Update(ctx context.Context, stock *model.Stock) error {
	if err := r.db.WithContext(ctx).Save(stock).Error; err != nil {
		r.logger.Error(ctx, "failed to update stock", "error", err)
		return err
	}
	return nil
}

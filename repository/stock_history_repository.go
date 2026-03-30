package repository

import (
	"context"
	"javan-inventory-barang/model"
	"javan-inventory-barang/transaction"
	"javan-inventory-barang/utils/logger"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StockHistoryRepository handles data access for the stock_histories table.
type StockHistoryRepository interface {
	FindAll(ctx context.Context) ([]model.StockHistory, error)
	FindByProductID(ctx context.Context, productID *uuid.UUID) ([]model.StockHistory, error)
	FindByStockID(ctx context.Context, stockID *uuid.UUID) ([]model.StockHistory, error)
	Create(ctx context.Context, history *model.StockHistory) error

	WithTx(tx transaction.Conn) StockHistoryRepository
}

type stockHistoryRepository struct {
	logger logger.Logger
	db     *gorm.DB
}

func NewStockHistoryRepository(logger logger.Logger, db *gorm.DB) StockHistoryRepository {
	return &stockHistoryRepository{logger: logger, db: db}
}

func (r *stockHistoryRepository) WithTx(tx transaction.Conn) StockHistoryRepository {
	return &stockHistoryRepository{logger: r.logger, db: tx.Tx}
}

func (r *stockHistoryRepository) FindAll(ctx context.Context) ([]model.StockHistory, error) {
	var histories []model.StockHistory
	if err := r.db.WithContext(ctx).Order("created_at DESC").Find(&histories).Error; err != nil {
		r.logger.Error(ctx, "failed to find all stock histories", "error", err)
		return nil, err
	}
	return histories, nil
}

func (r *stockHistoryRepository) FindByProductID(ctx context.Context, productID *uuid.UUID) ([]model.StockHistory, error) {
	if productID == nil {
		return nil, gorm.ErrRecordNotFound
	}
	var histories []model.StockHistory
	if err := r.db.WithContext(ctx).Where("product_id = ?", productID).Order("created_at DESC").Find(&histories).Error; err != nil {
		r.logger.Error(ctx, "failed to find stock histories by product id", "error", err)
		return nil, err
	}
	return histories, nil
}

func (r *stockHistoryRepository) FindByStockID(ctx context.Context, stockID *uuid.UUID) ([]model.StockHistory, error) {
	if stockID == nil {
		return nil, gorm.ErrRecordNotFound
	}
	var histories []model.StockHistory
	if err := r.db.WithContext(ctx).Where("stock_id = ?", stockID).Order("created_at DESC").Find(&histories).Error; err != nil {
		r.logger.Error(ctx, "failed to find stock histories by stock id", "error", err)
		return nil, err
	}
	return histories, nil
}

func (r *stockHistoryRepository) Create(ctx context.Context, history *model.StockHistory) error {
	if err := r.db.WithContext(ctx).Create(history).Error; err != nil {
		r.logger.Error(ctx, "failed to create stock history", "error", err)
		return err
	}
	return nil
}

package repository

import (
	"context"
	"javan-inventory-barang/model"
	"javan-inventory-barang/transaction"
	"javan-inventory-barang/utils/logger"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]model.Product, error)
	FindById(ctx context.Context, id *uuid.UUID) (*model.Product, error)
	Create(ctx context.Context, product *model.Product) error
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, id *uuid.UUID) error

	WithTx(tx transaction.Conn) ProductRepository
}

type productRepository struct {
	logger logger.Logger
	db     *gorm.DB
}

func NewProductRepository(logger logger.Logger, db *gorm.DB) ProductRepository {
	return &productRepository{logger: logger, db: db}
}

func (r *productRepository) WithTx(tx transaction.Conn) ProductRepository {
	return &productRepository{logger: r.logger, db: tx.Tx}
}

func (r *productRepository) FindAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
		r.logger.Error(ctx, "failed to find all products", "error", err)
		return nil, err
	}
	return products, nil
}

func (r *productRepository) FindById(ctx context.Context, id *uuid.UUID) (*model.Product, error) {
	if id == nil {
		return nil, gorm.ErrRecordNotFound
	}
	var product model.Product
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&product).Error; err != nil {
		r.logger.Error(ctx, "failed to find product by id", "error", err)
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	if err := r.db.WithContext(ctx).Create(product).Error; err != nil {
		r.logger.Error(ctx, "failed to create product", "error", err)
		return err
	}
	return nil
}

func (r *productRepository) Update(ctx context.Context, product *model.Product) error {
	if err := r.db.WithContext(ctx).Save(product).Error; err != nil {
		r.logger.Error(ctx, "failed to update product", "error", err)
		return err
	}

	return nil
}

func (r *productRepository) Delete(ctx context.Context, id *uuid.UUID) error {
	if id == nil {
		r.logger.Error(ctx, "id is nil")
		return gorm.ErrRecordNotFound
	}
	if err := r.db.WithContext(ctx).Delete(&model.Product{}, id).Error; err != nil {
		r.logger.Error(ctx, "failed to delete product", "error", err)
		return err
	}
	return r.db.WithContext(ctx).Delete(&model.Product{}, id).Error
}

package repository

import (
	"context"
	"javan-inventory-barang/model"
	"javan-inventory-barang/transaction"

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
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) WithTx(tx transaction.Conn) ProductRepository {
	return &productRepository{db: tx.Tx}
}

func (r *productRepository) FindAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
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
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepository) Update(ctx context.Context, product *model.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *productRepository) Delete(ctx context.Context, id *uuid.UUID) error {
	if id == nil {
		return gorm.ErrRecordNotFound
	}
	return r.db.WithContext(ctx).Delete(&model.Product{}, id).Error
}

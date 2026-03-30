package domain

import (
	"context"
	"javan-inventory-barang/model"
	"javan-inventory-barang/repository"
)

// ProductDomain defines product use cases.
type ProductDomain interface {
	GetProducts(ctx context.Context) ([]model.Product, error)
	GetProductByID(ctx context.Context, id int64) (*model.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) error
	UpdateProduct(ctx context.Context, product *model.Product) error
	DeleteProduct(ctx context.Context, id int64) error
}

type productDomain struct {
	repo repository.ProductRepository
}

// NewProductDomain constructs a ProductDomain.
func NewProductDomain(repo repository.ProductRepository) ProductDomain {
	return &productDomain{repo: repo}
}

func (d *productDomain) GetProducts(ctx context.Context) ([]model.Product, error) {
	return d.repo.FindAll(ctx)
}

func (d *productDomain) GetProductByID(ctx context.Context, id int64) (*model.Product, error) {
	return d.repo.FindById(ctx, id)
}

func (d *productDomain) CreateProduct(ctx context.Context, product *model.Product) error {
	return d.repo.Create(ctx, product)
}

func (d *productDomain) UpdateProduct(ctx context.Context, product *model.Product) error {
	if _, err := d.repo.FindById(ctx, product.ID); err != nil {
		return err
	}
	return d.repo.Update(ctx, product)
}

func (d *productDomain) DeleteProduct(ctx context.Context, id int64) error {
	if _, err := d.repo.FindById(ctx, id); err != nil {
		return err
	}
	return d.repo.Delete(ctx, id)
}

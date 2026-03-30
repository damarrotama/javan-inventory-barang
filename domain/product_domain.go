package domain

import (
	"context"

	"github.com/google/uuid"

	"javan-inventory-barang/model"
	"javan-inventory-barang/repository"
)

// ProductDomain defines product use cases.
type ProductDomain interface {
	GetProducts(ctx context.Context) ([]model.Product, error)
	GetProductByID(ctx context.Context, id *uuid.UUID) (*model.Product, error)
	CreateProduct(ctx context.Context, api *model.ProductAPI) (*model.Product, error)
	UpdateProduct(ctx context.Context, api *model.ProductAPI, id *uuid.UUID) (*model.Product, error)
	DeleteProduct(ctx context.Context, id *uuid.UUID) error
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

func (d *productDomain) GetProductByID(ctx context.Context, id *uuid.UUID) (*model.Product, error) {
	return d.repo.FindById(ctx, id)
}

func (d *productDomain) CreateProduct(ctx context.Context, api *model.ProductAPI) (*model.Product, error) {
	product := &model.Product{
		ProductAPI: *api,
	}
	if err := d.repo.Create(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}

func (d *productDomain) UpdateProduct(ctx context.Context, api *model.ProductAPI, id *uuid.UUID) (*model.Product, error) {
	product := &model.Product{
		Base:       model.Base{ID: id},
		ProductAPI: *api,
	}
	if _, err := d.repo.FindById(ctx, product.ID); err != nil {
		return nil, err
	}
	if err := d.repo.Update(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}

func (d *productDomain) DeleteProduct(ctx context.Context, id *uuid.UUID) error {
	if _, err := d.repo.FindById(ctx, id); err != nil {
		return err
	}
	return d.repo.Delete(ctx, id)
}

package domain

import (
	"context"

	"github.com/google/uuid"

	"javan-inventory-barang/model"
	"javan-inventory-barang/repository"
	"javan-inventory-barang/transaction"
	"javan-inventory-barang/utils/resp"
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
	txManager transaction.Manager
	repo      repository.ProductRepository
}

// NewProductDomain constructs a ProductDomain.
func NewProductDomain(txManager transaction.Manager, repo repository.ProductRepository) ProductDomain {
	return &productDomain{txManager: txManager, repo: repo}
}

func (d *productDomain) GetProducts(ctx context.Context) ([]model.Product, error) {
	products, err := d.repo.FindAll(ctx)
	if err != nil {
		return nil, resp.ErrorInternal()
	}
	return products, nil
}

func (d *productDomain) GetProductByID(ctx context.Context, id *uuid.UUID) (*model.Product, error) {
	product, err := d.repo.FindById(ctx, id)
	if err != nil {
		return nil, resp.ErrorInternal(err.Error())
	}
	return product, nil
}

func (d *productDomain) CreateProduct(ctx context.Context, api *model.ProductAPI) (*model.Product, error) {
	product := &model.Product{
		ProductAPI: *api,
	}
	if err := d.txManager.WithTx(ctx, func(tx transaction.Conn) error {
		return d.repo.WithTx(tx).Create(ctx, product)
	}); err != nil {
		return nil, resp.ErrorInternal()
	}
	return product, nil
}

func (d *productDomain) UpdateProduct(ctx context.Context, api *model.ProductAPI, id *uuid.UUID) (*model.Product, error) {
	if _, err := d.repo.FindById(ctx, id); err != nil {
		return nil, resp.ErrorNotFound()
	}

	product := &model.Product{
		Base:       model.Base{ID: id},
		ProductAPI: *api,
	}
	if err := d.txManager.WithTx(ctx, func(tx transaction.Conn) error {
		return d.repo.WithTx(tx).Update(ctx, product)
	}); err != nil {
		return nil, resp.ErrorInternal()
	}
	return product, nil
}

func (d *productDomain) DeleteProduct(ctx context.Context, id *uuid.UUID) error {
	if _, err := d.repo.FindById(ctx, id); err != nil {
		return resp.ErrorNotFound()
	}
	if err := d.txManager.WithTx(ctx, func(tx transaction.Conn) error {
		return d.repo.WithTx(tx).Delete(ctx, id)
	}); err != nil {
		return resp.ErrorInternal()
	}
	return nil
}

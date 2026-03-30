package domain

import "javan-inventory-barang/model"

type ProductDomain interface {
	GetProducts() ([]model.Product, error)
}

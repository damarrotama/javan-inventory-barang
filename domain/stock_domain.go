package domain

import "javan-inventory-barang/model"

type StockDomain interface {
	GetStocks() ([]model.Stock, error)
}

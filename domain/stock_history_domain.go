package domain

import "javan-inventory-barang/model"

type StockHistoryDomain interface {
	GetStockHistory() ([]model.Stock, error)
}

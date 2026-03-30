package migrations

import "javan-inventory-barang/model"

var ModelMigrations = []interface{}{
	&model.Product{},
	&model.Stock{},
	&model.StockHistory{},
}

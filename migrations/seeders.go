package migrations

import "javan-inventory-barang/model"

var (
	product model.Product
)

func DataSeeds() []any {
	return []any{
		product.Seed(),
	}
}

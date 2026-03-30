package config

var Environment map[string]any = map[string]any{
	"ENV":             "local",
	"PORT":            8000,
	"DB_HOST":         "localhost",
	"DB_PORT":         5432,
	"DB_USER":         "postgres",
	"DB_PASSWORD":     "postgres",
	"DB_NAME":         "inventory_barang",
	"DB_AUTO_MIGRATE": true,
	"DB_TABLE_PREFIX": "",
}

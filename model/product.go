package model

import "javan-inventory-barang/lib"

// Product maps to the products table.
type Product struct {
	Base
	ProductAPI
}
type ProductAPI struct {
	SKU         string  `gorm:"type:varchar(64);uniqueIndex;not null" json:"sku"`
	Name        string  `gorm:"type:varchar(255);not null;index" json:"name"`
	Description *string `gorm:"type:text" json:"description,omitempty"`
	Unit        string  `gorm:"type:varchar(32);not null;default:pcs" json:"unit"`
}

// TableName returns the GORM table name.
func (Product) TableName() string {
	return "products"
}

func (b *Product) Seed() *[]Product {
	return &[]Product{
		{
			ProductAPI: ProductAPI{
				SKU:         "ELC-0001",
				Name:        "USB-C Charging Cable",
				Description: lib.Pointer("Durable USB-C charging cable for fast and reliable power"),
				Unit:        "pcs",
			},
		},
		{
			ProductAPI: ProductAPI{
				SKU:         "ELC-0002",
				Name:        "20W USB-C Wall Charger",
				Description: lib.Pointer("Compact 20W USB-C wall charger with stable output"),
				Unit:        "pcs",
			},
		},
		{
			ProductAPI: ProductAPI{
				SKU:         "ELC-0003",
				Name:        "Bluetooth Earbuds",
				Description: lib.Pointer("Bluetooth earbuds with clear calls and comfortable fit"),
				Unit:        "pcs",
			},
		},
		{
			ProductAPI: ProductAPI{
				SKU:         "ELC-0004",
				Name:        "10,000mAh Power Bank",
				Description: lib.Pointer("10,000mAh power bank for charging phones on the go"),
				Unit:        "pcs",
			},
		},
		{
			ProductAPI: ProductAPI{
				SKU:         "ELC-0005",
				Name:        "Wireless Optical Mouse",
				Description: lib.Pointer("2.4G wireless optical mouse for smooth tracking"),
				Unit:        "pcs",
			},
		},
		{
			ProductAPI: ProductAPI{
				SKU:         "ELC-0006",
				Name:        "Adjustable Phone Stand",
				Description: lib.Pointer("Adjustable phone stand for desk or bedside viewing"),
				Unit:        "pcs",
			},
		},
		{
			ProductAPI: ProductAPI{
				SKU:         "ELC-0007",
				Name:        "HDMI Adapter",
				Description: lib.Pointer("HDMI adapter for connecting compatible devices to displays"),
				Unit:        "pcs",
			},
		},
		{
			ProductAPI: ProductAPI{
				SKU:         "ELC-0008",
				Name:        "USB Flash Drive 64GB",
				Description: lib.Pointer("High-speed USB flash drive (64GB) for backups and transfers"),
				Unit:        "pcs",
			},
		},
	}
}

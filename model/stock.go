package model

type Stock struct {
	Base
	ProductID int64   `gorm:"not null;uniqueIndex" json:"product_id"`
	Quantity  float64 `gorm:"not null;default:0" json:"quantity"`
}

// TableName returns the GORM table name.
func (Stock) TableName() string {
	return "stocks"
}

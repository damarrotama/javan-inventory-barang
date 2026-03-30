package transaction

import (
	"context"

	"gorm.io/gorm"
)

type Manager interface {
	WithTx(ctx context.Context, fn func(Conn) error) error
}

type Conn struct {
	Tx *gorm.DB
}

type gormAdapter struct {
	tx *gorm.DB
}

func NewManager(db *gorm.DB) Manager {
	return &gormAdapter{tx: db}
}

func (a *gormAdapter) WithTx(ctx context.Context, fn func(tx Conn) error) error {
	return a.tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(Conn{Tx: tx})
	})
}

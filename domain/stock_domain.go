package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"javan-inventory-barang/lib"
	"javan-inventory-barang/model"
	"javan-inventory-barang/repository"
	"javan-inventory-barang/utils/logger"
	"javan-inventory-barang/utils/transaction"
)

// Sentinel errors for stock movement validation.
var (
	ErrInvalidMovementType = errors.New("movement_type must be IN or OUT")
	ErrInvalidQuantity     = errors.New("quantity must be greater than zero")
	ErrInsufficientStock   = errors.New("insufficient stock: quantity cannot go negative")
	ErrProductIDRequired   = errors.New("product_id is required")
)

// StockMovementRequest is the payload for stock IN/OUT operations.
type StockMovementRequest struct {
	ProductID    *uuid.UUID              `json:"product_id"`
	MovementType model.StockMovementType `json:"movement_type"`
	Quantity     float64                 `json:"quantity"`
	Reference    *string                 `json:"reference,omitempty"`
	Note         *string                 `json:"note,omitempty"`
}

// StockMovementResponse is returned after a successful stock movement.
type StockMovementResponse struct {
	Stock   *model.Stock `json:"stock"`
	Message string       `json:"message"`
}

// StockDomain defines use cases for stock and history management.
type StockDomain interface {
	GetStocks(ctx context.Context) ([]model.Stock, error)
	GetStockByProductID(ctx context.Context, productID *uuid.UUID) (*model.Stock, error)
	GetStockHistories(ctx context.Context) ([]model.StockHistory, error)
	GetStockHistoriesByProductID(ctx context.Context, productID *uuid.UUID) ([]model.StockHistory, error)
	MoveStock(ctx context.Context, req *StockMovementRequest) (*StockMovementResponse, error)
}

type stockDomain struct {
	logger           logger.Logger
	txManager        transaction.Manager
	stockRepo        repository.StockRepository
	stockHistoryRepo repository.StockHistoryRepository
	productRepo      repository.ProductRepository
}

func NewStockDomain(
	logger logger.Logger,
	txManager transaction.Manager,
	stockRepo repository.StockRepository,
	stockHistoryRepo repository.StockHistoryRepository,
	productRepo repository.ProductRepository,
) StockDomain {
	return &stockDomain{
		logger:           logger,
		txManager:        txManager,
		stockRepo:        stockRepo,
		stockHistoryRepo: stockHistoryRepo,
		productRepo:      productRepo,
	}
}

func (d *stockDomain) GetStocks(ctx context.Context) ([]model.Stock, error) {
	return d.stockRepo.FindAll(ctx)
}

func (d *stockDomain) GetStockByProductID(ctx context.Context, productID *uuid.UUID) (*model.Stock, error) {
	return d.stockRepo.FindByProductID(ctx, productID)
}

func (d *stockDomain) GetStockHistories(ctx context.Context) ([]model.StockHistory, error) {
	return d.stockHistoryRepo.FindAll(ctx)
}

func (d *stockDomain) GetStockHistoriesByProductID(ctx context.Context, productID *uuid.UUID) ([]model.StockHistory, error) {
	return d.stockHistoryRepo.FindByProductID(ctx, productID)
}

// MoveStock processes a stock IN or OUT movement for a product.
// Flow: validate input → check product exists → transaction (find/create stock + update quantity) → record history async.
func (d *stockDomain) MoveStock(ctx context.Context, req *StockMovementRequest) (*StockMovementResponse, error) {
	if req.ProductID == nil {
		return nil, ErrProductIDRequired
	}
	if req.Quantity <= 0 {
		return nil, ErrInvalidQuantity
	}
	if req.MovementType != model.StockMovementIn && req.MovementType != model.StockMovementOut {
		return nil, ErrInvalidMovementType
	}

	// Ensure the product exists
	if _, err := d.productRepo.FindById(ctx, req.ProductID); err != nil {
		return nil, err
	}

	// Atomic stock update inside a transaction to prevent race conditions
	var updatedStock *model.Stock

	err := d.txManager.WithTx(ctx, func(conn transaction.Conn) error {
		txStockRepo := d.stockRepo.WithTx(conn)

		// Find existing stock or create a new one with quantity 0
		stock, err := txStockRepo.FindByProductID(ctx, req.ProductID)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			stock = &model.Stock{
				ProductID: req.ProductID,
				Quantity:  0,
			}
			if err := txStockRepo.Create(ctx, stock); err != nil {
				return err
			}
		}

		// IN adds, OUT subtracts
		delta := req.Quantity
		if req.MovementType == model.StockMovementOut {
			delta = -delta
		}

		newQuantity := stock.Quantity + delta

		// Stock must not go below zero
		if newQuantity < 0 {
			return ErrInsufficientStock
		}

		stock.Quantity = newQuantity
		if err := txStockRepo.Update(ctx, stock); err != nil {
			return err
		}

		updatedStock = stock
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Record history asynchronously so it doesn't block the response
	go func() {
		defer lib.Recover()

		history := &model.StockHistory{
			ProductID:     req.ProductID,
			StockID:       updatedStock.ID,
			MovementType:  req.MovementType,
			QuantityDelta: req.Quantity,
			QuantityAfter: updatedStock.Quantity,
			Reference:     req.Reference,
			Note:          req.Note,
		}

		if err := d.stockHistoryRepo.Create(context.Background(), history); err != nil {
			d.logger.Error(ctx, "failed to record history", "error", err)
		}
	}()

	return &StockMovementResponse{
		Stock:   updatedStock,
		Message: "stock movement recorded successfully",
	}, nil
}

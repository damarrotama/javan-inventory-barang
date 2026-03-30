package controller

import (
	"errors"
	"javan-inventory-barang/domain"
	"javan-inventory-barang/lib"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// StockController handles HTTP requests for stock and history operations.
type StockController interface {
	GetStocks(ctx *fiber.Ctx) error
	GetStockByProductID(ctx *fiber.Ctx) error
	GetStockHistories(ctx *fiber.Ctx) error
	GetStockHistoriesByProductID(ctx *fiber.Ctx) error
	MoveStock(ctx *fiber.Ctx) error
}

type stockController struct {
	stockDomain domain.StockDomain
}

func NewStockController(stockDomain domain.StockDomain) StockController {
	return &stockController{stockDomain: stockDomain}
}

// GetStocks godoc
// @Summary List all current stock levels
// @Tags stocks
// @Produce json
// @Success 200 {array} model.Stock
// @Failure 500 {object} map[string]string
// @Router /stocks [get]
func (c *stockController) GetStocks(ctx *fiber.Ctx) error {
	stocks, err := c.stockDomain.GetStocks(ctx.UserContext())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(stocks)
}

// GetStockByProductID godoc
// @Summary Get current stock level for a product
// @Tags stocks
// @Produce json
// @Param product_id path string true "Product ID" format(uuid)
// @Success 200 {object} model.Stock
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stocks/product/{product_id} [get]
func (c *stockController) GetStockByProductID(ctx *fiber.Ctx) error {
	stock, err := c.stockDomain.GetStockByProductID(ctx.UserContext(), lib.ParamsUUID(ctx, "product_id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "stock not found for this product"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(stock)
}

// GetStockHistories godoc
// @Summary List all stock movement history
// @Tags stock-histories
// @Produce json
// @Success 200 {array} model.StockHistory
// @Failure 500 {object} map[string]string
// @Router /stocks/histories [get]
func (c *stockController) GetStockHistories(ctx *fiber.Ctx) error {
	histories, err := c.stockDomain.GetStockHistories(ctx.UserContext())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(histories)
}

// GetStockHistoriesByProductID godoc
// @Summary List stock movement history for a product
// @Tags stock-histories
// @Produce json
// @Param product_id path string true "Product ID" format(uuid)
// @Success 200 {array} model.StockHistory
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stocks/histories/product/{product_id} [get]
func (c *stockController) GetStockHistoriesByProductID(ctx *fiber.Ctx) error {
	histories, err := c.stockDomain.GetStockHistoriesByProductID(ctx.UserContext(), lib.ParamsUUID(ctx, "product_id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no history found for this product"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(histories)
}

// MoveStock godoc
// @Summary Record a stock IN or OUT movement
// @Tags stocks
// @Accept json
// @Produce json
// @Param body body domain.StockMovementRequest true "Movement payload"
// @Success 200 {object} domain.StockMovementResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stocks/movement [post]
func (c *stockController) MoveStock(ctx *fiber.Ctx) error {
	req := new(domain.StockMovementRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
	}

	result, err := c.stockDomain.MoveStock(ctx.UserContext(), req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
		}
		if errors.Is(err, domain.ErrInvalidMovementType) ||
			errors.Is(err, domain.ErrInvalidQuantity) ||
			errors.Is(err, domain.ErrProductIDRequired) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, domain.ErrInsufficientStock) {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(result)
}

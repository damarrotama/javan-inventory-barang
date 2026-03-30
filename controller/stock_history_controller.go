package controller

import "github.com/gofiber/fiber/v2"

type StockHistoryController struct {
}

func NewStockHistoryController() *StockHistoryController {
	return &StockHistoryController{}
}

func (c *StockHistoryController) GetStockHistory(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

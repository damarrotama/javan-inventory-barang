package controller

import "github.com/gofiber/fiber/v2"

type StockController struct {
}

func NewStockController() *StockController {
	return &StockController{}
}

func (c *StockController) GetStocks(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello World!",
	})
}

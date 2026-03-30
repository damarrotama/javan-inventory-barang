package controller

import "github.com/gofiber/fiber/v2"

type ProductController struct {
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (c *ProductController) GetProducts(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

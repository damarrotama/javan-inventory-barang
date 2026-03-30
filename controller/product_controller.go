package controller

import (
	"javan-inventory-barang/model"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (c *ProductController) AddProducts(ctx *fiber.Ctx) error {
	var product model.Product

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := model.InsertProduct(&product)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(201).JSON(product)
}

func (c *ProductController) GetProducts(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

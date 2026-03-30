package routes

import (
	"javan-inventory-barang/services"

	"github.com/gofiber/fiber/v2"
)

func Handle(app *fiber.App, service *services.Service) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api/v1")

	productController := service.Controller.ProductController
	productAPI := api.Group("/products")
	productAPI.Get("/", productController.GetProducts)
	// productAPI.Get("/:id", controller.GetProductById)
	// productAPI.Post("/", controller.CreateProduct)
	// productAPI.Put("/:id", controller.UpdateProduct)
	// productAPI.Delete("/:id", controller.DeleteProduct)
}

package routes

import (
	"javan-inventory-barang/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Handle(app *fiber.App, service *services.Service) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api/v1")

	pc := service.Controller.ProductController
	products := api.Group("/products")
	products.Get("/", pc.GetProducts)
	products.Post("/", pc.CreateProduct)
	products.Get("/:id", pc.GetProductByID)
	products.Put("/:id", pc.UpdateProduct)
	products.Delete("/:id", pc.DeleteProduct)
}

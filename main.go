package main

import (
	"javan-inventory-barang/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()

	routes.Handle(app)
	if err := app.Listen(":" + viper.GetString("PORT")); err != nil {
		log.Fatal(err)
	}
}

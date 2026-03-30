package main

import (
	"javan-inventory-barang/config"
	"javan-inventory-barang/routes"
	"javan-inventory-barang/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// @title Javan Inventory Barang API
// @version 1.0.0
// @description Javan Inventory Barang API Documentation
// @contact.name Team 1 Javan Inventory Barang
// @contact.email team1@javan.co.id
// @host localhost:8000
// @schemes http
// @BasePath /api/v1
func main() {
	app := fiber.New()
	config.Connect()
	config.RunMigration()

	service := services.NewService()

	routes.Handle(app, service)
	if err := app.Listen(":" + viper.GetString("PORT")); err != nil {
		log.Fatal(err)
	}
}

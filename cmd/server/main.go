package main

import (
	"fmt"
	"log"
	"myapp/internal/config"
	"myapp/internal/database"
	"myapp/internal/model"
	"myapp/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectMySQL(cfg)
	database.DB.AutoMigrate(&model.User{})

	app := fiber.New()
	router.SetupRoutes(app)

	port := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("🚀 Server running on port %s", port)
	log.Fatal(app.Listen(port))
}

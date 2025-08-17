package api

import (
	"go-ecommerce/config"

	"github.com/gofiber/fiber/v2"
)

func StartServer(cfg config.AppConfig) {
	app := fiber.New()
	app.Listen(cfg.HostAddress)
}

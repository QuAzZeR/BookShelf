package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func NewHTTPServer() {
	app := fiber.New(fiber.Config{})
	// TODO: need to update this to use env instead
	log.Fatal(app.Listen("0.0.0.0:8000"))
}

package configuration

import (
	"github.com/donnyirianto/go-be-fiber/exception"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "Test App v1.0.1",
		ErrorHandler: exception.ErrorHandler,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	}
}

package server

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func NewRouter() *fiber.App {
	app := fiber.New(
		fiber.Config{
			JSONEncoder: jsoniter.ConfigCompatibleWithStandardLibrary.Marshal,
			JSONDecoder: jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal,
		},
	)

	return app
}

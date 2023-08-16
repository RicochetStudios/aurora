package presenter

import (
	"ricochet/aurora/config"

	"github.com/gofiber/fiber/v2"
)

// SetupSuccessResponse is the SuccessResponse that will be passed in the response by handler.
func SetupSuccessResponse(data config.Config) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// SetupErrorResponse is the singular ErrorResponse that will be passed in the response by handler.
func SetupErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

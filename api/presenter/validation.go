package presenter

import (
	"github.com/RicochetStudios/aurora/config"

	"github.com/gofiber/fiber/v2"
)

// VaildationSuccessResponse is the SuccessResponse that will be passed in the response by handler.
func VaildationSuccessResponse(data config.Config) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// VaildationErrorResponse is the singular ErrorResponse that will be passed in the response by handler.
func VaildationErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

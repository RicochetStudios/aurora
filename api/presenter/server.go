package presenter

import (
	"ricochet/aurora/types"

	"github.com/gofiber/fiber/v2"
)

// ServerSuccessResponse is the SuccessResponse that will be passed in the response by handler.
func ServerSuccessResponse(data *types.Server) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// ServerErrorResponse is the singular ErrorResponse that will be passed in the response by handler.
func ServerErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

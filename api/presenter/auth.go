package presenter

import (
	"github.com/gofiber/fiber/v2"
)

// SetupErrorResponse is the singular ErrorResponse that will be passed in the response by handler.
func AuthErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": 403,
		"data":   "forbidden Auth error ",
		"error":  err.Error(),
	}
}

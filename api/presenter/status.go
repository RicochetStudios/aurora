package presenter

import (
	"github.com/gofiber/fiber/v2"
)

type Status struct {
	Type    string `json:"type" yaml:"type" xml:"type" form:"type"`             // The type of status e.g. "healthy".
	Message string `json:"message" yaml:"message" xml:"message" form:"message"` // The message returned by the api.
}

// SetupSuccessResponse is the SuccessResponse that will be passed in the response by handler.
func StatusSuccessResponse(data Status) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// SetupErrorResponse is the singular ErrorResponse that will be passed in the response by handler.
func StatusErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

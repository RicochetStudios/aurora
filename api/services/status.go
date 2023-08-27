package services

import (
	"time"

	"github.com/RicochetStudios/aurora/api/presenter"

	"github.com/gofiber/fiber/v2"
)

// HitTest checks if the api server is running.
func HitTest() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var status presenter.Status = presenter.Status{
			Type:    "healthy",
			Message: string("Everything seems to be working, time is " + time.Now().Format("2006-01-02 15:04:05")),
		}

		return ctx.JSON(presenter.StatusSuccessResponse(status))
	}
}

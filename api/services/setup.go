package services

import (
	"fmt"
	"net/http"
	"ricochet/aurora/api/presenter"
	"ricochet/aurora/config"

	"github.com/gofiber/fiber/v2"
)

// Setup performs initialisation steps to prepare the app to take following instructions.
func Setup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var newConfig config.Config

		// Check for errors in body.
		if err := ctx.BodyParser(&newConfig); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.SetupErrorResponse(fmt.Errorf("error in provided body: \n%v", err)))
		}

		// Add or update the instance ID and cluster ID in the config.
		newConfig, err := config.Update(newConfig)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.SetupErrorResponse(fmt.Errorf("error updating local config: \n%v", err)))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.SetupSuccessResponse(newConfig))
	}
}

package services

import (
	"fmt"
	"net/http"

	"github.com/RicochetStudios/aurora/api/presenter"
	"github.com/RicochetStudios/aurora/config"

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

		// Add or update the cluster ID in the config.
		newConfig, err := config.Update(config.Config{ClusterID: newConfig.ClusterID})
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.SetupErrorResponse(fmt.Errorf("error updating local config: \n%v", err)))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.SetupSuccessResponse(newConfig))
	}
}

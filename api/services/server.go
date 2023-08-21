package services

import (
	"fmt"
	"net/http"

	"github.com/RicochetStudios/aurora/api/presenter"
	"github.com/RicochetStudios/aurora/config"
	"github.com/RicochetStudios/aurora/db"
	"github.com/RicochetStudios/aurora/docker"
	"github.com/RicochetStudios/aurora/schema"
	"github.com/RicochetStudios/aurora/types"

	"github.com/gofiber/fiber/v2"
)

// GetServer gets details about the currently configured game server instance.
func GetServer() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Read the config file.
		config, err := config.Read()
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ServerErrorResponse(fmt.Errorf("error reading config: \n%v", err)))
		}

		// Read the current server configuration.
		server, err := db.GetServer(ctx.Context(), config.ID)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ServerErrorResponse(fmt.Errorf("error reading server details from the database: \n%v", err)))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.ServerSuccessResponse(&server))
	}
}

// UpdateServer creates or updates a server.
func UpdateServer() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var server types.Server

		// Check for errors in body.
		if err := ctx.BodyParser(&server); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ServerErrorResponse(fmt.Errorf("error in provided body: \n%v", err)))
		}

		// Get the game schema.
		schema, err := schema.GetSchema("minecraft_java")
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ServerErrorResponse(fmt.Errorf("error reading schema: \n%v", err)))
		}

		// Create a container config.
		containerConfig, err := docker.NewContainerConfigFromSchema("my-unique-id", schema)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ServerErrorResponse(fmt.Errorf("error creating container config: \n%v", err)))
		}

		// Deploy and start the container.
		if _, err := docker.RunServer(ctx.Context(), containerConfig); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ServerErrorResponse(fmt.Errorf("error deploying container: \n%v", err)))
		}

		// Read the config file.
		id, err := config.GetId()
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ServerErrorResponse(fmt.Errorf("error getting config id: \n%v", err)))
		}

		// Add the server status.
		server.Status = "running"

		// Create or update the current server configuration.
		server, err = db.SetServer(ctx.Context(), id, server)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ServerErrorResponse(fmt.Errorf("error updating server details in the database: \n%v", err)))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.ServerSuccessResponse(&server))
	}
}

// RemoveServer stops and deletes a server.
func RemoveServer() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Remove the container.
		if err := docker.RemoveServer(ctx.Context()); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ServerErrorResponse(fmt.Errorf("error removing container: \n%v", err)))
		}

		// Read the config file.
		id, err := config.GetId()
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ServerErrorResponse(fmt.Errorf("error getting config id: \n%v", err)))
		}

		// Delete the current server configuration.
		_, err = db.SetServer(ctx.Context(), id, types.Server{Status: "deallocated"})
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ServerErrorResponse(fmt.Errorf("error updating server details in the database: \n%v", err)))
		}

		// Return success if the server is deleted.
		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.ServerSuccessResponse(&types.Server{}))
	}
}

package routes

import (
	"github.com/RicochetStudios/aurora/api/models"
	"github.com/RicochetStudios/aurora/api/services"

	"github.com/gofiber/fiber/v2"
)

// AuthRouter is the router for all auth methods.
func AuthRouter(app fiber.Router) {
	// set user auth.
	app.Post("/setAuthUser", func(c *fiber.Ctx) error {

		// get uid from body.
		response := new(models.UserUid)

		if err := c.BodyParser(response); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"error": err.Error(),
			})
		}
	
		// get uid from body.
		uid := response.Uid

		added, err := services.SetAuthUser(c, uid)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": true,
			"error": nil,
			"userAdded": added,
		})
	})


	// Verify user jwt token.
	app.Post("/verifyAuthUser", func(c *fiber.Ctx) error {

		// get token from body.
		response := new(models.UserToken)

		if err := c.BodyParser(response); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"error": err.Error(),
			})
		}

		// get token from body.
		token := response.Token

		decoded, err := services.CheckMemeber(c, token)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": true,
			"error": nil,
			"isMember": decoded,
		})

	})
}

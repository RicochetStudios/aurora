package routes

import (
	"log"
	"time"

	"ricochet/aurora/api/services"
	"ricochet/aurora/config"

	"github.com/gofiber/fiber/v2"
)

// METHOD: GET
// ROUTE: /
// DESC: Hit test to check if the server is running
func HitTest(c *fiber.Ctx) error {
	return c.SendString("Everything seems to be working, time is " + time.Now().Format("2006-01-02 15:04:05"))
}

// METHOD: POST
// ROUTE: /setup
// DESC: Performs initialisation steps to prepare the app to take following instructions.
func Setup(c *fiber.Ctx) error {
	var config config.Config

	// Check for errors in body.
	if err := c.BodyParser(&config); err != nil {
		log.Fatalf("Error in provided body: \n%v", err)
	}

	config = services.Setup(config)

	return c.JSON(config)
}

// TESTING
// METHOD: POST
// ROUTE: /server/firebase
// DESC: Update data into firebase
func UpdateServerFromFirebase(c *fiber.Ctx) error {
	services.UpdateServerFromFirebase(c)
	return c.SendString("Updated data into firebase")
}

// TESTING
// METHOD: GET
// ROUTE: /server/firebase
// DESC: Get data from firebase
func GetServerFromFirebase(c *fiber.Ctx) error {
	response := services.GetServerFromFirebase(c)

	return c.JSON(response)
}

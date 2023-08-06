package routes

import (
	"time"

	"ricochet/aurora/api/services"

	"github.com/gofiber/fiber/v2"
)

// METHOD: GET
// ROUTE: /
// DESC: Hit test to check if the server is running
func HitTest(c *fiber.Ctx) error {	
	return c.SendString("Everything seems to be working, time is " + time.Now().Format("2006-01-02 15:04:05"))
}

// METHOD: POST
// ROUTE: /server
// DESC: Get server details
func GetServer(c *fiber.Ctx) error {
	server, _ := services.GetServer()
	return c.JSON(server)
}

// METHOD: PUT
// ROUTE: /server
// DESC: Update server details
func UpdateServer(c *fiber.Ctx) error {
	server := services.NewServer()

	// Check for errors in body.
	if err := c.BodyParser(server); err != nil {
		return err
	}

	err := services.UpdateServer(c, server)
	if err != nil {
		return err
	}

	return c.JSON(server)
}

// TESTING
// METHOD: POST
// ROUTE: /server/firebase
// DESC: Update data into firebase
func UpdateServerFromFirebase(c *fiber.Ctx) error {
	services.UpdateServerFromFirebase()
	return c.SendString("Updated server from firebase")
}



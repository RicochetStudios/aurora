package routes

import (
	"log"
	"time"

	"ricochet/aurora/api/services"
	"ricochet/aurora/config"
	"ricochet/aurora/types"

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

// METHOD: POST
// ROUTE: /server
// DESC: Get server details
func GetServer(c *fiber.Ctx) error {
	server := services.GetServer(c)
	return c.JSON(server)
}

// METHOD: PUT
// ROUTE: /server
// DESC: Update server details
func UpdateServer(c *fiber.Ctx) error {
	var server types.Server

	// Check for errors in body.
	if err := c.BodyParser(&server); err != nil {
		log.Fatalf("Error in provided body: \n%v", err)
	}

	server = services.UpdateServer(c, server)

	return c.JSON(server)
}

// METHOD: DELETE
// ROUTE: /server
// DESC: Delete server
func RemoveServer(c *fiber.Ctx) error {
	services.RemoveServer(c)

	// Return success if the server is deleted.
	return c.JSON(struct {
		Status string `json:"status"`
	}{Status: "success"})
}

// TESTING
// METHOD: POST
// ROUTE: /server/firebase
// DESC: Update data into firebase
func UpdateServerFromFirebase(c *fiber.Ctx) error {
	services.UpdateServerFromFirebase()
	return c.SendString("Updated data into firebase")
}

// TESTING
// METHOD: GET
// ROUTE: /server/firebase
// DESC: Get data from firebase
func GetServerFromFirebase(c *fiber.Ctx) error {
	response := services.GetServerFromFirebase()

	return c.JSON(response)
}

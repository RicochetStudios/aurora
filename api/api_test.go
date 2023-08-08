package api_test

import (
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"ricochet/aurora/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/utils"
)

// go test -run -v Test_Handler
func Test_HitTest(t *testing.T) {
	app := fiber.New()
	app.Use(cors.New())

	// set routes
	app.Get("/", routes.HitTest)

	// set request
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")

	// call api
	resp, _ := app.Test(req)

	// Test 1 : if status code is 200 (API is working)
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")

	if resp.StatusCode == fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		// Test 2 if body return current time (API is working)
		utils.AssertEqual(t, "Everything seems to be working, time is "+time.Now().Format("2006-01-02 15:04:05"), string(body), "Body")
	}
}

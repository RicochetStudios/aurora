package api_test

import (
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/RicochetStudios/aurora/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"bytes"

	"github.com/gofiber/utils"
)

// TestHitTest calls HitTest with a GET method,
// checking for a valid status in response.
func TestHitTest(t *testing.T) {
	app := fiber.New()
	app.Use(cors.New())

	// set routes
	api := app.Group("/api")
	routes.StatusRouter(api)

	// set request
	req := httptest.NewRequest("GET", "/api", nil)
	req.Header.Set("Content-Type", "application/json")

	// call api
	resp, _ := app.Test(req)

	// Test 1 : if status code is 200 (API is working)
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")

	if resp.StatusCode == fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		// Test 2 if body return current time (API is working)
		var want string = `{"data":{"type":"healthy","message":"Everything seems to be working, time is ` +
			time.Now().Format("2006-01-02 15:04:05") + `"},"error":null,"status":true}`
		utils.AssertEqual(t, want, string(body), "Body")
	}
}


// SetupTest calls Setup with a POST method,
// checking for a valid status in response.
func TestSetupTest(t *testing.T) {
	app := fiber.New()
	app.Use(cors.New())

	// set routes
	api := app.Group("/api")
	routes.SetupRouter(api)

	// Setup request body
	reqBody := `{
		"clusterId": "00000001"
	}`
	bodyJson := []byte(reqBody)

	// set request
	req := httptest.NewRequest("POST", "/api/setup", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")

	// call api
	resp, _ := app.Test(req)

	// print response
	// fmt.Println(resp)

	// Test 1 : if status code is 200 (API is working)
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")

	if resp.StatusCode == fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		// Test 2 if body return current time (API is working)
		var want string = `{"data":{"id":"","clusterId":"00000001"},"error":null,"status":true}`

		utils.AssertEqual(t, want, string(body), "Body")
	}
}
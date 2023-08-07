package api_test

import (
	"net/http/httptest"
	"testing"

	"ricochet/aurora/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/utils"
)

// go test -run -v Test_Handler
func Test_HitTest(t *testing.T) {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	api.Get("/", routes.HitTest)

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 404, resp.StatusCode, "Status code")
}



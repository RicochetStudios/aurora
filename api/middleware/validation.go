package middleware

import (
	"fmt"

	"reflect"

	"github.com/RicochetStudios/aurora/api/presenter"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


var Validator = validator.New()

// ValidatePayload will validate the payload of the request
func ValidatePayload(bodyType interface{}, ctx *fiber.Ctx) error {
    
    // create a new instance of the body type
    body := reflect.New(reflect.TypeOf(bodyType)).Interface()
    ctx.BodyParser(&body)

    // validate the body
    err := Validator.Struct(body)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(presenter.StatusErrorResponse(fmt.Errorf("invalid payload")))
    }
    return ctx.Next()
}

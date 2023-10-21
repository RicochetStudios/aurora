package middleware

import (
	"fmt"

	"github.com/RicochetStudios/aurora/api/presenter"
	"github.com/RicochetStudios/aurora/db"
	"github.com/gofiber/fiber/v2"
)

// Auth funtion will check the user claims from jwt token
func Auth(ctx *fiber.Ctx) error {

	// Get an auth client from the firebase.App
	client, err := db.FirebaseAuth(ctx.Context())
	if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.StatusErrorResponse(fmt.Errorf("firebase auth error")))
    }

    // check if bear token is empty
    if(ctx.Get("Authorization") == "") {
        return ctx.Status(fiber.StatusBadRequest).JSON(presenter.StatusErrorResponse(fmt.Errorf("authorization header is empty")))
    }

	// get token from from bearer header.
	token := ctx.Get("Authorization")[7:]

	// verify token
	decoded, err := client.VerifyIDToken(ctx.Context(), token)
	if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(presenter.StatusErrorResponse(fmt.Errorf("token verification error")))
    }

	// get user claims from token
	claims := decoded.Claims

	// check if user has member claim
	if claims["member"] == true {
        return ctx.Next()
    } else {
        return ctx.Status(fiber.StatusUnauthorized).JSON(presenter.StatusErrorResponse(fmt.Errorf("user doesnt not have the correct role")))
	}
}

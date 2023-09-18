package middleware

import (
	"fmt"

	"github.com/RicochetStudios/aurora/db"
	"github.com/gofiber/fiber/v2"
)

// ProtectRoute will check the user claims from jwt token
func ProtectRoute(ctx *fiber.Ctx) error {

	// Get an auth client from the firebase.App
	client, err := db.FirebaseAuth(ctx.Context())
	if err != nil {
		return err
	}

	// get token from from bearer header.
	token := ctx.Get("Authorization")[7:]

	// verify token
	decoded, err := client.VerifyIDToken(ctx.Context(), token)
	if err != nil {
		return err
	}

	// get user claims from token
	claims := decoded.Claims

	// check if user has member claim
	if claims["member"] == true {
		return nil
	} else {
		return fmt.Errorf("user is not member")
	}
}

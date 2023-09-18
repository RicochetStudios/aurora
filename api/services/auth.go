package services

import (
	"log"

	"github.com/RicochetStudios/aurora/db"

	"github.com/gofiber/fiber/v2"
)

// SetAuthUser will give the user access to create/delete servers
func SetAuthUser(ctx *fiber.Ctx, uid string) (bool, error) {

	// Get an auth client from the firebase.App
	client, err := db.FirebaseAuth(ctx.Context())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return false, err
	}

	// Set admin privilege on the user corresponding to uid.
	claims := map[string]interface{}{"member": true}
	
	err = client.SetCustomUserClaims(ctx.Context(), uid, claims)
	if err != nil {
		log.Fatalf("error setting custom claims %v\n", err)
		return false, err
	}

	// Verify the change
	user, err := client.GetUser(ctx.Context(), uid)
	if err != nil {
		log.Fatalf("error getting user %v\n", err)
		return false, err
	}

	log.Printf("Claims after setting custom claims: %v\n", user.CustomClaims)

	return true, nil
}



// CheckMemeber will check the user claims from jwt token
func CheckMemeber(ctx *fiber.Ctx, token string) (bool, error) {

	// Get an auth client from the firebase.App
	client, err := db.FirebaseAuth(ctx.Context())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return false, err
	}

	decoded, err := client.VerifyIDToken(ctx.Context(), token)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
		return false, err
	}

	// see user claims
	claims := decoded.Claims

	// check if user has admin claim
	if claims["member"] == false {
		log.Printf("User is a memeber")
		return true, nil
	} else {
		log.Printf("User is not member")
		return false, nil
	}
}

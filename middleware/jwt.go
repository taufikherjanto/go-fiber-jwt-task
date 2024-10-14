package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTAuthorization(c *fiber.Ctx) error {
	/*
		return jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
			ContextKey: "jwt",
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				// Failed authentication return status 401
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":   true,
					"message": err.Error(),
				})
			},
		})(c)
	*/

	// simple version readable
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey: "jwt",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Failed authentication return status 401
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":   true,
					"message": err.Error(),
				})
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		},
	}

	return jwtware.New(config)(c)

}

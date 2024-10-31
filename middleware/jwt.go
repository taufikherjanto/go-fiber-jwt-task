package middleware

import (
	"strings"

	"go-fiber-jwt-task/utils"

	"github.com/gofiber/fiber/v2"
)

// JWTAuthorization middleware memeriksa keabsahan token JWT.
func JWTAuthorization(c *fiber.Ctx) error {
	// Mengambil header Authorization dari permintaan.
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "message": "JWT hilang atau tidak valid"})
	}

	// Memastikan format token benar (Bearer token).
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "message": "Format Authorization tidak valid"})
	}

	// Memverifikasi token dan mengekstrak klaim.
	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "message": "Token tidak valid atau telah kedaluwarsa"})
	}

	// Menyimpan klaim dalam konteks untuk digunakan di handler selanjutnya.
	c.Locals("jwt", claims)
	return c.Next() // Melanjutkan ke handler berikutnya jika token valid.
}

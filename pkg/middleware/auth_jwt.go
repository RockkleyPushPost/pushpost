package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"pushpost/pkg/jwt"
	"strings"
)

func AuthJWTMiddleware(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header missing",
			})
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization header format",
			})
		}
		token := parts[1]
		claims, err := jwt.VerifyToken(token, secret)
		if err != nil {
			//log.Printf("Token verification failed: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		rawUUID, ok := claims["userUUID"].(string)
		if !ok {
			log.Println("User UUID not found in token claims")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		userUUID, err := uuid.Parse(rawUUID)
		if err != nil {
			log.Printf("Invalid UUID format: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		//log.Printf("User UUID from token: %s", userUUID)

		c.Locals("userUUID", userUUID)

		return c.Next()
	}
}

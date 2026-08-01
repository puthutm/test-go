package middlewares

import (
	"data-referensi/handlers"
	"data-referensi/pkg/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	tokenString := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
	if tokenString == "" {
		return handlers.SendFailed(c, fiber.StatusUnauthorized, nil, "authentication failed")
	}

	userClaims, err := auth.VerifyTokenSpesifik(tokenString)
	if err != nil {
		return handlers.SendFailed(c, fiber.StatusUnauthorized, nil, "authentication failed")
	}

	auth.Mu.Lock()
	auth.UserClaimsGlobal = userClaims
	auth.Mu.Unlock()

	return c.Next()
}

func ApiKey(c *fiber.Ctx) error {
	// header := c.Get("x-api-key")
	return c.Next()
}

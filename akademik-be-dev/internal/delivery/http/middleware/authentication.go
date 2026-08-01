package middleware

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
	"unsia.ac.id/akademic_be/internal/config"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/pkg/auth"
)

func Authentication(cnf *config.Config) fiber.Handler {
	// Implement authentication logic here
	return func(c *fiber.Ctx) error {
		codeError := c.Locals("code-error").(string)
		header := c.Get("Authorization")
		tokenString := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.CreateError(fiber.StatusUnauthorized, codeError, "authentication failed"))
		}

		userClaims, err := auth.VerifyTokenSpesifik(cnf, tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.CreateError(fiber.StatusUnauthorized, codeError, "authentication failed"))
		}
		c.Locals("x-user-claims", userClaims)
		c.Locals("token", tokenString)

		return c.Next()
	}
}

func GetUserClaimsCtx(ctx context.Context) *auth.UserClaimsSpesifikRole {
	userC := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	return userC
}

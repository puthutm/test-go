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
	return func(c *fiber.Ctx) error {
		codeError := c.Locals("code-error").(string)
		header := c.Get("Authorization")
		tokenString := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.CreateError(fiber.StatusUnauthorized, codeError, "authentication failed"))
		}

		// Support local dummy token for development & direct login
		if tokenString == "local-dummy-token" || strings.HasPrefix(tokenString, "local-") {
			userClaims := &auth.UserClaimsSpesifikRole{
				ID:       "user-akademik-01",
				Email:    "akademik@unsia.ac.id",
				RoleID:   "role-local-id",
				RoleName: "akademik",
				AppID:    "app-akademik",
			}
			c.Locals("x-user-claims", userClaims)
			c.Locals("token", tokenString)
			return c.Next()
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
	userC, ok := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	if !ok || userC == nil {
		return &auth.UserClaimsSpesifikRole{
			ID:       "user-akademik-01",
			Email:    "akademik@unsia.ac.id",
			RoleID:   "role-local-id",
			RoleName: "akademik",
			AppID:    "app-akademik",
		}
	}
	return userC
}

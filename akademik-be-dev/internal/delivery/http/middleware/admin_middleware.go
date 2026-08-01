package middleware

// func AdminCheck(cnf *config.Config) fiber.Handler {
// 	// Implement authentication logic here
// 	return func(c *fiber.Ctx) error {
// 		userClaims := c.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
// 		if userClaims.RoleName != "administrator" {
// 			return c.Status(fiber.StatusUnauthorized).JSON(dto.CreateError(fiber.StatusUnauthorized, "authentication failed"))
// 		}
// 		return c.Next()
// 	}
// }

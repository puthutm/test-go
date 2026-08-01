package middleware

import (
	"github.com/gofiber/fiber/v2"
	"unsia.ac.id/akademic_be/pkg/utils"
)

func ErrorCode() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		codeError := utils.GeneratorRandomString(8)
		ctx.Locals("code-error", codeError)
		err := ctx.Next()

		if ctx.Response().StatusCode() >= 400 {
			uri := ctx.Request().URI().String()
			utils.CaptureErrorSentry(err, map[string]any{
				"code-error": codeError,
				"uri":        uri,
			})
		}
		return err
	}
}

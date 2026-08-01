package config

import (
	"github.com/gofiber/fiber/v2"
	"unsia.ac.id/akademic_be/internal/dto"
)

func NewFiber(config *Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      config.AppConfig.AppName,
		ErrorHandler: NewErrorHandler(),
		Prefork:      config.Server.Prefork,
	})
	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		codeError := ctx.Locals("code-error").(string)
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(dto.CreateError(code, codeError, err.Error()))
	}
}

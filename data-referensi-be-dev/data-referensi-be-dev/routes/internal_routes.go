package routes

import (
	controllers "data-referensi/app/controllers/pmb"
	"data-referensi/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InternalRoute(app fiber.Router) {
	internal := app.Group("/internal", middlewares.ApiKey)

	pmbRouteInternal(internal)
}

func pmbRouteInternal(app fiber.Router) {
	pmb := app.Group("/pmb")

	/* Academic Period */
	academicPeriod := pmb.Group("/academic-periods")
	academicPeriod.Get("/:id", controllers.GetAcademicPeriodDetailWithSession)
}

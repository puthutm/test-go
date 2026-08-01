package routes

import (
	"data-referensi/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api", middlewares.CORSMiddleware())

	RegionRoute(api)
	BiodataRoute(api)
	EducationRoute(api)
	PositionRoute(api)
	DocumentRoute(api)
	EnrollmentRoute(api)
	StudentRoute(api)
	SelectionRoute(api)
	PMBRoute(api)
	AcademicRoute(api)
	PublicRoute(api)
	InternalRoute(api)
}

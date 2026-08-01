package routes

import (
	controllers "data-referensi/app/controllers/position"
	"data-referensi/app/middlewares"
	"data-referensi/app/requests"

	"github.com/gofiber/fiber/v2"
)

func PositionRoute(app fiber.Router) {
	position := app.Group("/position", middlewares.Auth)

	/* Academic Positions */
	academic := position.Group("academic-positions")
	academicTrash := academic.Group("trashs")
	academicTrash.Get("/", requests.ValidatePagination, controllers.GetTrashAcademicPositions)
	academicTrash.Put("/:id", controllers.RestoreAcademicPosition)

	academic.Get("/", requests.ValidatePagination, controllers.GetAcademicPositions)
	academic.Get("/search", controllers.SearchAcademicPositions)
	academic.Get("/export", controllers.ExportAcademicPositions)
	academic.Get("/:id", controllers.GetAcademicPosition)
	academic.Post("/", requests.ValidateAcademicPosition, controllers.CreateAcademicPosition)
	academic.Post("/import", controllers.ImportAcademicPositions)
	academic.Put("/:id", requests.ValidateAcademicPosition, controllers.UpdateAcademicPosition)
	academic.Delete("/:id", controllers.DeleteAcademicPosition)

	/* Functional Positions */
	functional := position.Group("functional-positions")
	functionalTrash := functional.Group("trashs")
	functionalTrash.Get("/", requests.ValidatePagination, controllers.GetTrashFunctionalPositions)
	functionalTrash.Put("/:id", controllers.RestoreFunctionalPosition)

	functional.Get("/", requests.ValidatePagination, controllers.GetFunctionalPositions)
	functional.Get("/search", controllers.SearchFunctionalPositions)
	functional.Get("/export", controllers.ExportFunctionalPositions)
	functional.Get("/:id", controllers.GetFunctionalPosition)
	functional.Post("/", requests.ValidateFunctionalPosition, controllers.CreateFunctionalPosition)
	functional.Post("/import", controllers.ImportFunctionalPositions)
	functional.Put("/:id", requests.ValidateFunctionalPosition, controllers.UpdateFunctionalPosition)
	functional.Delete("/:id", controllers.DeleteFunctionalPosition)
}

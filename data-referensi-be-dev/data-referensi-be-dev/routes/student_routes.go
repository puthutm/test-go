package routes

import (
	controllers "data-referensi/app/controllers/student"
	"data-referensi/app/middlewares"
	"data-referensi/app/requests"

	"github.com/gofiber/fiber/v2"
)

func StudentRoute(app fiber.Router) {
	student := app.Group("/student", middlewares.Auth)

	/* Student Type */
	studentStatus := student.Group("student-statuses")
	studentStatusTrash := studentStatus.Group("trashs")
	studentStatusTrash.Get("/", requests.ValidatePagination, controllers.GetTrashStudentStatuses)
	studentStatusTrash.Put("/:id", controllers.RestoreStudentStatus)

	studentStatus.Get("/", requests.ValidatePagination, controllers.GetStudentStatuses)
	studentStatus.Get("/search", controllers.SearchStudentStatuses)
	studentStatus.Get("/export", controllers.ExportStudentStatuses)
	studentStatus.Get("/:id", controllers.GetStudentStatus)
	studentStatus.Post("/", requests.ValidateStudentStatus, controllers.CreateStudentStatus)
	studentStatus.Post("/import", controllers.ImportStudentStatuses)
	studentStatus.Put("/:id", requests.ValidateStudentStatus, controllers.UpdateStudentStatus)
	studentStatus.Delete("/:id", controllers.DeleteStudentStatus)

	/* Type Of Stays*/
	typeOfStay := student.Group("type-of-stays")
	typeOfStayTrash := typeOfStay.Group("trashs")
	typeOfStayTrash.Get("/", requests.ValidatePagination, controllers.GetTrashTypeOfStays)
	typeOfStayTrash.Put("/:id", controllers.RestoreTypeOfStay)

	typeOfStay.Get("/", requests.ValidatePagination, controllers.GetTypeOfStays)
	typeOfStay.Get("/search", controllers.SearchTypeOfStays)
	typeOfStay.Get("/export", controllers.ExportTypeOfStays)
	typeOfStay.Get("/:id", controllers.GetTypeOfStay)
	typeOfStay.Post("/", requests.ValidateTypeOfStay, controllers.CreateTypeOfStay)
	typeOfStay.Post("/import", controllers.ImportTypeOfStays)
	typeOfStay.Put("/:id", requests.ValidateTypeOfStay, controllers.UpdateTypeOfStay)
	typeOfStay.Delete("/:id", controllers.DeleteTypeOfStay)
}

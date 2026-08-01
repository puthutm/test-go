package routes

import (
	controllers "data-referensi/app/controllers/pmb"
	"data-referensi/app/middlewares"
	"data-referensi/app/requests"

	"github.com/gofiber/fiber/v2"
)

func PMBRoute(app fiber.Router) {
	pmb := app.Group("/pmb", middlewares.Auth)

	/* Academic Year */
	academicYear := pmb.Group("academic-years")
	academicYearTrash := academicYear.Group("trashs")
	academicYearTrash.Get("/", requests.ValidatePagination, controllers.GetTrashAcademicYears)
	academicYearTrash.Put("/:id", controllers.RestoreAcademicYear)

	academicYear.Get("/", requests.ValidatePagination, controllers.GetAcademicYears)
	academicYear.Get("/search", controllers.SearchAcademicYears)
	academicYear.Get("/export", controllers.ExportAcademicYears)
	academicYear.Get("/:id", controllers.GetAcademicYear)
	academicYear.Post("/", requests.ValidateAcademicYear, controllers.CreateAcademicYear)
	academicYear.Post("/import", controllers.ImportAcademicYears)
	academicYear.Put("/:id", requests.ValidateAcademicYear, controllers.UpdateAcademicYear)
	academicYear.Delete("/:id", controllers.DeleteAcademicYear)

	/* Semester */
	semester := pmb.Group("semesters")
	semesterTrash := semester.Group("trashs")
	semesterTrash.Get("/", requests.ValidatePagination, controllers.GetTrashSemesters)
	semesterTrash.Put("/:id", controllers.RestoreSemester)

	semester.Get("/", requests.ValidatePagination, controllers.GetSemesters)
	semester.Get("/search", controllers.SearchSemesters)
	semester.Get("/export", controllers.ExportSemesters)
	semester.Get("/:id", controllers.GetSemester)
	semester.Post("/", requests.ValidateSemester, controllers.CreateSemester)
	semester.Post("/import", controllers.ImportSemesters)
	semester.Put("/:id", requests.ValidateSemester, controllers.UpdateSemester)
	semester.Delete("/:id", controllers.DeleteSemester)

	/* Semester Number */
	semesterNumber := pmb.Group("semester-numbers")
	semesterNumberTrash := semesterNumber.Group("trashs")
	semesterNumberTrash.Get("/", requests.ValidatePagination, controllers.GetTrashSemesterNumbers)
	semesterNumberTrash.Put("/:id", controllers.RestoreSemesterNumber)

	semesterNumber.Get("/", requests.ValidatePagination, controllers.GetSemesterNumbers)
	semesterNumber.Get("/search", controllers.SearchSemesterNumbers)
	semesterNumber.Get("/export", controllers.ExportSemesterNumbers)
	semesterNumber.Get("/:id", controllers.GetSemesterNumber)
	semesterNumber.Post("/", requests.ValidateSemesterNumber, controllers.CreateSemesterNumber)
	semesterNumber.Post("/import", controllers.ImportSemesterNumbers)
	semesterNumber.Put("/:id", requests.ValidateSemesterNumber, controllers.UpdateSemesterNumber)
	semesterNumber.Delete("/:id", controllers.DeleteSemesterNumber)

	/* Academic Period */
	academicPeriod := pmb.Group("academic-periods")
	academicPeriodTrash := academicPeriod.Group("trashs")
	academicPeriodTrash.Get("/", requests.ValidatePagination, controllers.GetTrashAcademicPeriods)
	academicPeriodTrash.Put("/:id", controllers.RestoreAcademicPeriod)

	academicPeriod.Get("/", requests.ValidatePagination, controllers.GetAcademicPeriods)
	academicPeriod.Get("/search", controllers.SearchAcademicPeriods)
	academicPeriod.Get("/export", controllers.ExportAcademicPeriods)
	academicPeriod.Get("/:id", controllers.GetAcademicPeriod)
	academicPeriod.Post("/", requests.ValidateAcademicPeriod, controllers.CreateAcademicPeriod)
	academicPeriod.Post("/import", controllers.ImportAcademicPeriods)
	academicPeriod.Put("/:id", requests.ValidateAcademicPeriod, controllers.UpdateAcademicPeriod)
	academicPeriod.Delete("/:id", controllers.DeleteAcademicPeriod)

	/* School Type */
	schoolType := pmb.Group("school-types")
	schoolTypeTrash := schoolType.Group("trashs")
	schoolTypeTrash.Get("/", requests.ValidatePagination, controllers.GetTrashSchoolTypes)
	schoolTypeTrash.Put("/:id", controllers.RestoreSchoolType)

	schoolType.Get("/", requests.ValidatePagination, controllers.GetSchoolTypes)
	schoolType.Get("/search", controllers.SearchSchoolTypes)
	schoolType.Get("/export", controllers.ExportSchoolTypes)
	schoolType.Get("/:id", controllers.GetSchoolType)
	schoolType.Post("/", requests.ValidateSchoolType, controllers.CreateSchoolType)
	schoolType.Post("/import", controllers.ImportSchoolTypes)
	schoolType.Put("/:id", requests.ValidateSchoolType, controllers.UpdateSchoolType)
	schoolType.Delete("/:id", controllers.DeleteSchoolType)
}

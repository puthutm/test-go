package routes

import (
	controllers "data-referensi/app/controllers/selection"
	"data-referensi/app/middlewares"
	"data-referensi/app/requests"

	"github.com/gofiber/fiber/v2"
)

func SelectionRoute(app fiber.Router) {
	selection := app.Group("/selection", middlewares.Auth)

	/* Composition */
	composition := selection.Group("compositions")
	compositionTrash := composition.Group("trashs")
	compositionTrash.Get("/", requests.ValidatePagination, controllers.GetTrashCompositions)
	compositionTrash.Put("/:id", controllers.RestoreComposition)

	composition.Get("/", requests.ValidatePagination, controllers.GetCompositions)
	composition.Get("/search", controllers.SearchCompositions)
	composition.Get("/export", controllers.ExportCompositions)
	composition.Get("/:id", controllers.GetComposition)
	composition.Post("/", requests.ValidateComposition, controllers.CreateComposition)
	composition.Post("/import", controllers.ImportCompositions)
	composition.Put("/:id", requests.ValidateComposition, controllers.UpdateComposition)
	composition.Delete("/:id", controllers.DeleteComposition)

	/* Type Of Conditions */
	typeOfCondition := selection.Group("type-of-conditions")
	typeOfConditionTrash := typeOfCondition.Group("trashs")
	typeOfConditionTrash.Get("/", requests.ValidatePagination, controllers.GetTrashTypeOfConditions)
	typeOfConditionTrash.Put("/:id", controllers.RestoreTypeOfCondition)

	typeOfCondition.Get("/", requests.ValidatePagination, controllers.GetTypeOfConditions)
	typeOfCondition.Get("/search", controllers.SearchTypeOfConditions)
	typeOfCondition.Get("/export", controllers.ExportTypeOfConditions)
	typeOfCondition.Get("/:id", controllers.GetTypeOfCondition)
	typeOfCondition.Post("/", requests.ValidateTypeOfCondition, controllers.CreateTypeOfCondition)
	typeOfCondition.Post("/import", controllers.ImportTypeOfConditions)
	typeOfCondition.Put("/:id", requests.ValidateTypeOfCondition, controllers.UpdateTypeOfCondition)
	typeOfCondition.Delete("/:id", controllers.DeleteTypeOfCondition)

	/* Conditions */
	condition := selection.Group("conditions")
	conditionTrash := condition.Group("trashs")
	conditionTrash.Get("/", requests.ValidatePagination, controllers.GetTrashConditions)
	conditionTrash.Put("/:id", controllers.RestoreCondition)

	condition.Get("/", requests.ValidatePagination, controllers.GetConditions)
	condition.Get("/search", controllers.SearchConditions)
	condition.Get("/export", controllers.ExportConditions)
	condition.Get("/:id", controllers.GetCondition)
	condition.Post("/", requests.ValidateCondition, controllers.CreateCondition)
	condition.Post("/import", controllers.ImportConditions)
	condition.Put("/:id", requests.ValidateCondition, controllers.UpdateCondition)
	condition.Delete("/:id", controllers.DeleteCondition)

	/* Subjects */
	subject := selection.Group("subjects")
	subjectTrash := subject.Group("trashs")
	subjectTrash.Get("/", requests.ValidatePagination, controllers.GetTrashSubjects)
	subjectTrash.Put("/:id", controllers.RestoreSubject)

	subject.Get("/", requests.ValidatePagination, controllers.GetSubjects)
	subject.Get("/search", controllers.SearchSubjects)
	subject.Get("/export", controllers.ExportSubjects)
	subject.Get("/:id", controllers.GetSubject)
	subject.Post("/", requests.ValidateSubject, controllers.CreateSubject)
	subject.Post("/import", controllers.ImportSubjects)
	subject.Put("/:id", requests.ValidateSubject, controllers.UpdateSubject)
	subject.Delete("/:id", controllers.DeleteSubject)

	/* Report Card Assessments */
	reportCardAssessment := selection.Group("report-card-assessments")
	reportCardAssessmentTrash := reportCardAssessment.Group("trashs")
	reportCardAssessmentTrash.Get("/", requests.ValidatePagination, controllers.GetTrashReportCardAssessments)
	reportCardAssessmentTrash.Put("/:id", controllers.RestoreReportCardAssessment)

	reportCardAssessment.Get("/", requests.ValidatePagination, controllers.GetReportCardAssessments)
	reportCardAssessment.Get("/search", controllers.SearchReportCardAssessments)
	reportCardAssessment.Get("/export", controllers.ExportReportCardAssessments)
	reportCardAssessment.Get("/:id", controllers.GetReportCardAssessment)
	reportCardAssessment.Post("/", requests.ValidateReportCardAssessment, controllers.CreateReportCardAssessment)
	reportCardAssessment.Post("/import", controllers.ImportReportCardAssessments)
	reportCardAssessment.Put("/:id", requests.ValidateReportCardAssessment, controllers.UpdateReportCardAssessment)
	reportCardAssessment.Delete("/:id", controllers.DeleteReportCardAssessment)

}

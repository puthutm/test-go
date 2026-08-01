package routes

import (
	controllers "data-referensi/app/controllers/document"
	"data-referensi/app/middlewares"
	"data-referensi/app/requests"

	"github.com/gofiber/fiber/v2"
)

func DocumentRoute(app fiber.Router) {
	document := app.Group("/document", middlewares.Auth)

	/* Document Type */
	documentBatch := document.Group("document-types")
	documentBatchTrash := documentBatch.Group("trashs")
	documentBatchTrash.Get("/", requests.ValidatePagination, controllers.GetTrashDocumentTypes)
	documentBatchTrash.Put("/:id", controllers.RestoreDocumentType)

	documentBatch.Get("/", requests.ValidatePagination, controllers.GetDocumentTypes)
	documentBatch.Get("/search", controllers.SearchDocumentTypes)
	documentBatch.Get("/export", controllers.ExportDocumentTypes)
	documentBatch.Get("/:id", controllers.GetDocumentType)
	documentBatch.Post("/", requests.ValidateDocumentType, controllers.CreateDocumentType)
	documentBatch.Post("/import", controllers.ImportDocumentTypes)
	documentBatch.Put("/:id", requests.ValidateDocumentType, controllers.UpdateDocumentType)
	documentBatch.Delete("/:id", controllers.DeleteDocumentType)
}

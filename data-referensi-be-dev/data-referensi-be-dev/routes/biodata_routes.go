package routes

import (
	controllers "data-referensi/app/controllers/biodata"
	"data-referensi/app/middlewares"
	"data-referensi/app/requests"

	"github.com/gofiber/fiber/v2"
)

func BiodataRoute(app fiber.Router) {
	biodata := app.Group("/biodata", middlewares.Auth)

	/* Religions */
	religion := biodata.Group("religions")
	religionTrash := religion.Group("trashs")
	religionTrash.Get("/", requests.ValidatePagination, controllers.GetTrashReligions)
	religionTrash.Put("/:id", controllers.RestoreReligion)

	religion.Get("/", requests.ValidatePagination, controllers.GetReligions)
	religion.Get("/export", controllers.ExportReligions)
	religion.Get("/search", controllers.SearchReligions)
	religion.Get("/:id", controllers.GetReligion)
	religion.Post("/", requests.ValidateReligion, controllers.CreateReligion)
	religion.Post("/import", controllers.ImportReligions)
	religion.Put("/:id", requests.ValidateReligion, controllers.UpdateReligion)
	religion.Delete("/:id", controllers.DeleteReligion)

	/* Jobs */
	job := biodata.Group("jobs")
	jobTrash := job.Group("trashs")
	jobTrash.Get("/", requests.ValidatePagination, controllers.GetTrashJobs)
	jobTrash.Put("/:id", controllers.RestoreJob)

	job.Get("/", requests.ValidatePagination, controllers.GetJobs)
	job.Get("/export", controllers.ExportJobs)
	job.Get("/search", controllers.SearchJobs)
	job.Get("/:id", controllers.GetJob)
	job.Post("/", requests.ValidateJob, controllers.CreateJob)
	job.Post("/import", controllers.ImportJobs)
	job.Put("/:id", requests.ValidateJob, controllers.UpdateJob)
	job.Delete("/:id", controllers.DeleteJob)

	/* Earnings */
	earning := biodata.Group("earnings")
	earningTrash := earning.Group("trashs")
	earningTrash.Get("/", requests.ValidatePagination, controllers.GetTrashEarnings)
	earningTrash.Put("/:id", controllers.RestoreEarning)

	earning.Get("/", requests.ValidatePagination, controllers.GetEarnings)
	earning.Get("/export", controllers.ExportEarnings)
	earning.Get("/search", controllers.SearchEarnings)
	earning.Get("/:id", controllers.GetEarning)
	earning.Post("/", requests.ValidateEarning, controllers.CreateEarning)
	earning.Post("/import", controllers.ImportEarnings)
	earning.Put("/:id", requests.ValidateEarning, controllers.UpdateEarning)
	earning.Delete("/:id", controllers.DeleteEarning)

	/* Ethnics */
	ethnic := biodata.Group("ethnics")
	ethnicTrash := ethnic.Group("trashs")
	ethnicTrash.Get("/", requests.ValidatePagination, controllers.GetTrashEthnics)
	ethnicTrash.Put("/:id", controllers.RestoreEthnic)

	ethnic.Get("/", requests.ValidatePagination, controllers.GetEthnics)
	ethnic.Get("/export", controllers.ExportEthnics)
	ethnic.Get("/search", controllers.SearchEthnics)
	ethnic.Get("/:id", controllers.GetEthnic)
	ethnic.Post("/", requests.ValidateEthnic, controllers.CreateEthnic)
	ethnic.Post("/import", controllers.ImportEthnics)
	ethnic.Put("/:id", requests.ValidateEthnic, controllers.UpdateEthnic)
	ethnic.Delete("/:id", controllers.DeleteEthnic)

	/* Almamater Sizes */
	almamaterSize := biodata.Group("almamater-sizes")
	almamaterSizeTrash := almamaterSize.Group("trashs")
	almamaterSizeTrash.Get("/", requests.ValidatePagination, controllers.GetTrashAlmamaterSizes)
	almamaterSizeTrash.Put("/:id", controllers.RestoreAlmamaterSize)

	almamaterSize.Get("/", requests.ValidatePagination, controllers.GetAlmamaterSizes)
	almamaterSize.Get("/export", controllers.ExportAlmamaterSizes)
	almamaterSize.Get("/search", controllers.SearchAlmamaterSizes)
	almamaterSize.Get("/:id", controllers.GetAlmamaterSize)
	almamaterSize.Post("/", requests.ValidateAlmamaterSize, controllers.CreateAlmamaterSize)
	almamaterSize.Post("/import", controllers.ImportAlmamaterSizes)
	almamaterSize.Put("/:id", requests.ValidateAlmamaterSize, controllers.UpdateAlmamaterSize)
	almamaterSize.Delete("/:id", controllers.DeleteAlmamaterSize)

	/* Marriage Statuses */
	marriageStatus := biodata.Group("marriage-statuses")
	marriageStatusTrash := marriageStatus.Group("trashs")
	marriageStatusTrash.Get("/", requests.ValidatePagination, controllers.GetTrashMarriageStatuses)
	marriageStatusTrash.Put("/:id", controllers.RestoreMarriageStatus)

	marriageStatus.Get("/", requests.ValidatePagination, controllers.GetMarriageStatuses)
	marriageStatus.Get("/export", controllers.ExportMarriageStatuses)
	marriageStatus.Get("/search", controllers.SearchMarriageStatuses)
	marriageStatus.Get("/:id", controllers.GetMarriageStatus)
	marriageStatus.Post("/", requests.ValidateMarriageStatus, controllers.CreateMarriageStatus)
	marriageStatus.Post("/import", controllers.ImportMarriageStatuses)
	marriageStatus.Put("/:id", requests.ValidateMarriageStatus, controllers.UpdateMarriageStatus)
	marriageStatus.Delete("/:id", controllers.DeleteMarriageStatus)

	/* Banks */
	bank := biodata.Group("banks")
	bankTrash := bank.Group("trashs")
	bankTrash.Get("/", requests.ValidatePagination, controllers.GetTrashBanks)
	bankTrash.Put("/:id", controllers.RestoreBank)

	bank.Get("/", requests.ValidatePagination, controllers.GetBanks)
	bank.Get("/export", controllers.ExportBanks)
	bank.Get("/search", controllers.SearchBanks)
	bank.Get("/:id", controllers.GetBank)
	bank.Post("/", requests.ValidateBank, controllers.CreateBank)
	bank.Post("/import", controllers.ImportBanks)
	bank.Put("/:id", requests.ValidateBank, controllers.UpdateBank)
	bank.Delete("/:id", controllers.DeleteBank)

	/* Work Units */
	workUnit := biodata.Group("work-units")
	workUnitTrash := workUnit.Group("trashs")
	workUnitTrash.Get("/", requests.ValidatePagination, controllers.GetTrashWorkUnits)
	workUnitTrash.Put("/:id", controllers.RestoreWorkUnit)

	workUnit.Get("/", requests.ValidatePagination, controllers.GetWorkUnits)
	workUnit.Get("/export", controllers.ExportWorkUnits)
	workUnit.Get("/search", controllers.SearchWorkUnits)
	workUnit.Get("/:id", controllers.GetWorkUnit)
	workUnit.Post("/", requests.ValidateWorkUnit, controllers.CreateWorkUnit)
	workUnit.Post("/import", controllers.ImportWorkUnits)
	workUnit.Put("/:id", requests.ValidateWorkUnit, controllers.UpdateWorkUnit)
	workUnit.Delete("/:id", controllers.DeleteWorkUnit)

	/* Transportation */
	transportation := biodata.Group("transportations")
	transportationTrash := transportation.Group("trashs")
	transportationTrash.Get("/", requests.ValidatePagination, controllers.GetTrashTransportations)
	transportationTrash.Put("/:id", controllers.RestoreTransportation)

	transportation.Get("/", requests.ValidatePagination, controllers.GetTransportations)
	transportation.Get("/export", controllers.ExportTransportations)
	transportation.Get("/search", controllers.SearchTransportations)
	transportation.Get("/:id", controllers.GetTransportation)
	transportation.Post("/", requests.ValidateTransportation, controllers.CreateTransportation)
	transportation.Post("/import", controllers.ImportTransportations)
	transportation.Put("/:id", requests.ValidateTransportation, controllers.UpdateTransportation)
	transportation.Delete("/:id", controllers.DeleteTransportation)

	/* Employee Type */
	employeeType := biodata.Group("employee-types")
	employeeTypeTrash := employeeType.Group("trashs")
	employeeTypeTrash.Get("/", requests.ValidatePagination, controllers.GetTrashEmployeeTypes)
	employeeTypeTrash.Put("/:id", controllers.RestoreEmployeeType)

	employeeType.Get("/", requests.ValidatePagination, controllers.GetEmployeeTypes)
	employeeType.Get("/export", controllers.ExportEmployeeTypes)
	employeeType.Get("/search", controllers.SearchEmployeeTypes)
	employeeType.Get("/:id", controllers.GetEmployeeType)
	employeeType.Post("/", requests.ValidateEmployeeType, controllers.CreateEmployeeType)
	employeeType.Post("/import", controllers.ImportEmployeeTypes)
	employeeType.Put("/:id", requests.ValidateEmployeeType, controllers.UpdateEmployeeType)
	employeeType.Delete("/:id", controllers.DeleteEmployeeType)

	/* Employee Status */
	employeeStatus := biodata.Group("employee-statuses")
	employeeStatusTrash := employeeStatus.Group("trashs")
	employeeStatusTrash.Get("/", requests.ValidatePagination, controllers.GetTrashEmployeeStatuses)
	employeeStatusTrash.Put("/:id", controllers.RestoreEmployeeStatus)

	employeeStatus.Get("/", requests.ValidatePagination, controllers.GetEmployeeStatuses)
	employeeStatus.Get("/export", controllers.ExportEmployeeStatuses)
	employeeStatus.Get("/search", controllers.SearchEmployeeStatuses)
	employeeStatus.Get("/:id", controllers.GetEmployeeStatus)
	employeeStatus.Post("/", requests.ValidateEmployeeStatus, controllers.CreateEmployeeStatus)
	employeeStatus.Post("/import", controllers.ImportEmployeeStatuses)
	employeeStatus.Put("/:id", requests.ValidateEmployeeStatus, controllers.UpdateEmployeeStatus)
	employeeStatus.Delete("/:id", controllers.DeleteEmployeeStatus)

	/* Blood Type */
	bloodType := biodata.Group("blood-types")
	bloodTypeTrash := bloodType.Group("trashs")
	bloodTypeTrash.Get("/", requests.ValidatePagination, controllers.GetTrashBloodTypes)
	bloodTypeTrash.Put("/:id", controllers.RestoreBloodType)

	bloodType.Get("/", requests.ValidatePagination, controllers.GetBloodTypes)
	bloodType.Get("/export", controllers.ExportBloodTypes)
	bloodType.Get("/search", controllers.SearchBloodTypes)
	bloodType.Get("/:id", controllers.GetBloodType)
	bloodType.Post("/", requests.ValidateBloodType, controllers.CreateBloodType)
	bloodType.Post("/import", controllers.ImportBloodTypes)
	bloodType.Put("/:id", requests.ValidateBloodType, controllers.UpdateBloodType)
	bloodType.Delete("/:id", controllers.DeleteBloodType)

	/* Working Relationship */
	workingRelationship := biodata.Group("working-relationships")
	workingRelationshipTrash := workingRelationship.Group("trashs")
	workingRelationshipTrash.Get("/", requests.ValidatePagination, controllers.GetTrashWorkingRelationships)
	workingRelationshipTrash.Put("/:id", controllers.RestoreWorkingRelationship)

	workingRelationship.Get("/", requests.ValidatePagination, controllers.GetWorkingRelationships)
	workingRelationship.Get("/export", controllers.ExportWorkingRelationships)
	workingRelationship.Get("/search", controllers.SearchWorkingRelationships)
	workingRelationship.Get("/:id", controllers.GetWorkingRelationship)
	workingRelationship.Post("/", requests.ValidateWorkingRelationship, controllers.CreateWorkingRelationship)
	workingRelationship.Post("/import", controllers.ImportWorkingRelationships)
	workingRelationship.Put("/:id", requests.ValidateWorkingRelationship, controllers.UpdateWorkingRelationship)
	workingRelationship.Delete("/:id", controllers.DeleteWorkingRelationship)
}

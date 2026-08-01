package routes

import (
	controllers "data-referensi/app/controllers/enrollment"
	"data-referensi/app/middlewares"
	"data-referensi/app/requests"

	"github.com/gofiber/fiber/v2"
)

func EnrollmentRoute(app fiber.Router) {
	enrollment := app.Group("/enrollment", middlewares.Auth)

	/* Enrollment Batch */
	enrollmentBatch := enrollment.Group("enrollment-batchs")
	enrollmentBatchTrash := enrollmentBatch.Group("trashs")
	enrollmentBatchTrash.Get("/", requests.ValidatePagination, controllers.GetTrashEnrollmentBatchs)
	enrollmentBatchTrash.Put("/:id", controllers.RestoreEnrollmentBatch)

	enrollmentBatch.Get("/", requests.ValidatePagination, controllers.GetEnrollmentBatchs)
	enrollmentBatch.Get("/search", controllers.SearchEnrollmentBatchs)
	enrollmentBatch.Get("/export", controllers.ExportEnrollmentBatchs)
	enrollmentBatch.Get("/:id", controllers.GetEnrollmentBatch)
	enrollmentBatch.Post("/", requests.ValidateEnrollmentBatch, controllers.CreateEnrollmentBatch)
	enrollmentBatch.Post("/import", controllers.ImportEnrollmentBatchs)
	enrollmentBatch.Put("/:id", requests.ValidateEnrollmentBatch, controllers.UpdateEnrollmentBatch)
	enrollmentBatch.Delete("/:id", controllers.DeleteEnrollmentBatch)

	/* Lecture System */
	lectureSystem := enrollment.Group("lecture-systems")
	lectureSystemTrash := lectureSystem.Group("trashs")
	lectureSystemTrash.Get("/", requests.ValidatePagination, controllers.GetTrashLectureSystems)
	lectureSystemTrash.Put("/:id", controllers.RestoreLectureSystem)

	lectureSystem.Get("/", requests.ValidatePagination, controllers.GetLectureSystems)
	lectureSystem.Get("/search", controllers.SearchLectureSystems)
	lectureSystem.Get("/export", controllers.ExportLectureSystems)
	lectureSystem.Get("/:id", controllers.GetLectureSystem)
	lectureSystem.Post("/", requests.ValidateLectureSystem, controllers.CreateLectureSystem)
	lectureSystem.Post("/import", controllers.ImportLectureSystems)
	lectureSystem.Put("/:id", requests.ValidateLectureSystem, controllers.UpdateLectureSystem)
	lectureSystem.Delete("/:id", controllers.DeleteLectureSystem)

	/* Program Type */
	programType := enrollment.Group("program-types")
	programTypeTrash := programType.Group("trashs")
	programTypeTrash.Get("/", requests.ValidatePagination, controllers.GetTrashProgramTypes)
	programTypeTrash.Put("/:id", controllers.RestoreProgramType)

	programType.Get("/", requests.ValidatePagination, controllers.GetProgramTypes)
	programType.Get("/search", controllers.SearchProgramTypes)
	programType.Get("/export", controllers.ExportProgramTypes)
	programType.Get("/:id", controllers.GetProgramType)
	programType.Post("/", requests.ValidateProgramType, controllers.CreateProgramType)
	programType.Post("/import", controllers.ImportProgramTypes)
	programType.Put("/:id", requests.ValidateProgramType, controllers.UpdateProgramType)
	programType.Delete("/:id", controllers.DeleteProgramType)

	/* Registration Path */
	registrationPath := enrollment.Group("registration-paths")
	registrationPathTrash := registrationPath.Group("trashs")
	registrationPathTrash.Get("/", requests.ValidatePagination, controllers.GetTrashRegistrationPaths)
	registrationPathTrash.Put("/:id", controllers.RestoreRegistrationPath)

	registrationPath.Get("/", requests.ValidatePagination, controllers.GetRegistrationPaths)
	registrationPath.Get("/search", controllers.SearchRegistrationPaths)
	registrationPath.Get("/export", controllers.ExportRegistrationPaths)
	registrationPath.Get("/:id", controllers.GetRegistrationPath)
	registrationPath.Post("/", requests.ValidateRegistrationPath, controllers.CreateRegistrationPath)
	registrationPath.Post("/import", controllers.ImportRegistrationPaths)
	registrationPath.Put("/:id", requests.ValidateRegistrationPath, controllers.UpdateRegistrationPath)
	registrationPath.Delete("/:id", controllers.DeleteRegistrationPath)

	/* Registration Status */
	registrationStatus := enrollment.Group("registration-statuses")
	registrationStatusTrash := registrationStatus.Group("trashs")
	registrationStatusTrash.Get("/", requests.ValidatePagination, controllers.GetTrashRegistrationStatuses)
	registrationStatusTrash.Put("/:id", controllers.RestoreRegistrationStatus)

	registrationStatus.Get("/", requests.ValidatePagination, controllers.GetRegistrationStatuses)
	registrationStatus.Get("/search", controllers.SearchRegistrationStatuses)
	registrationStatus.Get("/export", controllers.ExportRegistrationStatuses)
	registrationStatus.Get("/:id", controllers.GetRegistrationStatus)
	registrationStatus.Post("/", requests.ValidateRegistrationStatus, controllers.CreateRegistrationStatus)
	registrationStatus.Post("/import", controllers.ImportRegistrationStatuses)
	registrationStatus.Put("/:id", requests.ValidateRegistrationStatus, controllers.UpdateRegistrationStatus)
	registrationStatus.Delete("/:id", controllers.DeleteRegistrationStatus)

	/* Referral */
	referral := enrollment.Group("referrals")
	referralTrash := referral.Group("trashs")
	referralTrash.Get("/", requests.ValidatePagination, controllers.GetTrashReferrals)
	referralTrash.Put("/:id", controllers.RestoreReferral)

	referral.Get("/", requests.ValidatePagination, controllers.GetReferrals)
	referral.Get("/search", controllers.SearchReferrals)
	referral.Get("/export", controllers.ExportReferrals)
	referral.Get("/:id", controllers.GetReferral)
	referral.Post("/", requests.ValidateReferral, controllers.CreateReferral)
	referral.Post("/import", controllers.ImportReferrals)
	referral.Put("/:id", requests.ValidateReferral, controllers.UpdateReferral)
	referral.Delete("/:id", controllers.DeleteReferral)
}

package router

import "github.com/gofiber/fiber/v2"

func (r *RouterConfig) SetupStudentRoute(api fiber.Router) {
	/* Student */
	student := api.Group("/student")

	// Home
	r.SetupStudentHomeRoute(student)
	// Student
	r.SetupStudentStudentRoute(student)
	// Academic
	r.SetupStudentAcademicRoute(student)
	// Finance
	r.SetupStudentFinanceRoute(student)
	// Course
	r.SetupStudentCourseRoute(student)
}

func (r *RouterConfig) SetupStudentHomeRoute(student fiber.Router) {
	// home := student.Group("/home")
}

func (r *RouterConfig) SetupStudentStudentRoute(student fiber.Router) {
	// Biodata
	r.SetupStudentStudentBiodataRoute(student)

	student = student.Group("/student")

	// Academic
	r.SetupStudentStudentAcademicRoute(student)
	// Final Project
	r.SetupStudentStudentFinalProjectRoute(student)
	// Internship
	r.SetupStudentStudentInternshipRoute(student)
	// Administration
	r.SetupStudentStudentAdministrationRoute(student)
	// Academic Calendar
	r.SetupStudentStudentAcademicCalendarRoute(student)
	// KTP
	r.SetupStudentStudentKTPRoute(student)
}

func (r *RouterConfig) SetupStudentStudentBiodataRoute(student fiber.Router) {
	biodata := student.Group("/biodata")
	biodataBiodata := biodata.Group("/biodatas")
	biodataBiodata.Get("/", r.MstStudentBioController.GetBioGeneralOnlyUser)
	biodataBiodata.Put("/", r.MstStudentBioController.UpdateBioGeneralOnlyUser)

	// information
	biodataInformation := biodata.Group("/informations")
	biodataInformation.Put("/",
		r.MstStudentBioController.UpdateBioInformationOnlyUser,
	)
	biodataInformation.Get(
		"/",
		r.MstStudentBioController.GetBioInformationOnlyUser,
	)

	// completeness
	biodataCompleteness := biodata.Group("/completeness")
	biodataCompleteness.Put(
		"/",
		r.MstStudentBioController.UpdateBioCompletenessOnlyUser,
	)
	biodataCompleteness.Get(
		"/",
		r.MstStudentBioController.GetBioCompletenessOnlyUser,
	)

	// Address
	biodataAddress := biodata.Group("/addresses")
	biodataAddress.Get(
		"/",
		r.MstStudentDomicileController.GetByStudentID,
	)
	biodataAddress.Put(
		"/",
		r.MstStudentDomicileController.UpdateByStudentID,
	)

	// Parent / Family
	biodataParent := biodata.Group("/parents")
	biodataParent.Get(
		"/:parent_type",
		r.MstStudentFamilyController.GetByStudentID,
	)
	biodataParent.Put(
		"/:parent_type",
		r.MstStudentFamilyController.UpdateByStudentID,
	)

	// Bank Account
	biodataBankAccount := biodata.Group("/bank-accounts")
	biodataBankAccount.Put(
		"/",
		r.MstStudentBioController.UpdateBioBankAccountOnlyUser,
	)
	biodataBankAccount.Get(
		"/",
		r.MstStudentBioController.GetBioBankAccountOnlyUser,
	)

	// Document
	document := biodata.Group("/documents")
	document.Put(
		"/",
		r.MstStudentBioController.UpdateBioDocumentOnlyUser,
	)
	document.Get(
		"/",
		r.MstStudentBioController.GetBioDocumentOnlyUser,
	)

	// Original Education
	biodataOriginalEducation := biodata.Group("/original-educations")
	biodataOriginalEducation.Get(
		"/",
		r.MstStudentEducationController.GetByStudentID,
	)
	biodataOriginalEducation.Put(
		"/",
		r.MstStudentEducationController.UpdateByID,
	)

	// study program
	studyPorgram := biodata.Group("/study-program")
	studyPorgram.Get(
		"/", r.MstStudentStudyProgramController.GetAllWithCountSearchByStudyProgram,
	)
}

func (r *RouterConfig) SetupStudentStudentAcademicRoute(student fiber.Router) {
}

func (r *RouterConfig) SetupStudentStudentFinalProjectRoute(student fiber.Router) {
	finalProjectProposal := student.Group("/final-project-proposals")

	finalProjectProposal.Get("/", r.TrxFinalProjectProposalController.GetAllForStudent)
	finalProjectProposal.Post("/", r.TrxFinalProjectProposalController.Create)
	finalProjectProposal.Get("/:proposal_id", r.TrxFinalProjectProposalController.GetByIDForStudent)
}

func (r *RouterConfig) SetupStudentStudentInternshipRoute(student fiber.Router) {
}

func (r *RouterConfig) SetupStudentStudentAdministrationRoute(student fiber.Router) {
}

func (r *RouterConfig) SetupStudentStudentAcademicCalendarRoute(student fiber.Router) {
}

func (r *RouterConfig) SetupStudentStudentKTPRoute(student fiber.Router) {
}

func (r *RouterConfig) SetupStudentAcademicRoute(student fiber.Router) {
	academic := student.Group("/academic")

	fillingKRS := academic.Group("/filling-krs")
	fillingKRS.Get("/pick", r.KrsController.GetPickClassesByUserID)
	fillingKRS.Post("/pick/take", r.KrsController.TakeClass)

	fillingKRS.Get("/saved", r.KrsController.GetSavedByUserID)
	fillingKRS.Delete("/saved/:krs_item_id", r.KrsController.DeleteSavedByKrsItemID)

	fillingKRS.Get("/info", r.KrsController.GetKrsMaxSksInfo)

	academic.Get("/khs", r.KhsController.GetKHS)
}

func (r *RouterConfig) SetupStudentFinanceRoute(student fiber.Router) {
	// finance := student.Group("/finance")
}

func (r *RouterConfig) SetupStudentCourseRoute(student fiber.Router) {
	// course := student.Group("/course")
}

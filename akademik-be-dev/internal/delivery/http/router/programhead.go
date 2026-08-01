package router

import "github.com/gofiber/fiber/v2"

func (r *RouterConfig) SetupProgramHeadRoute(api fiber.Router) {
	/* Program Head */
	programHead := api.Group("/program-head")

	// Home
	r.SetupHomeRoute(programHead)

	// Portal
	r.SetupPortalRoute(programHead)

	// Curriculum
	r.SetupCurriculumRoute(programHead)

	// Course
	r.SetupCourseRoute(programHead)

	// Lectures
	r.SetupLecturesRoute(programHead)
}

func (r *RouterConfig) SetupHomeRoute(programHead fiber.Router) {
}

func (r *RouterConfig) SetupPortalRoute(programHead fiber.Router) {
	// Portal
	portal := programHead.Group("/portal")

	/* Student */
	student := portal.Group("/students")
	student.Get("/", r.MstStudentStudyProgramController.GetAllWithCountByProgramHeadID)

	/* Lecturer */
	lecturer := portal.Group("/lecturers")
	lecturer.Get("/", r.MstLecturerController.GetAllWithCountByProgramHeadID)
}

func (r *RouterConfig) SetupCurriculumRoute(programHead fiber.Router) {
	// Curriculum
	curriculum := programHead.Group("/curriculum")

	/* Academic Period */
	academicPeriod := curriculum.Group("/academic-period")

	/* Class */
	class := academicPeriod.Group("/classes")

	class.Post("/", r.MstClassController.CreateByProgramHead)
	class.Put("/:class_id", r.MstClassController.UpdateByIDAndProgramHead)
	class.Get("/", r.MstClassController.GetAllWithCountByProgramHeadID)
	class.Get("/:class_id", r.MstClassController.GetByID)
	class.Delete("/:class_id", r.MstClassController.DeleteByIDAndProgramHead)

	classContract := class.Group("/:class_id/contract")
	classContract.Put("/", r.MstClassController.UpdateContractByIDAndProgramHead)

	class = academicPeriod.Group("/classes/:class_id")

	classSchedule := class.Group("/schedules")
	classSchedule.Get("/", r.MstClassScheduleController.GetByClassID)
	classSchedule.Post("/", r.MstClassScheduleController.CreateByProgramHead)
	classSchedule.Get("/generate", r.MstClassScheduleController.GenerateByProgramHead)

	classLecturer := class.Group("/lecturers")
	classLecturer.Get("/", r.MstCLassLecturerController.GetByClassID)

	classParticipant := class.Group("/participants")
	classParticipant.Get("/", r.MstCLassParticipantController.GetAllWithCount)
	classParticipant.Post("/", r.MstCLassParticipantController.CreateByProgramHead)
	classParticipant.Delete("/:participant_id", r.MstCLassParticipantController.DeleteByIDAndProgramHead)

	classScheduleTemplate := classSchedule.Group("/template")
	classScheduleTemplate.Post("", r.MstClassScheduleTemplateController.CreateByProgramHead)
	classScheduleTemplate.Put("/:id", r.MstClassScheduleTemplateController.UpdateByProgramHead)
	classScheduleTemplate.Delete("/:id", r.MstClassScheduleTemplateController.DeleteByIDProgramHead)
	classScheduleTemplate.Get("", r.MstClassScheduleTemplateController.GetByClassIDAndProgramHead)
	classScheduleTemplate.Get("/:id", r.MstClassScheduleTemplateController.GetByIDAndProgramHead)

	classScheduleOfDate := class.Group("/schedules-as-of-date")
	classScheduleOfDate.Get("/", r.MstClassScheduleController.GetByClassAsDate)

	/* Study Program Curriculum */
	studyProgramCurriculum := academicPeriod.Group("/study-program-curriculums")

	studyProgramCurriculumTrash := academicPeriod.Group("/study-program-curriculums/trash")
	studyProgramCurriculumTrash.Get("/", r.MstStudyProgramCurriculumController.GetByStudyProgramIDAndSemesterIDAndProgramHeadTrash)
	studyProgramCurriculumTrash.Put("/:study_program_curriculum_id", r.MstStudyProgramCurriculumController.RestoreByIDAndProgramHead)

	studyProgramCurriculum.Get("/", r.MstStudyProgramCurriculumController.GetByStudyProgramIDAndSemesterIDAndProgramHead)
	studyProgramCurriculum.Post("/", r.MstStudyProgramCurriculumController.CreateByProgramHead)
	studyProgramCurriculum.Put("/package", r.MstStudyProgramCurriculumController.UpdateBlastPackageBySemesterWithoutProgramStudy)
	studyProgramCurriculum.Get("/:study_program_curriculum_id", r.MstStudyProgramCurriculumController.GetByIDAndProgramHead)
	studyProgramCurriculum.Put("/:study_program_curriculum_id", r.MstStudyProgramCurriculumController.UpdateByIDAndProgramHead)
	studyProgramCurriculum.Delete("/:study_program_curriculum_id", r.MstStudyProgramCurriculumController.DeleteByIDAndProgramHead)
	studyProgramCurriculum.Get("/search/subject", r.MstStudyProgramCurriculumController.GetByStudyProgramAndSemesterAndCuricullumForSubjectDataProgramHead)
}

func (r *RouterConfig) SetupCourseRoute(programHead fiber.Router) {
	course := programHead.Group("/course")

	// Final Project Proposal
	finalProjectProposal := course.Group("/final-project-proposals")
	finalProjectProposal.Get("/", r.TrxFinalProjectProposalController.GetAllWithCountByProgramHeadID)
	finalProjectProposal.Get("/by-student/:student_id/study-program/:study_program_id", r.TrxFinalProjectProposalController.GetByStudentIDandStudyProgramID)
	finalProjectProposal.Get("/:user_id/group-by-student", r.TrxFinalProjectProposalController.GetProposalStudentByUser)
	finalProjectProposal.Get("/:proposal_id", r.TrxFinalProjectProposalController.GetByID)
	finalProjectProposal.Put("/:proposal_id/status", r.TrxFinalProjectProposalController.UpdateByID)
	finalProjectProposal.Post("/:proposal_id/assign-academic-supervisor", r.TrxFinalProjectProposalController.AsignAcademicSupervisor)
}

func (r *RouterConfig) SetupLecturesRoute(programHead fiber.Router) {
	// Lectures
	lectures := programHead.Group("/lectures")

	// KRS Requests
	lectures.Get("/krs-requests", r.KrsController.GetKrsProgramHeadClasses)
}

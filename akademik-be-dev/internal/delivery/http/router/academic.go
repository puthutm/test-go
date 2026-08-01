package router

import "github.com/gofiber/fiber/v2"

func (r *RouterConfig) SetupAcademicRoute(api fiber.Router) {
	/* Academic */
	academic := api.Group("/academic")

	// Setting
	r.SetupSettingRoute(academic)

	// Curriculum
	r.SetupAcademicCurriculumRoute(academic)

	// Portal
	r.SetupAcademicPortalRoute(academic)
}

func (r *RouterConfig) SetupSettingRoute(academic fiber.Router) {
	// Setting
	setting := academic.Group("/setting")

	/* SKS Limit */
	sksLimit := setting.Group("/sks-limits")

	sksLimitTrash := sksLimit.Group("trash")
	sksLimitTrash.Get("/", r.MstSKSLimitController.GetAllTrashWithCount)
	sksLimitTrash.Put("/:sks_limit_id", r.MstSKSLimitController.RestoreByID)

	sksLimit.Get("/", r.MstSKSLimitController.GetAllAndCount)
	sksLimit.Post("/", r.MstSKSLimitController.Create)
	sksLimit.Get("/:sks_limit_id", r.MstSKSLimitController.GetByID)
	sksLimit.Put("/:sks_limit_id", r.MstSKSLimitController.UpdateByID)
	sksLimit.Delete("/:sks_limit_id", r.MstSKSLimitController.DeleteByID)

	// Value scale
	valueScale := setting.Group("/value-scales")
	valueScaleTrash := valueScale.Group("trash")
	valueScaleTrash.Get("/", r.MstValueScaleController.GetAllTrashWithCount)
	valueScaleTrash.Put("/:value_scale_id", r.MstValueScaleController.RestoreByID)

	valueScale.Get("/", r.MstValueScaleController.GetAllWithCount)
	valueScale.Post("/", r.MstValueScaleController.Create)
	valueScale.Get("/:value_scale_id", r.MstValueScaleController.GetByID)
	valueScale.Put("/:value_scale_id", r.MstValueScaleController.UpdateByID)
	valueScale.Delete("/:value_scale_id", r.MstValueScaleController.DeleteByID)

	// Value scale
	valueComposition := setting.Group("/value-compositions")
	valueCompositionTrash := valueComposition.Group("trash")
	valueCompositionTrash.Get("/", r.MstValueCompositionController.GetAllTrashWithCount)
	valueCompositionTrash.Put("/:value_composition_id", r.MstValueCompositionController.RestoreByID)

	valueComposition.Get("/", r.MstValueCompositionController.GetAllWithCount)
	valueComposition.Get("/academic-periods", r.MstValueCompositionController.GetAcademicPeriodsWithCount)
	valueComposition.Post("/duplicate", r.MstValueCompositionController.Duplicate)
	valueComposition.Post("/", r.MstValueCompositionController.Create)
	valueComposition.Get("/:value_composition_id", r.MstValueCompositionController.GetByID)
	valueComposition.Put("/:value_composition_id", r.MstValueCompositionController.UpdateByID)
	valueComposition.Delete("/:value_composition_id", r.MstValueCompositionController.DeleteByID)

	// Subject
	subject := setting.Group("/subjects")
	subjectTrash := subject.Group("trash")
	subjectTrash.Get("/", r.MstSubjectController.GetAllTrashWithCount)
	subjectTrash.Put("/:subject_id", r.MstSubjectController.RestoreByID)

	subject.Get("/", r.MstSubjectController.GetAllWithCount)
	subject.Get("/search", r.MstSubjectController.GetByStudyProgramIDAndCurriculumYearID)
	subject.Get("/search/program-head", r.MstSubjectController.GetByProgramHeadAndCurriculumYearID)
	subject.Post("/", r.MstSubjectController.Create)
	subject.Get("/:subject_id", r.MstSubjectController.GetByID)
	subject.Put("/:subject_id", r.MstSubjectController.UpdateByID)
	subject.Delete("/:subject_id", r.MstSubjectController.DeleteByID)

	// academic period
	// Class
	academicPeriod := setting.Group("/academic-period")
	class := academicPeriod.Group("/classes")
	classTrash := class.Group("trash")
	classTrash.Get("/", r.MstClassController.GetAllTrashWithCount)
	classTrash.Put("/:class_id", r.MstClassController.RestoreByID)
	class.Get("/", r.MstClassController.GetAllWithCount)
	class.Post("/", r.MstClassController.Create)
	class.Get("/:class_id", r.MstClassController.GetByID)
	class.Put("/:class_id", r.MstClassController.UpdateByID)
	class.Delete("/:class_id", r.MstClassController.DeleteByID)

	// Open Close Values
	openCloseValues := academicPeriod.Group("/:academic_periode_id/classes/open-close-values")
	openCloseValues.Get("/", r.MstClassController.CheckSaveButton)
	openCloseValues.Put("/", r.MstClassController.UpdateStatusLockedByAcademicPeriodID)

	// Classes by Study Program
	classByStudyProgram := academicPeriod.Group("/:academic_periode_id/classes/:study_program_id")
	classByStudyProgram.Get("/", r.MstClassController.GetAllWithCountByStudyProgramID)

	// Class Score
	classScore := academicPeriod.Group("/:academic_periode_id/classes/:class_id/class-score")
	classScore.Get("/", r.ClassScoreController.GetByClassID)

	// Class Score Open Close Values
	classScoreOpenCloseValues := classScore.Group("/open-close-values")
	classScoreOpenCloseValues.Get("/", r.ClassScoreController.CheckSaveButton)
	classScoreOpenCloseValues.Put("/", r.ClassScoreController.UpdateStatusLock)

	classContract := class.Group("/:class_id/contract")
	classContract.Put("/", r.MstClassController.UpdateContractByID)

	// class participan
	classes := academicPeriod.Group("/classes/:class_id")
	classParticipant := classes.Group("/participants")
	classParticipant.Post("/", r.MstCLassParticipantController.Create)
	classParticipant.Get("/", r.MstCLassParticipantController.GetAllWithCount)
	classParticipant.Get("/:participant_id", r.MstCLassParticipantController.GetByID)
	classParticipant.Delete("/:participant_id", r.MstCLassParticipantController.DeleteByID)
	// class lecturer
	classLecturer := classes.Group("/lecturers")
	classLecturer.Post("/", r.MstCLassLecturerController.Create)
	classLecturer.Get("/", r.MstCLassLecturerController.GetByClassID)
	classLecturer.Put("/:lecturer_id", r.MstCLassLecturerController.Update)
	classLecturer.Get("/:lecturer_id", r.MstCLassLecturerController.GetByID)
	classLecturer.Delete("/:lecturer_id", r.MstCLassLecturerController.DeleteByID)

	classSchedule := classes.Group("/schedules")
	classSchedule.Post("/", r.MstClassScheduleController.Create)
	classSchedule.Get("/", r.MstClassScheduleController.GetByClassID)
	classSchedule.Get("/generate", r.MstClassScheduleController.GenerateByAcademic)

	classScheduleDayTime := classes.Group("/schedules/day-time")
	classScheduleDayTime.Put("/", r.MstClassScheduleController.UpdateByDayTime)
	classScheduleDayTime.Get("/", r.MstClassScheduleController.GetByDayTime)
	classScheduleDayTime.Delete("/", r.MstClassScheduleController.DeleteByID)

	classScheduleTemplate := classSchedule.Group("/template")
	classScheduleTemplate.Post("", r.MstClassScheduleTemplateController.Create)
	classScheduleTemplate.Put("/:id", r.MstClassScheduleTemplateController.Update)
	classScheduleTemplate.Delete("/:id", r.MstClassScheduleTemplateController.DeleteByID)
	classScheduleTemplate.Get("", r.MstClassScheduleTemplateController.GetByClassID)
	classScheduleTemplate.Get("/:id", r.MstClassScheduleTemplateController.GetByID)

	classScheduleOfDate := classes.Group("/schedules-as-of-date")
	classScheduleOfDate.Get("/", r.MstClassScheduleController.GetByClassAsDate)

	// curriculum year
	curriculumYear := setting.Group("/curriculum_year")
	curriculumYearID := curriculumYear.Group("/:curriculum_year_id")

	// subject
	subjectCurriculumYearID := curriculumYearID.Group("/subjects")
	subjectCurriculumYearID.Get("", r.MstSubjectController.GetAllWithCountForCurriculumYear)

	// Assessment Weight
	assessmentWeight := setting.Group("/assessment-weight")
	assessmentWeight.Get("/", r.MstAssessmentWeightController.GetFirst)
	assessmentWeight.Post("/", r.MstAssessmentWeightController.Create)

	// Presence
	presence := setting.Group("/presence")
	presenceStudent := presence.Group("/student")
	presenceStudent.Post("/", r.MstStudentPresenceSettingController.Create)
	presenceStudent.Post("/duplicate", r.MstStudentPresenceSettingController.Duplicate)
	presenceStudent.Get("/", r.MstStudentPresenceSettingController.GetAllWithCount)
}

func (r *RouterConfig) SetupAcademicCurriculumRoute(academic fiber.Router) {
	// Curriculum
	curriculum := academic.Group("/curriculum")

	/* Study Program Curriculum */
	studyProgramCurriculum := curriculum.Group("/study-program-curriculums")

	studyProgramCurriculumTrash := curriculum.Group("/study-program-curriculums/trash")
	studyProgramCurriculumTrash.Get("/", r.MstStudyProgramCurriculumController.GetByStudyProgramIDAndSemesterIDTrash)
	studyProgramCurriculumTrash.Put("/:study_program_curriculum_id", r.MstStudyProgramCurriculumController.RestoreByID)

	studyProgramCurriculum.Get("/", r.MstStudyProgramCurriculumController.GetByStudyProgramIDAndSemesterID)
	studyProgramCurriculum.Post("/", r.MstStudyProgramCurriculumController.Create)
	studyProgramCurriculum.Put("/package", r.MstStudyProgramCurriculumController.UpdateBlastPackageBySemesterWithProgramStudy)
	studyProgramCurriculum.Get("/:study_program_curriculum_id", r.MstStudyProgramCurriculumController.GetByID)
	studyProgramCurriculum.Put("/:study_program_curriculum_id", r.MstStudyProgramCurriculumController.UpdateByID)
	studyProgramCurriculum.Delete("/:study_program_curriculum_id", r.MstStudyProgramCurriculumController.DeleteByID)
	studyProgramCurriculum.Get("/search/subject", r.MstStudyProgramCurriculumController.GetByStudyProgramAndSemesterAndCuricullumForSubjectDataAcademic)
}

func (r *RouterConfig) SetupAcademicPortalRoute(academic fiber.Router) {
	// Portal
	portal := academic.Group("/portal")
	portalStudents := portal.Group("/students")
	portalStudents.Get("/batches", r.PortalStudentController.GetAllBatches)
	portalStudents.Get("/", r.PortalStudentController.GetStudentList)
	portalStudents.Post("/", r.PortalStudentController.CreateStudent)
	portalStudents.Post("/bulk", r.PortalStudentController.CreateStudentBulk)
}

package router

import "github.com/gofiber/fiber/v2"

func (r *RouterConfig) SetupLecturerRoute(api fiber.Router) {
	/* Lecturer */
	lecturer := api.Group("/lecturer")

	// Curriculum
	r.SetupLecturerCurriculumRoute(lecturer)

	// Academic
	r.SetupLecturerAcademicRoute(lecturer)

	// Lectures
	r.SetupLecturerLecturesRoute(lecturer)
}

func (r *RouterConfig) SetupLecturerCurriculumRoute(lecturer fiber.Router) {
	// Curriculum
	curriculum := lecturer.Group("/curriculum")

	// Subject
	subject := curriculum.Group("/subjects")
	subject.Get("/", r.MstSubjectController.GetAllWithCountByLecuturerID)
	subject.Get("/:subject_id", r.MstSubjectController.GetByID)

	// Coordinator MK
	coordinatorMK := curriculum.Group("/coordinator-mks")
	coordinatorMK.Get("/", r.MstSubjectController.GetAllWithCountByCoordinatorLecuturerID)
	coordinatorMK.Get("/:subject_id", r.MstSubjectController.GetByID)
}

func (r *RouterConfig) SetupLecturerAcademicRoute(lecturer fiber.Router) {
	// Academic
	academic := lecturer.Group("/academic")

	// Class Schedule
	classSchedule := academic.Group("/class-schedules")
	classSchedule.Get("/", r.MstClassScheduleController.GetAllWithCountByLecuturerID)
	classSchedule.Get("/:class_id", r.MstClassController.GetByID)

	classSchedule = classSchedule.Group("/:class_id")
	classSchedule.Get("/weekly-schedule", r.MstClassScheduleController.GetByClassID)
	classSchedule.Get("/distribution-of-study-programs", r.MstStudyProgramController.GetByLecturerIDandActiveAcademicPeriod)
	classSchedule.Get("/academic-system-distributions", r.MstClassScheduleController.GetByLecturerIDandActiveAcademicPeriod)
	classSchedule.Get("/student-class-distributions", r.MstClassController.GetByLecturerIDandActiveAcademicPeriod)
	// classSchedule.Get("/teaching-lecture", r.MstBiodataOfUserController.GetByUserID)
	classSchedule.Get("/class-participants", r.MstCLassParticipantController.GetAllWithCountByClassIDForLecturer)
	classSchedule.Get("/class-schedules", r.MstClassScheduleController.GetByClassAsDate)
	classSchedule.Get("/class-schedules/:class_schedule_id", r.MstClassScheduleController.GetByID)
	classSchedule.Post("/class-schedules/:class_schedule_id", r.MstClassScheduleController.UpdateByIDForLecturer)
	classSchedule.Delete("/class-schedules/:schedule_id", r.MstClassScheduleController.DeleteByID)
	classSchedule.Get("/class-attendances", r.MstClassScheduleController.GetByClassAsDate)

	courseAssisment := classSchedule.Group("/course-assisments")
	courseAssisment.Get("/", r.MstClassScheduleTaskController.GetAll)
	courseAssisment.Post("/", r.MstClassScheduleTaskController.Create)
	courseAssisment.Get("/:course_assisment_id", r.MstClassScheduleTaskController.GetByID)

	presence := academic.Group("/presence")
	presence.Get(
		"/",
		r.MstCLassLecturerController.GetSubjectByClassLecturerWithCount,
	)

	// student presence
	presenceStudent := presence.Group("/students")
	presenceStudent.Get("/sessions/:session_id", r.TrxStudentPresenceSettingController.GetStudentPresenceBySessionWithCount)
	presenceStudent.Post("/sessions/:session_id", r.TrxStudentPresenceSettingController.CreateOrUpdateStudentPresence)
	presenceStudent.Post("/sessions/:session_id/bulk", r.TrxStudentPresenceSettingController.CreateOrUpdateStudentPresenceSlice)
	presenceStudent.Get("/class_schedules/:class_schedule_id", r.MstClassScheduleController.GetByIDForPresence)
	presenceStudent.Get("/components/sessions/:session_id", r.TrxStudentPresenceSettingController.GetComponentBySession)

	academicPeriodAndSubject := presence.Group("/academic-periods/:academic_period_id/subjects/:subject_id")

	academicPeriodAndSubject.Post("/class", r.TrxStudentPresenceSettingController.CreateOrUpdate)
	academicPeriodAndSubject.Get(
		"/class",
		r.MstCLassLecturerController.GetClassByAcademicPeriodAndSubjectAndUserForLecturerWithCount,
	)
	academicPeriodAndSubject.Get(
		"/component", r.TrxStudentPresenceSettingController.GetComponentForLecturer,
	)
	academicPeriodAndSubject.Get(
		"/class/:class_id",
		r.TrxStudentPresenceSettingController.GetSessionPresenceByClassID,
	)
}

func (r *RouterConfig) SetupLecturerLecturesRoute(lecturer fiber.Router) {
	// Lectures
	lectures := lecturer.Group("/lectures")

	// KRS Requests
	lectures.Get("/krs-requests", r.KrsController.GetKrsLecturerStudents)
	lectures.Get("/krs-requests/:krsID", r.KrsController.GetKrsLecturerStudentDetailByKrsHeaderID)
	lectures.Put("/krs-requests/:krsItemID", r.KrsController.UpdateKrsItemStatusByKrsItemID)
}

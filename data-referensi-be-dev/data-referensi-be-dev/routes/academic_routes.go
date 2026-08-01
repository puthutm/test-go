package routes

import (
	controllers "data-referensi/app/controllers/academic"
	"data-referensi/app/middlewares"
	"data-referensi/app/requests"

	"github.com/gofiber/fiber/v2"
)

func AcademicRoute(app fiber.Router) {
	academic := app.Group("/academic", middlewares.Auth)

	/* Course Type */
	courseType := academic.Group("course-types")
	courseTypeTrash := courseType.Group("trashs")
	courseTypeTrash.Get("/", requests.ValidatePagination, controllers.GetTrashCourseTypes)
	courseTypeTrash.Put("/:id", controllers.RestoreCourseType)

	courseType.Get("/", requests.ValidatePagination, controllers.GetCourseTypes)
	courseType.Get("/search", controllers.SearchCourseTypes)
	courseType.Get("/export", controllers.ExportCourseTypes)
	courseType.Get("/:id", controllers.GetCourseType)
	courseType.Post("/", requests.ValidateCourseType, controllers.CreateCourseType)
	courseType.Post("/import", controllers.ImportCourseTypes)
	courseType.Put("/:id", requests.ValidateCourseType, controllers.UpdateCourseType)
	courseType.Delete("/:id", controllers.DeleteCourseType)

	/* Course Group */
	courseGroup := academic.Group("course-groups")
	courseGroupTrash := courseGroup.Group("trashs")
	courseGroupTrash.Get("/", requests.ValidatePagination, controllers.GetTrashCourseGroups)
	courseGroupTrash.Put("/:id", controllers.RestoreCourseGroup)

	courseGroup.Get("/", requests.ValidatePagination, controllers.GetCourseGroups)
	courseGroup.Get("/search", controllers.SearchCourseGroups)
	courseGroup.Get("/export", controllers.ExportCourseGroups)
	courseGroup.Get("/:id", controllers.GetCourseGroup)
	courseGroup.Post("/", requests.ValidateCourseGroup, controllers.CreateCourseGroup)
	courseGroup.Post("/import", controllers.ImportCourseGroups)
	courseGroup.Put("/:id", requests.ValidateCourseGroup, controllers.UpdateCourseGroup)
	courseGroup.Delete("/:id", controllers.DeleteCourseGroup)

	/* Value Element */
	valueElement := academic.Group("value-elements")
	valueElementTrash := valueElement.Group("trashs")
	valueElementTrash.Get("/", requests.ValidatePagination, controllers.GetTrashValueElements)
	valueElementTrash.Put("/:id", controllers.RestoreValueElement)

	valueElement.Get("/", requests.ValidatePagination, controllers.GetValueElements)
	valueElement.Get("/search", controllers.SearchValueElements)
	valueElement.Get("/export", controllers.ExportValueElements)
	valueElement.Get("/:id", controllers.GetValueElement)
	valueElement.Post("/", requests.ValidateValueElement, controllers.CreateValueElement)
	valueElement.Post("/import", controllers.ImportValueElements)
	valueElement.Put("/:id", requests.ValidateValueElement, controllers.UpdateValueElement)
	valueElement.Delete("/:id", controllers.DeleteValueElement)

	/* Lecture Status */
	lectureStatus := academic.Group("lecture-statuses")
	lectureStatusTrash := lectureStatus.Group("trashs")
	lectureStatusTrash.Get("/", requests.ValidatePagination, controllers.GetTrashLecturerStatuses)
	lectureStatusTrash.Put("/:id", controllers.RestoreLecturerStatus)

	lectureStatus.Get("/", requests.ValidatePagination, controllers.GetLecturerStatuses)
	lectureStatus.Get("/search", controllers.SearchLecturerStatuses)
	lectureStatus.Get("/export", controllers.ExportLecturerStatuses)
	lectureStatus.Get("/:id", controllers.GetLecturerStatus)
	lectureStatus.Post("/", requests.ValidateLecturerStatus, controllers.CreateLecturerStatus)
	lectureStatus.Post("/import", controllers.ImportLecturerStatuses)
	lectureStatus.Put("/:id", requests.ValidateLecturerStatus, controllers.UpdateLecturerStatus)
	lectureStatus.Delete("/:id", controllers.DeleteLecturerStatus)

	/* Special Need */
	specialNeed := academic.Group("special-needs")
	specialNeedTrash := specialNeed.Group("trashs")
	specialNeedTrash.Get("/", requests.ValidatePagination, controllers.GetTrashSpecialNeeds)
	specialNeedTrash.Put("/:id", controllers.RestoreSpecialNeed)

	specialNeed.Get("/", requests.ValidatePagination, controllers.GetSpecialNeeds)
	specialNeed.Get("/search", controllers.SearchSpecialNeeds)
	specialNeed.Get("/export", controllers.ExportSpecialNeeds)
	specialNeed.Get("/:id", controllers.GetSpecialNeed)
	specialNeed.Post("/", requests.ValidateSpecialNeed, controllers.CreateSpecialNeed)
	specialNeed.Post("/import", controllers.ImportSpecialNeeds)
	specialNeed.Put("/:id", requests.ValidateSpecialNeed, controllers.UpdateSpecialNeed)
	specialNeed.Delete("/:id", controllers.DeleteSpecialNeed)

	/* Attendance Status */
	attendanceStatus := academic.Group("attendance-statuses")
	attendanceStatusTrash := attendanceStatus.Group("trashs")
	attendanceStatusTrash.Get("/", requests.ValidatePagination, controllers.GetTrashAttendanceStatuses)
	attendanceStatusTrash.Put("/:id", controllers.RestoreAttendanceStatus)

	attendanceStatus.Get("/", requests.ValidatePagination, controllers.GetAttendanceStatuses)
	attendanceStatus.Get("/search", controllers.SearchAttendanceStatuses)
	attendanceStatus.Get("/export", controllers.ExportAttendanceStatuses)
	attendanceStatus.Get("/:id", controllers.GetAttendanceStatus)
	attendanceStatus.Post("/", requests.ValidateAttendanceStatus, controllers.CreateAttendanceStatus)
	attendanceStatus.Post("/import", controllers.ImportAttendanceStatuses)
	attendanceStatus.Put("/:id", requests.ValidateAttendanceStatus, controllers.UpdateAttendanceStatus)
	attendanceStatus.Delete("/:id", controllers.DeleteAttendanceStatus)

	/* Class Status */
	classStatus := academic.Group("class-statuses")
	classStatusTrash := classStatus.Group("trashs")
	classStatusTrash.Get("/", requests.ValidatePagination, controllers.GetTrashClassStatuses)
	classStatusTrash.Put("/:id", controllers.RestoreClassStatus)

	classStatus.Get("/", requests.ValidatePagination, controllers.GetClassStatuses)
	classStatus.Get("/search", controllers.SearchClassStatuses)
	classStatus.Get("/export", controllers.ExportClassStatuses)
	classStatus.Get("/:id", controllers.GetClassStatus)
	classStatus.Post("/", requests.ValidateClassStatus, controllers.CreateClassStatus)
	classStatus.Post("/import", controllers.ImportClassStatuses)
	classStatus.Put("/:id", requests.ValidateClassStatus, controllers.UpdateClassStatus)
	classStatus.Delete("/:id", controllers.DeleteClassStatus)

	/* Academic Activity */
	academicActivity := academic.Group("academic-activities")
	academicActivityTrash := academicActivity.Group("trashs")
	academicActivityTrash.Get("/", requests.ValidatePagination, controllers.GetTrashAcademicActivities)
	academicActivityTrash.Put("/:id", controllers.RestoreAcademicActivity)

	academicActivity.Get("/", requests.ValidatePagination, controllers.GetAcademicActivities)
	academicActivity.Get("/search", controllers.SearchAcademicActivities)
	academicActivity.Get("/export", controllers.ExportAcademicActivities)
	academicActivity.Get("/:id", controllers.GetAcademicActivity)
	academicActivity.Post("/", requests.ValidateAcademicActivity, controllers.CreateAcademicActivity)
	academicActivity.Post("/import", controllers.ImportAcademicActivities)
	academicActivity.Put("/:id", requests.ValidateAcademicActivity, controllers.UpdateAcademicActivity)
	academicActivity.Delete("/:id", controllers.DeleteAcademicActivity)

	/* SKS Limit */
	sksLimit := academic.Group("sks-limits")
	sksLimitTrash := sksLimit.Group("trashs")
	sksLimitTrash.Get("/", requests.ValidatePagination, controllers.GetTrashSksLimits)
	sksLimitTrash.Put("/:id", controllers.RestoreSksLimit)

	sksLimit.Get("/", requests.ValidatePagination, controllers.GetSksLimits)
	sksLimit.Get("/search", controllers.SearchSksLimits)
	sksLimit.Get("/export", controllers.ExportSksLimits)
	sksLimit.Get("/:id", controllers.GetSksLimit)
	sksLimit.Post("/", requests.ValidateSksLimit, controllers.CreateSksLimit)
	sksLimit.Post("/import", controllers.ImportSksLimits)
	sksLimit.Put("/:id", requests.ValidateSksLimit, controllers.UpdateSksLimit)
	sksLimit.Delete("/:id", controllers.DeleteSksLimit)

	/* Value Scale */
	valueScale := academic.Group("value-scales")
	valueScaleTrash := valueScale.Group("trashs")
	valueScaleTrash.Get("/", requests.ValidatePagination, controllers.GetTrashValueScales)
	valueScaleTrash.Put("/:id", controllers.RestoreValueScale)

	valueScale.Get("/", requests.ValidatePagination, controllers.GetValueScales)
	valueScale.Get("/search", controllers.SearchValueScales)
	valueScale.Get("/export", controllers.ExportValueScales)
	valueScale.Get("/:id", controllers.GetValueScale)
	valueScale.Post("/", requests.ValidateValueScale, controllers.CreateValueScale)
	valueScale.Post("/import", controllers.ImportValueScales)
	valueScale.Put("/:id", requests.ValidateValueScale, controllers.UpdateValueScale)
	valueScale.Delete("/:id", controllers.DeleteValueScale)

	/* SKS Weight */
	sksWeight := academic.Group("sks-weight")
	sksWeightTrash := sksWeight.Group("trashs")
	sksWeightTrash.Get("/", requests.ValidatePagination, controllers.GetTrashSksWeights)
	sksWeightTrash.Put("/:id", controllers.RestoreSksWeight)

	sksWeight.Get("/", requests.ValidatePagination, controllers.GetSksWeights)
	sksWeight.Get("/search", controllers.SearchSksWeights)
	sksWeight.Get("/export", controllers.ExportSksWeights)
	sksWeight.Get("/:id", controllers.GetSksWeight)
	sksWeight.Post("/", requests.ValidateSksWeight, controllers.CreateSksWeight)
	sksWeight.Post("/import", controllers.ImportSksWeights)
	sksWeight.Put("/:id", requests.ValidateSksWeight, controllers.UpdateSksWeight)
	sksWeight.Delete("/:id", controllers.DeleteSksWeight)

	/* Predicate */
	predicate := academic.Group("predicates")
	predicateTrash := predicate.Group("trashs")
	predicateTrash.Get("/", requests.ValidatePagination, controllers.GetTrashPredicates)
	predicateTrash.Put("/:id", controllers.RestorePredicate)

	predicate.Get("/", requests.ValidatePagination, controllers.GetPredicates)
	predicate.Get("/search", controllers.SearchPredicates)
	predicate.Get("/export", controllers.ExportPredicates)
	predicate.Get("/:id", controllers.GetPredicate)
	predicate.Post("/", requests.ValidatePredicate, controllers.CreatePredicate)
	predicate.Post("/import", controllers.ImportPredicates)
	predicate.Put("/:id", requests.ValidatePredicate, controllers.UpdatePredicate)
	predicate.Delete("/:id", controllers.DeletePredicate)

	/* Faculty */
	faculty := academic.Group("faculties")
	facultyTrash := faculty.Group("trashs")
	facultyTrash.Get("/", requests.ValidatePagination, controllers.GetTrashFaculties)
	facultyTrash.Put("/:id", controllers.RestoreFaculty)

	faculty.Get("/", requests.ValidatePagination, controllers.GetFaculties)
	faculty.Get("/search", controllers.SearchFaculties)
	faculty.Get("/export", controllers.ExportFaculties)
	faculty.Get("/:id", controllers.GetFaculty)
	faculty.Post("/", requests.ValidateFaculty, controllers.CreateFaculty)
	faculty.Post("/import", controllers.ImportFaculties)
	faculty.Put("/:id", requests.ValidateFaculty, controllers.UpdateFaculty)
	faculty.Delete("/:id", controllers.DeleteFaculty)

	/* Course */
	course := academic.Group("courses")
	courseTrash := course.Group("trashs")
	courseTrash.Get("/", requests.ValidatePagination, controllers.GetTrashCourses)
	courseTrash.Put("/:id", controllers.RestoreCourse)

	course.Get("/", requests.ValidatePagination, controllers.GetCourses)
	course.Get("/search", controllers.SearchCourses)
	course.Get("/export", controllers.ExportCourses)
	course.Get("/:id", controllers.GetCourse)
	course.Post("/", requests.ValidateCourse, controllers.CreateCourse)
	course.Post("/import", controllers.ImportCourses)
	course.Put("/:id", requests.ValidateCourse, controllers.UpdateCourse)
	course.Delete("/:id", controllers.DeleteCourse)

	/* Curriculum Year */
	curriculumYear := academic.Group("curriculum-years")
	curriculumYearTrash := curriculumYear.Group("trashs")
	curriculumYearTrash.Get("/", requests.ValidatePagination, controllers.GetTrashCurriculumYears)
	curriculumYearTrash.Put("/:id", controllers.RestoreCurriculumYear)

	curriculumYear.Get("/", requests.ValidatePagination, controllers.GetCurriculumYears)
	curriculumYear.Get("/search", controllers.SearchCurriculumYears)
	curriculumYear.Get("/export", controllers.ExportCurriculumYears)
	curriculumYear.Get("/:id", controllers.GetCurriculumYear)
	curriculumYear.Post("/", requests.ValidateCurriculumYear, controllers.CreateCurriculumYear)
	curriculumYear.Post("/import", controllers.ImportCurriculumYears)
	curriculumYear.Put("/:id", requests.ValidateCurriculumYear, controllers.UpdateCurriculumYear)
	curriculumYear.Delete("/:id", controllers.DeleteCurriculumYear)

	/* Final Assignment Stage */
	finalAssignmentStage := academic.Group("final-assignment-stages")
	finalAssignmentStageTrash := finalAssignmentStage.Group("trashs")
	finalAssignmentStageTrash.Get("/", requests.ValidatePagination, controllers.GetTrashFinalAssignmentStages)
	finalAssignmentStageTrash.Put("/:id", controllers.RestoreFinalAssignmentStage)

	finalAssignmentStage.Get("/", requests.ValidatePagination, controllers.GetFinalAssignmentStages)
	finalAssignmentStage.Get("/search", controllers.SearchFinalAssignmentStages)
	finalAssignmentStage.Get("/export", controllers.ExportFinalAssignmentStages)
	finalAssignmentStage.Get("/:id", controllers.GetFinalAssignmentStage)
	finalAssignmentStage.Post("/", requests.ValidateFinalAssignmentStage, controllers.CreateFinalAssignmentStage)
	finalAssignmentStage.Post("/import", controllers.ImportFinalAssignmentStages)
	finalAssignmentStage.Put("/:id", requests.ValidateFinalAssignmentStage, controllers.UpdateFinalAssignmentStage)
	finalAssignmentStage.Delete("/:id", controllers.DeleteFinalAssignmentStage)

	/* Student Certificate */
	studenCertificate := academic.Group("student-certificates")
	studenCertificateTrash := studenCertificate.Group("trashs")
	studenCertificateTrash.Get("/", requests.ValidatePagination, controllers.GetTrashStudentCertificates)
	studenCertificateTrash.Put("/:id", controllers.RestoreStudentCertificate)

	studenCertificate.Get("/", requests.ValidatePagination, controllers.GetStudentCertificates)
	studenCertificate.Get("/search", controllers.SearchStudentCertificates)
	studenCertificate.Get("/export", controllers.ExportStudentCertificates)
	studenCertificate.Get("/:id", controllers.GetStudentCertificate)
	studenCertificate.Post("/", requests.ValidateStudentCertificate, controllers.CreateStudentCertificate)
	studenCertificate.Post("/import", controllers.ImportStudentCertificates)
	studenCertificate.Put("/:id", requests.ValidateStudentCertificate, controllers.UpdateStudentCertificate)
	studenCertificate.Delete("/:id", controllers.DeleteStudentCertificate)

	/* Approval */
	approval := academic.Group("approvals")
	approvalTrash := approval.Group("trashs")
	approvalTrash.Get("/", requests.ValidatePagination, controllers.GetTrashApprovals)
	approvalTrash.Put("/:id", controllers.RestoreApproval)

	approval.Get("/", requests.ValidatePagination, controllers.GetApprovals)
	approval.Get("/search", controllers.SearchApprovals)
	approval.Get("/export", controllers.ExportApprovals)
	approval.Get("/:id", controllers.GetApproval)
	approval.Post("/", requests.ValidateApproval, controllers.CreateApproval)
	approval.Post("/import", controllers.ImportApprovals)
	approval.Put("/:id", requests.ValidateApproval, controllers.UpdateApproval)
	approval.Delete("/:id", controllers.DeleteApproval)

	/* Grade */
	grade := academic.Group("grades")
	gradeTrash := grade.Group("trashs")
	gradeTrash.Get("/", requests.ValidatePagination, controllers.GetTrashGrades)
	gradeTrash.Put("/:id", controllers.RestoreGrade)

	grade.Get("/", requests.ValidatePagination, controllers.GetGrades)
	grade.Get("/search", controllers.SearchGrades)
	grade.Get("/export", controllers.ExportGrades)
	grade.Get("/:id", controllers.GetGrade)
	grade.Post("/", requests.ValidateGrade, controllers.CreateGrade)
	grade.Post("/import", controllers.ImportGrades)
	grade.Put("/:id", requests.ValidateGrade, controllers.UpdateGrade)
	grade.Delete("/:id", controllers.DeleteGrade)

	/* Field Study Concentration */
	fieldStudyConcentration := academic.Group("field-study-concentrations")
	fieldStudyConcentrationTrash := fieldStudyConcentration.Group("trashs")
	fieldStudyConcentrationTrash.Get("/", requests.ValidatePagination, controllers.GetTrashFieldStudyConcentrations)
	fieldStudyConcentrationTrash.Put("/:id", controllers.RestoreFieldStudyConcentration)

	fieldStudyConcentration.Get("/", requests.ValidatePagination, controllers.GetFieldStudyConcentrations)
	fieldStudyConcentration.Get("/search", controllers.SearchFieldStudyConcentrations)
	fieldStudyConcentration.Get("/export", controllers.ExportFieldStudyConcentrations)
	fieldStudyConcentration.Get("/:id", controllers.GetFieldStudyConcentration)
	fieldStudyConcentration.Post("/", requests.ValidateFieldStudyConcentration, controllers.CreateFieldStudyConcentration)
	fieldStudyConcentration.Post("/import", controllers.ImportFieldStudyConcentrations)
	fieldStudyConcentration.Put("/:id", requests.ValidateFieldStudyConcentration, controllers.UpdateFieldStudyConcentration)
	fieldStudyConcentration.Delete("/:id", controllers.DeleteFieldStudyConcentration)
}

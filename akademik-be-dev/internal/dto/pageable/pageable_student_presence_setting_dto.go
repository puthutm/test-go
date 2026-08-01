// Package pageable
package pageable

type PageableStudentSettingGets struct {
	PageableRequest
	StudyProgramID   *string `query:"study_program_id"`
	AcademicPeriodID *string `query:"academic_periode_id"`
}

type PageableStudentPresenceBySession struct {
	PageableRequest
	Status    *string `query:"status"`
	SessionID string  `query:"-"`
}

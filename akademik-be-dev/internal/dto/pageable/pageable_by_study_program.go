package pageable

// TODO: PageableRequest for class participant
type PageableRequestByStudyProgram struct {
	PageableRequest
	StudyProgramID string `query:"study_program_id"`
}

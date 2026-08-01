package dto

import "github.com/google/uuid"

type GetStudyProgramCurriculumRequest struct {
	// Filter utama
	StudyProgramID   string `query:"study_program_id"`
	SemesterNumberID string `query:"semester_number_id"`
	CurriculumYearID string `query:"curriculum_year_id"`
	UserID           string `query:"-"`
}

type MstStudyProgramCurriculumSubjectPrerequisiteResponse struct {
	ID                              uuid.UUID `json:"id"`
	StudyProgramCurriculumID        uuid.UUID `json:"study_program_curriculum_id"`
	StudyProgramCurriculumSubjectID uuid.UUID `json:"study_program_curriculum_subject_id"`
	// Tambahkan kolom tambahan hasil JOIN (dari tabel mst_subject atau lainnya)
	SubjectNameID string `json:"subject_name_id"`
	SubjectNameEN string `json:"subject_name_en"`
	SubjectCode   string `json:"subject_code"`
}

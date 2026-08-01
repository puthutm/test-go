package model

import "github.com/google/uuid"

type MstStudyProgramCurriculumSubjectPrerequisite struct {
	ID                              uuid.UUID `gorm:"type:char(36);primaryKey"`
	StudyProgramCurriculumID        uuid.UUID `gorm:"type:char(36);not null;column:study_program_curriculum_id"`
	StudyProgramCurriculumSubjectID uuid.UUID `gorm:"type:char(36);not null;column:study_program_curriculum_subject_id"`
	// Tambahkan kolom tambahan hasil JOIN (dari tabel mst_subject atau lainnya)
	SubjectNameID string `gorm:"column:subject_name_id"`
	SubjectNameEN string `gorm:"column:subject_name_en"`
	SubjectCode   string `gorm:"column:subject_code"`
}

// TableName overrides the table name
func (MstStudyProgramCurriculumSubjectPrerequisite) TableName() string {
	return "mst_study_program_curriculum_subject_prerequisites"
}

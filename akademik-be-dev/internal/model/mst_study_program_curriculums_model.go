package model

import "github.com/google/uuid"

type MstStudyProgramCurriculum struct {
	ID                        uuid.UUID  `gorm:"type:char(36);primaryKey"`
	CurriculumYearID          uuid.UUID  `gorm:"type:char(36);column:curriculum_year_id; not null"`
	SubjectID                 uuid.UUID  `gorm:"type:char(36);column:subject_id; not null"`
	SemesterNumberID          uuid.UUID  `gorm:"type:char(36);column:semester_number_id; not null"`
	LimitGradeID              *string    `gorm:"type:char(36);column:limit_grade_id;null"`
	StudyProgramID            *uuid.UUID `gorm:"type:char(36);column:study_program_id;null"`
	IsMandatory               *bool      `gorm:"type:tinyint(1);column:is_mandatory;null"`
	IsPackage                 *bool      `gorm:"type:tinyint(1);column:is_package;null"`
	FieldStudyConcentrationID *uuid.UUID `gorm:"type:char(36);column:field_study_concentration_id"`
	CreatedAt                 int64      `gorm:"type:bigint;column:created_at"`
	CreatedBy                 *uuid.UUID `gorm:"type:char(36);column:created_by"`
	UpdatedAt                 *int64     `gorm:"type:bigint;column:updated_at"`
	UpdatedBy                 *uuid.UUID `gorm:"type:char(36);column:updated_by"`
	DeletedAt                 *int64     `gorm:"type:bigint;column:deleted_at"`
	DeletedBy                 *uuid.UUID `gorm:"type:char(36);column:deleted_by"`

	/* Relation */
	SubjectCode                 *string `gorm:"column:subject_code"`
	SubjectNameID               *string `gorm:"column:subject_name_id"`
	SubjectNameEN               *string `gorm:"column:subject_name_en"`
	SubjectTotalSKS             *int    `gorm:"column:subject_total_sks"`
	GradeName                   *string `gorm:"column:grade_name"`
	GradeCode                   *string `gorm:"column:grade_code"`
	FieldStudyConcentrationName *string `gorm:"column:field_study_concentration_name"`
	FieldStudyConcentrationCode *string `gorm:"column:field_study_concentration_code"`
}

func (MstStudyProgramCurriculum) TableName() string {
	return "mst_study_program_curriculums"
}

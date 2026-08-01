package dto

import "github.com/google/uuid"

/* Request */
type MstStudyProgramCurriculumRequest struct {
	ID                        string      `json:"-"`
	CurriculumYearID          string      `json:"curriculum_year_id" validate:"required,uuid"`
	SubjectID                 string      `json:"subject_id" validate:"required,uuid"`
	SemesterNumberID          string      `json:"semester_number_id" validate:"required,uuid"`
	LimitGradeID              *string     `json:"limit_grade_id" validate:"omitempty,uuid"`
	StudyProgramID            *string     `json:"study_program_id" validate:"omitempty,uuid"`
	IsMandatory               *bool       `json:"is_mandatory" validate:""`
	FieldStudyConcentrationID *uuid.UUID  `json:"field_study_concentration_id"`
	SubjectPrerequisite       []uuid.UUID `json:"subject_prerequisites"`
}

// Package
type UpdatePackageMstStudyProgramCurriculumWithStudyProgramRequest struct {
	IsPackage        bool   `json:"is_package"`
	CurriculumYearID string `json:"curriculum_year_id" validate:"required"`
	StudyProgramID   string `json:"study_program_id" validate:"required"`
	SemesterNumberID string `json:"semester_number_id" validate:"required"`
}
type UpdatePackageMstStudyProgramCurriculumWithoutStudyProgramRequest struct {
	IsPackage        bool   `json:"is_package"`
	CurriculumYearID string `json:"curriculum_year_id" validate:"required"`
	SemesterNumberID string `json:"semester_number_id" validate:"required"`
	UserID           string `json:"-"`
}

/* Response */
type MstStudyProgramCurriculumResponse struct {
	ID                        uuid.UUID  `json:"id" gorm:"column:id"`
	SubjectCode               *string    `json:"subject_code" gorm:"column:subject_code"`
	CurriculumYearID          uuid.UUID  `json:"curriculum_year_id" gorm:"column:curriculum_year_id"`
	SubjectID                 uuid.UUID  `json:"subject_id" gorm:"column:subject_id"`
	SemesterNumberID          uuid.UUID  `json:"semester_number_id" gorm:"column:semester_number_id"`
	LimitGradeID              *string    `json:"limit_grade_id" gorm:"column:limit_grade_id"`
	FieldStudyConcentrationID *uuid.UUID `json:"field_study_concentration_id"`
	StudyProgramID            *uuid.UUID `json:"study_program_id" gorm:"column:study_program_id"`
	IsPackage                 *bool      `json:"is_package" validate:""`
	IsMandatory               *bool      `json:"is_mandatory" gorm:"column:is_mandatory"`

	SubjectNameID               *string `json:"subject_name_id" gorm:"column:subject_name_id"`
	SubjectNameEN               *string `json:"subject_name_en" gorm:"column:subject_name_en"`
	SubjectTotalSKS             *int    `json:"subject_total_sks" gorm:"column:subject_total_sks"`
	GradeName                   *string `json:"grade_name" gorm:"column:grade_name"`
	GradeCode                   *string `json:"grade_code" gorm:"column:grade_code"`
	FieldStudyConcentrationName *string `json:"field_study_concentration_name"`
	FieldStudyConcentrationCode *string `json:"field_study_concentration_code"`

	SubjectPrerequisite []MstStudyProgramCurriculumSubjectPrerequisiteResponse `json:"subject_prerequisites"`
}

type MstStudyProgramCurriculumResponseDetail struct {
	ID                        uuid.UUID  `json:"id" gorm:"column:id"`
	SubjectCode               *string    `json:"subject_code" gorm:"column:subject_code"`
	CurriculumYearID          uuid.UUID  `json:"curriculum_year_id" gorm:"column:curriculum_year_id"`
	SubjectID                 uuid.UUID  `json:"subject_id" gorm:"column:subject_id"`
	SemesterNumberID          uuid.UUID  `json:"semester_number_id" gorm:"column:semester_number_id"`
	LimitGradeID              *string    `json:"limit_grade_id" gorm:"column:limit_grade_id"`
	StudyProgramID            *uuid.UUID `json:"study_program_id" gorm:"column:study_program_id"`
	FieldStudyConcentrationID *uuid.UUID `json:"field_study_concentration_id"`
	IsMandatory               *bool      `json:"is_mandatory" gorm:"column:is_mandatory"`
	IsPackage                 *bool      `json:"is_package" validate:""`
	CreatedAt                 int64      `json:"created_at" gorm:"column:created_at"`
	CreatedBy                 *uuid.UUID `json:"created_by" gorm:"column:created_by"`
	UpdatedAt                 *int64     `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy                 *uuid.UUID `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt                 *int64     `json:"deleted_at" gorm:"column:deleted_at"`
	DeletedBy                 *uuid.UUID `json:"deleted_by" gorm:"column:deleted_by"`

	SubjectNameID               *string `json:"subject_name_id" gorm:"column:subject_name_id"`
	SubjectNameEN               *string `json:"subject_name_en" gorm:"column:subject_name_en"`
	SubjectTotalSKS             *int    `json:"subject_total_sks" gorm:"column:subject_total_sks"`
	GradeName                   *string `json:"grade_name" gorm:"column:grade_name"`
	GradeCode                   *string `json:"grade_code" gorm:"column:grade_code"`
	FieldStudyConcentrationName *string `json:"field_study_concentration_name"`
	FieldStudyConcentrationCode *string `json:"field_study_concentration_code"`

	SubjectPrerequisite []MstStudyProgramCurriculumSubjectPrerequisiteResponse `json:"subject_prerequisites"`
}

type TotalForStudyProgramCurriculum struct {
	SKS         int `json:"sks"`
	Mandatory   int `json:"mandatory"`
	NoMandatory int `json:"no_mandatory"`
}

type MstStudyProgramCurriculumWithTotalResponse struct {
	MstStudyProgramCurriculumsResponse []MstStudyProgramCurriculumResponse `json:"data"`
	Total                              TotalForStudyProgramCurriculum      `json:"total"`
}

package model

import (
	"github.com/google/uuid"
)

type MstSubject struct {
	ID                 uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	CurriculumYearID   uuid.UUID  `gorm:"type:char(36);column:curriculum_year_id;not null"`
	StudyProgramID     string     `gorm:"type:varchar(36);column:study_program_id;not null"`
	Code               string     `gorm:"type:varchar(50);column:code;not null"`
	NameID             string     `gorm:"type:nvarchar(225);column:name_id;not null"`
	NameEN             string     `gorm:"type:nvarchar(225);column:name_en;not null"`
	CourseTypeID       uuid.UUID  `gorm:"type:char(36);column:course_type_id;not null"`
	CourseGroupID      uuid.UUID  `gorm:"type:char(36);column:course_group_id;not null"`
	FaceToFaceSKS      int        `gorm:"type:int;column:face_to_face_sks;default:0;not null"`
	PracticumSKS       int        `gorm:"type:int;column:practicum_sks;default:0;not null"`
	FieldPracticeSKS   int        `gorm:"type:int;column:field_practice_sks;default:0;not null"`
	SimulationSKS      int        `gorm:"type:int;column:simulation_sks"`
	TotalSKS           int        `gorm:"type:int;column:total_sks;default:0;not null"`
	FieldOfStudiesID   uuid.UUID  `gorm:"type:char(36);column:field_of_studies_id;not null"`
	IsMKU              bool       `gorm:"type:bool;column:is_mku;default:false;not null"`
	IsSAP              bool       `gorm:"type:bool;column:is_sap;default:false;not null"`
	IsSilabus          bool       `gorm:"type:bool;column:is_silabus;default:false;not null"`
	IsTeachingMaterial bool       `gorm:"type:bool;column:is_teaching_material;default:false;not null"`
	IsDiktat           bool       `gorm:"type:bool;column:is_diktat;default:false;not null"`
	CreatedAt          int64      `gorm:"type:bigint;column:created_at"`
	CreatedBy          *uuid.UUID `gorm:"type:char(36);column:created_by"`
	UpdatedAt          *int64     `gorm:"type:bigint;column:updated_at"`
	UpdatedBy          *uuid.UUID `gorm:"type:char(36);column:updated_by"`
	DeletedAt          *int64     `gorm:"type:bigint;column:deleted_at"`
	DeletedBy          *uuid.UUID `gorm:"type:char(36);column:deleted_by"`

	// Additional Fields
	StudyProgramName   string `gorm:"column:study_program_name"`
	CurriculumYearName string `gorm:"column:curriculum_year_name"`
	CourseTypeName     string `gorm:"column:course_type_name"`
	CourseGroupName    string `gorm:"column:course_group_name"`
	FieldStudyName     string `gorm:"column:field_study_name"`

	SubjectID          string `json:"subject_id"`
	LecturerID         string `json:"lecturer_id"`
	SubjectNameID      string `json:"subject_name_id"`
	SubjectNameEN      string `json:"subject_name_en"`
	LecturerName       string `json:"lecturer_name"`
	LecturerFrontTitle string `json:"lecturer_front_title"`
	LecturerBackTitle  string `json:"lecturer_back_title"`
}

func (MstSubject) TableName() string {
	return "mst_subjects"
}

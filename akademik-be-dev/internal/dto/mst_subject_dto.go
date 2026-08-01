package dto

import (
	"github.com/google/uuid"
)

/* Request */

type MstSubjectRequest struct {
	CurriculumYearID   uuid.UUID `json:"curriculum_year_id" validate:"required"`
	StudyProgramID     string    `json:"study_program_id" validate:"required"`
	CourseTypeID       uuid.UUID `json:"course_type_id" validate:"required"`
	CourseGroupID      uuid.UUID `json:"course_group_id" validate:"required"`
	Code               string    `json:"code" validate:"required"`
	NameID             string    `json:"name_id" validate:"required"`
	NameEN             string    `json:"name_en" validate:"required"`
	FaceToFaceSKS      int       `json:"face_to_face_sks" validate:"gte=0"`
	PracticumSKS       int       `json:"practicum_sks" validate:"gte=0"`
	FieldPracticeSKS   int       `json:"field_practice_sks" validate:"gte=0"`
	SimulationSKS      int       `json:"simulation_sks" validate:"omitempty,gte=0"`
	TotalSKS           int       `json:"-"`
	FieldOfStudiesID   uuid.UUID `json:"field_of_studies_id" validate:"required"`
	IsMKU              bool      `json:"is_mku" validate:"bool"`
	IsSAP              bool      `json:"is_sap" validate:"bool"`
	IsSilabus          bool      `json:"is_silabus" validate:"bool"`
	IsTeachingMaterial bool      `json:"is_teaching_material" validate:"bool"`
	IsDiktat           bool      `json:"is_diktat" validate:"bool"`
	UserID             uuid.UUID `json:"-" validate:"omitempty"`

	SupportingLecturerID        []uuid.UUID `json:"supporting_lecturer_id" validate:"required,dive,required"`
	DeveloperRPSLecuterID       []uuid.UUID `json:"developer_rps_lecturer_id" validate:"required,dive,required"`
	SubjectCoordinatorLecuterID []uuid.UUID `json:"subject_coordinator_lecturer_id" validate:"required,dive,required"`
}

type MstSubjectUpdate struct {
	ID                 uuid.UUID `json:"-" validate:"required"`
	CurriculumYearID   uuid.UUID `json:"curriculum_year_id" validate:"required"`
	StudyProgramID     string    `json:"study_program_id" validate:"required"`
	CourseTypeID       uuid.UUID `json:"course_type_id" validate:"required"`
	CourseGroupID      uuid.UUID `json:"course_group_id" validate:"required"`
	Code               string    `json:"code" validate:"required"`
	NameID             string    `json:"name_id" validate:"required"`
	NameEN             string    `json:"name_en" validate:"required"`
	FaceToFaceSKS      int       `json:"face_to_face_sks" validate:"required,gte=0"`
	PracticumSKS       int       `json:"practicum_sks" validate:"gte=0"`
	FieldPracticeSKS   int       `json:"field_practice_sks" validate:"gte=0"`
	SimulationSKS      int       `json:"simulation_sks" validate:"omitempty,gte=0"`
	TotalSKS           int       `json:"-"`
	FieldOfStudiesID   uuid.UUID `json:"field_of_studies_id" validate:"required"`
	IsMKU              bool      `json:"is_mku" validate:"bool"`
	IsSAP              bool      `json:"is_sap" validate:"bool"`
	IsSilabus          bool      `json:"is_silabus" validate:"bool"`
	IsTeachingMaterial bool      `json:"is_teaching_material" validate:"bool"`
	IsDiktat           bool      `json:"is_diktat" validate:"bool"`
	UserID             uuid.UUID `json:"-" validate:"omitempty"`

	SupportingLecturerID        []uuid.UUID `json:"supporting_lecturer_id" validate:"required,dive,required"`
	DeveloperRPSLecuterID       []uuid.UUID `json:"developer_rps_lecturer_id" validate:"required,dive,required"`
	SubjectCoordinatorLecuterID []uuid.UUID `json:"subject_coordinator_lecturer_id" validate:"required,dive,required"`
}

/* Response */
type MstSubjectResponse struct {
	ID                 uuid.UUID `json:"id"`
	CurriculumYearID   uuid.UUID `json:"curriculum_year_id"`
	StudyProgramID     string    `json:"study_program_id"`
	Code               string    `json:"code"`
	NameID             string    `json:"name_id"`
	NameEN             string    `json:"name_en"`
	CourseTypeID       uuid.UUID `json:"course_type_id"`
	CourseGroupID      uuid.UUID `json:"course_group_id"`
	FaceToFaceSKS      int       `json:"face_to_face_sks"`
	PracticumSKS       int       `json:"practicum_sks"`
	FieldPracticeSKS   int       `json:"field_practice_sks"`
	SimulationSKS      int       `json:"simulation_sks"`
	TotalSKS           int       `json:"total_sks"`
	FieldOfStudiesID   uuid.UUID `json:"field_of_studies_id"`
	IsMKU              bool      `json:"is_mku"`
	IsSAP              bool      `json:"is_sap"`
	IsSilabus          bool      `json:"is_silabus"`
	IsTeachingMaterial bool      `json:"is_teaching_material"`
	IsDiktat           bool      `json:"is_diktat"`
	CreatedAt          int64     `json:"created_at"`
	UpdatedAt          *int64    `json:"updated_at"`
	DeletedAt          *int64    `json:"deleted_at"`

	// Fields
	StudyProgramName   string `json:"study_program_name"`
	CurriculumYearName string `json:"curriculum_year_name"`
	CourseTypeName     string `json:"course_type_name"`
	CourseGroupName    string `json:"course_group_name"`
	FieldStudyName     string `json:"field_study_name"`

	SupportingLecturers         []Lecturer `json:"supporting_lecturers"`
	DeveloperRPSLecturers       []Lecturer `json:"developer_rps_lecturers"`
	SubjectCoordinatorLecturers []Lecturer `json:"subject_coordinator_lecturers"`
}

type Lecturer struct {
	ID                 string `json:"id"`
	SubjectID          string `json:"subject_id"`
	LecturerID         string `json:"lecturer_id"`
	SubjectNameID      string `json:"subject_name_id"`
	SubjectNameEN      string `json:"subject_name_en"`
	LecturerName       string `json:"lecturer_name"`
	LecturerFrontTitle string `json:"lecturer_front_title"`
	LecturerBackTitle  string `json:"lecturer_back_title"`
}

type MstSubjectResponseForSearch struct {
	ID     uuid.UUID `json:"id"`
	Code   string    `json:"code"`
	NameID string    `json:"name_id"`
	NameEN string    `json:"name_en"`
}

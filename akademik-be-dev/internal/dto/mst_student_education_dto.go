package dto

import (
	"github.com/google/uuid"
)

/* Request */
type MstStudentEducationRequest struct {
	ID                  string  `form:"-"`
	StudentID           string  `form:"-"`
	InstitutionName     *string `form:"institution_name" validate:"omitempty,stringMax=225"`
	SchoolMajor         *string `form:"school_major" validate:"omitempty,stringMax=50"`
	NISN                *string `form:"nisn" validate:"omitempty,stringMax=30"`
	NationalExamScore   *string `form:"national_exam_score" validate:"omitempty,numeric"`
	CertificateNumber   *string `form:"certificate_number" validate:"omitempty"`
	CertificateFilepath *string `form:"certificate_filepath" validate:"omitempty"`
	TranscriptsFilepath *string `form:"transcripts_filepath" validate:"omitempty"`
}

/* Response */
type MstStudentEducationResponse struct {
	ID                  uuid.UUID  `json:"id"`
	StudentID           uuid.UUID  `json:"student_id"`
	SchoolID            *uuid.UUID `json:"school_id"`
	EducationalLevelID  *uuid.UUID `json:"educational_level_id"`
	InstitutionName     *string    `json:"institution_name"`
	SchoolMajor         *string    `json:"school_major"`
	YearOfGraduation    *string    `json:"year_of_graduation"`
	NISN                *string    `json:"nisn"`
	ProvinceID          *uuid.UUID `json:"province_id"`
	CityID              *uuid.UUID `json:"city_id"`
	NationalExamScore   *float64   `json:"national_exam_score"`
	CertificateNumber   *string    `json:"certificate_number"`
	CertificateFilepath *string    `json:"certificate_filepath"`
	CertificateFileURL  *string    `json:"certificate_file_url"`
	TranscriptsFilepath *string    `json:"transcripts_filepath"`
	TranscriptsFileURL  *string    `json:"transcripts_file_url"`
	CreatedAt           int64      `json:"created_at"`
	UpdatedAt           *int64     `json:"updated_at"`
}

package model

import (
	"github.com/google/uuid"
)

type MstStudentEducation struct {
	ID                  uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	StudentID           uuid.UUID  `gorm:"type:char(36);column:student_id;not null"`
	SchoolID            *uuid.UUID `gorm:"type:char(36);column:school_id"`
	EducationalLevelID  *uuid.UUID `gorm:"type:char(36);column:educational_level_id"`
	InstitutionName     *string    `gorm:"type:nvarchar(225);column:institution_name"`
	SchoolMajor         *string    `gorm:"type:nvarchar(50);column:school_major"`
	YearOfGraduation    *string    `gorm:"type:varchar(4);column:year_of_graduation"`
	NISN                *string    `gorm:"type:varchar(30);column:nisn"`
	ProvinceID          *uuid.UUID `gorm:"type:char(36);column:province_id"`
	CityID              *uuid.UUID `gorm:"type:char(36);column:city_id"`
	NationalExamScore   *float64   `gorm:"type:decimal(5,2);column:national_exam_score"`
	CertificateNumber   *string    `gorm:"type:varchar(100);column:certificate_number"`
	CertificateFilepath *string    `gorm:"type:varchar(max);column:certificate_filepath"`
	TranscriptsFilepath *string    `gorm:"type:varchar(max);column:transcripts_filepath"`
	CreatedAt           int64      `gorm:"type:bigint;column:created_at"`
	UpdatedAt           *int64     `gorm:"type:bigint;column:updated_at"`
	DeletedAt           *int64     `gorm:"type:bigint;column:deleted_at"`

	// Additional Fields
	CertificateFileURL *string `gorm:"column:certificate_file_url"`
	TranscriptsFileURL *string `gorm:"column:transcripts_file_url"`
}

func (MstStudentEducation) TableName() string {
	return "mst_student_educations"
}

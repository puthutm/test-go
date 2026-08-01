package model

import "github.com/google/uuid"

type MstStudentDocument struct {
	ID                    uuid.UUID `gorm:"type:char(36);column:id;primaryKey"`
	StudentID             uuid.UUID `gorm:"type:char(36);column:student_id;not null"`
	StudentStudyProgramID uuid.UUID `gorm:"type:char(36);column:student_study_program_id;not null"`
	DocumentID            uuid.UUID `gorm:"type:char(36);column:document_id;not null"`
	DocumentPath          string    `gorm:"type:varchar(max);column:document_path;not null"`
	Status                string    `gorm:"type:varchar(100);column:status;not null"`
	CreatedAt             int64     `gorm:"type:bigint;column:created_at"`
	UpdatedAt             *int64    `gorm:"type:bigint;column:updated_at"`
	DeletedAt             *int64    `gorm:"type:bigint;column:deleted_at"`
}

func (MstStudentDocument) TableName() string {
	return "mst_student_documents"
}

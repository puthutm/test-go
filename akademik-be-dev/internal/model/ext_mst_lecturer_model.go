package model

import "github.com/google/uuid"

type MstLecturer struct {
	ID               uuid.UUID `gorm:"type:char(36);column:lecturer_id"`
	NIP              *string   `gorm:"type:varchar(50);column:lecturer_nip"`
	Name             *string   `gorm:"type:nvarchar(255);column:lecturer_name"`
	NIDN             *string   `gorm:"type:varchar(200);column:lecturer_nidn"`
	Gender           *string   `gorm:"type:varchar(10);column:lecturer_gender"`
	Phone            *string   `gorm:"type:varchar(20);column:lecturer_phone"`
	Email            *string   `gorm:"type:varchar(150);column:lecturer_email"`
	Status           *string   `gorm:"type:varchar(50);column:lecturer_status"`
	StudyProgramName *string   `gorm:"type:varchar(150);column:study_program_name"`
}

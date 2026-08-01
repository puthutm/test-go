package model

import (
	"time"

	"github.com/google/uuid"
)

type MstStudentFamily struct {
	ID                 uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	StudentID          uuid.UUID  `gorm:"type:char(36);column:student_id;not null"`
	Name               string     `gorm:"type:nvarchar(225);column:name;not null"`
	NIK                string     `gorm:"type:varchar(16);column:nik;not null"`
	EducationalLevelID *uuid.UUID `gorm:"type:char(36);column:educational_level_id"`
	Type               *string    `gorm:"type:varchar(100);column:type"`
	Phone              *string    `gorm:"type:nvarchar(225);column:phone"`
	Phone2             *string    `gorm:"type:nvarchar(225);column:phone2"`
	Email              *string    `gorm:"type:nvarchar(225);column:email"`
	Kinship            *string    `gorm:"type:nvarchar(225);column:kinship"`
	StatusKinship      *string    `gorm:"type:nvarchar(225);column:status_kinship"`
	LifeStatus         *string    `gorm:"type:nvarchar(225);column:life_status"`
	Address            *string    `gorm:"type:varchar(max);column:address"`
	BirthPlaceID       *uuid.UUID `gorm:"type:char(36);column:birth_place_id"`
	BirthDate          *time.Time `gorm:"type:date;column:birth_date"`
	JobID              *uuid.UUID `gorm:"type:char(36);column:job_id"`
	Income             *float64   `gorm:"type:decimal(18,2);column:income"`
	CreatedAt          int64      `gorm:"type:bigint;column:created_at"`
	UpdatedAt          *int64     `gorm:"type:bigint;column:updated_at"`
	DeletedAt          *int64     `gorm:"type:bigint;column:deleted_at"`

	// Additional Fields
	EducationalLevelName *string `gorm:"column:educational_level_name"`
	BirthPlaceName       *string `gorm:"column:birth_place_name"`
	JobName              *string `gorm:"column:job_name"`
}

func (MstStudentFamily) TableName() string {
	return "mst_student_families"
}

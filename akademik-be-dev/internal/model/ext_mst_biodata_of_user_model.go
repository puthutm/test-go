package model

import (
	"time"

	"github.com/google/uuid"
)

type MstBiodataOfUser struct {
	ID               uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	NipNim           string     `gorm:"type:nvarchar(20);column:nip_nim;not null"`
	FrontTitle       *string    `gorm:"type:nvarchar(50);column:front_title"`
	BackTitle        *string    `gorm:"type:nvarchar(50);column:back_title"`
	BirthPlace       *string    `gorm:"type:nvarchar(200);column:birth_place"`
	BirthDate        *time.Time `gorm:"type:date;column:birth_date"`
	Gender           string     `gorm:"type:nvarchar(30);column:gender;not null"`
	MarriageStatusID uuid.UUID  `gorm:"type:char(36);column:marriage_status_id;not null"`
	ReligionID       uuid.UUID  `gorm:"type:char(36);column:religion_id;not null"`
	EthnicID         uuid.UUID  `gorm:"type:char(36);column:ethnic_id;not null"`
	Weight           float64    `gorm:"type:float;column:weight;not null"`
	Height           float64    `gorm:"type:float;column:height;not null"`
	BloodTypeID      *uuid.UUID `gorm:"type:char(36);column:blood_type_id"`
	CountryID        *uuid.UUID `gorm:"type:char(36);column:country_id"`
	Status           *string    `gorm:"type:nvarchar(50);column:status"`
	UserID           uuid.UUID  `gorm:"type:char(36);column:user_id;not null"`
	CreatedAt        *int64     `gorm:"type:bigint;column:created_at"`
	UpdatedAt        *int64     `gorm:"type:bigint;column:updated_at"`
	DeletedAt        *int64     `gorm:"type:bigint;column:deleted_at"`
}

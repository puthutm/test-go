package model

import "github.com/google/uuid"

type MstUser struct {
	ID              uuid.UUID `gorm:"type:char(36);column:id;primaryKey"`
	Name            string    `gorm:"type:nvarchar(255);column:name;not null"`
	Username        string    `gorm:"type:nvarchar(255);column:username;not null"`
	NIK             *string   `gorm:"type:varchar(18);column:nik"`
	Email           *string   `gorm:"type:nvarchar(255);column:email"`
	EmailVerifiedAt *int64    `gorm:"type:bigint;column:email_verified_at"`
	EmailPrivate    *string   `gorm:"type:nvarchar(255);column:email_private"`
	Phone           *string   `gorm:"type:nvarchar(25);column:phone"`
	Avatar          *string   `gorm:"type:nvarchar(255);column:avatar"`
	GoogleAvatar    *string   `gorm:"type:nvarchar(255);column:google_avatar"`
	RememberToken   *string   `gorm:"type:nvarchar(100);column:remember_token"`
	CreatedAt       *int64    `gorm:"type:bigint;column:created_at"`
	UpdatedAt       *int64    `gorm:"type:bigint;column:updated_at"`
	DeletedAt       *int64    `gorm:"type:bigint;column:deleted_at"`
}

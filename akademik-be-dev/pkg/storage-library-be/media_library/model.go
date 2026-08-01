package medialibrary

import (
	"github.com/google/uuid"
)

type TrxMediaLibrary struct {
	ID        uuid.UUID  `gorm:"column:id;type:char(36);primaryKey" json:"id"`
	Filename  string     `gorm:"column:filename;type:nvarchar(max)" json:"filename"`
	SubjectID uuid.UUID  `gorm:"column:subject_id;type:char(36)" json:"subject_id"`
	FolderID  uuid.UUID  `gorm:"column:folder_id;type:char(36)" json:"folder_id"`
	Filepath  string     `gorm:"column:filepath;type:varchar(max)" json:"filepath"`
	Filetype  int        `gorm:"column:filetype" json:"filetype"`
	Filesize  float64    `gorm:"column:filesize;type:decimal(18,2)" json:"filesize"`
	Type      string     `gorm:"column:type;type:varchar(50);default:public" json:"type"`
	UserID    *uuid.UUID `gorm:"column:user_id;type:char(36)" json:"user_id,omitempty"`

	CreatedAt *int64     `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `gorm:"column:created_by;type:char(36)" json:"created_by,omitempty"`
	UpdatedAt *int64     `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by;type:char(36)" json:"updated_by,omitempty"`
	DeletedAt *int64     `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by;type:char(36)" json:"deleted_by,omitempty"`

	NameOfCreated string `gorm:"column:name_of_created" json:"name_of_created"`
}

func (TrxMediaLibrary) TableName() string {
	return "trx_media_libraries"
}

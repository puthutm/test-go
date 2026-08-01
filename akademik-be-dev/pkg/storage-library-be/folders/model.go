package folders

import "github.com/google/uuid"

type MstFolder struct {
	ID       uuid.UUID  `gorm:"column:id;type:char(36);primaryKey" json:"id"`
	Name     string     `gorm:"column:name;type:nvarchar(225)" json:"name"`
	ParentID *uuid.UUID `gorm:"column:parent_id;type:char(36)" json:"parent_id,omitempty"`
	Type     string     `gorm:"column:type;type:varchar(50);default:public" json:"type"`
	UserID   *uuid.UUID `gorm:"column:user_id;type:char(36)" json:"user_id,omitempty"`

	CreatedAt *int64     `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `gorm:"column:created_by;type:char(36)" json:"created_by,omitempty"`
	UpdatedAt *int64     `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by;type:char(36)" json:"updated_by,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by;type:char(36)" json:"deleted_by,omitempty"`
	DeletedAt *int64     `gorm:"column:deleted_at" json:"deleted_at,omitempty"`

	NameOfCreated string `gorm:"column:name_of_created" json:"name_of_created"`
}

func (MstFolder) TableName() string {
	return "mst_folders"
}

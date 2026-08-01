package folders

import "github.com/google/uuid"

type MstFolderResponse struct {
	ID       uuid.UUID  `json:"id"`
	Name     string     `json:"name"`
	ParentID *uuid.UUID `json:"parent_id,omitempty"`
	Type     string     `json:"type"`
	UserID   *uuid.UUID `json:"user_id,omitempty"`

	CreatedAt *int64     `json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
	UpdatedAt *int64     `json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
	DeletedBy *uuid.UUID `json:"deleted_by,omitempty"`
	DeletedAt *int64     `json:"deleted_at,omitempty"`

	// relation
	NameOfCreated string `json:"name_of_created"`
}

type MstFolderRequest struct {
	Name     string     `json:"name"`
	ParentID *uuid.UUID `json:"parent_id,omitempty"`
}

type MstFolderRequest_GetFileByFolderAndSubject struct {
	FolderID  uuid.UUID `json:"folder_id"`
	SubjectID uuid.UUID `json:"subject_id"`
}

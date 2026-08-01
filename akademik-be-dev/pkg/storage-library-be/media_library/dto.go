package medialibrary

import "github.com/google/uuid"

type TrxMediaLibraryResponse struct {
	ID        uuid.UUID  `json:"id"`
	Filename  string     `json:"filename"`
	SubjectID uuid.UUID  `json:"subject_id"`
	FolderID  uuid.UUID  `json:"folder_id"`
	Filepath  string     `json:"filepath"`
	Filetype  int        `json:"filetype"`
	Filesize  float64    `json:"filesize"`
	Type      string     `json:"type"`
	UserID    *uuid.UUID `json:"user_id,omitempty"`

	CreatedAt *int64     `json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
	UpdatedAt *int64     `json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
	DeletedAt *int64     `json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `json:"deleted_by,omitempty"`

	NameOfCreated string `json:"name_of_created"`
}

package command

import (
	"github.com/google/uuid"
	"unsia.ac.id/akademic_be/pkg/storage-library-be/folders"
)

type MstMediaLibraryRequest_GetFileByFolderAndSubject struct {
	FolderID  uuid.UUID
	SubjectID uuid.UUID
}

func ConvertDTOtoCommandMediaLibrary_GetFileByFolderAndSubject(
	req folders.MstFolderRequest_GetFileByFolderAndSubject,
) MstMediaLibraryRequest_GetFileByFolderAndSubject {
	return MstMediaLibraryRequest_GetFileByFolderAndSubject{
		FolderID:  req.FolderID,
		SubjectID: req.SubjectID,
	}
}

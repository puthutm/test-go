package command

import (
	"github.com/google/uuid"
	"gitlab.unsia.ac.id/icems/storage-library-be/folders"
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

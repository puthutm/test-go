package medialibrary

func ConvertToTrxMediaLibraryResponse(media TrxMediaLibrary) TrxMediaLibraryResponse {
	return TrxMediaLibraryResponse{
		ID:            media.ID,
		Filename:      media.Filename,
		SubjectID:     media.SubjectID,
		FolderID:      media.FolderID,
		Filepath:      media.Filepath,
		Filetype:      media.Filetype,
		Filesize:      media.Filesize,
		Type:          media.Type,
		UserID:        media.UserID,
		CreatedAt:     media.CreatedAt,
		CreatedBy:     media.CreatedBy,
		UpdatedAt:     media.UpdatedAt,
		UpdatedBy:     media.UpdatedBy,
		DeletedAt:     media.DeletedAt,
		DeletedBy:     media.DeletedBy,
		NameOfCreated: media.NameOfCreated,
	}
}

func ConvertToTrxMediaLibraryResponsePointer(media TrxMediaLibrary) *TrxMediaLibraryResponse {
	return &TrxMediaLibraryResponse{
		ID:            media.ID,
		Filename:      media.Filename,
		SubjectID:     media.SubjectID,
		FolderID:      media.FolderID,
		Filepath:      media.Filepath,
		Filetype:      media.Filetype,
		Filesize:      media.Filesize,
		Type:          media.Type,
		UserID:        media.UserID,
		CreatedAt:     media.CreatedAt,
		CreatedBy:     media.CreatedBy,
		UpdatedAt:     media.UpdatedAt,
		UpdatedBy:     media.UpdatedBy,
		DeletedAt:     media.DeletedAt,
		DeletedBy:     media.DeletedBy,
		NameOfCreated: media.NameOfCreated,
	}
}

package folders

func ConvertToMstFolderResponsePointer(folder MstFolder) *MstFolderResponse {
	return &MstFolderResponse{
		ID:            folder.ID,
		Name:          folder.Name,
		ParentID:      folder.ParentID,
		Type:          folder.Type,
		UserID:        folder.UserID,
		CreatedAt:     folder.CreatedAt,
		CreatedBy:     folder.CreatedBy,
		UpdatedAt:     folder.UpdatedAt,
		UpdatedBy:     folder.UpdatedBy,
		DeletedAt:     folder.DeletedAt,
		DeletedBy:     folder.DeletedBy,
		NameOfCreated: folder.NameOfCreated,
	}
}

func ConvertToMstFolderResponse(folder MstFolder) MstFolderResponse {
	return MstFolderResponse{
		ID:            folder.ID,
		Name:          folder.Name,
		ParentID:      folder.ParentID,
		Type:          folder.Type,
		UserID:        folder.UserID,
		CreatedAt:     folder.CreatedAt,
		CreatedBy:     folder.CreatedBy,
		UpdatedAt:     folder.UpdatedAt,
		UpdatedBy:     folder.UpdatedBy,
		DeletedAt:     folder.DeletedAt,
		DeletedBy:     folder.DeletedBy,
		NameOfCreated: folder.NameOfCreated,
	}
}

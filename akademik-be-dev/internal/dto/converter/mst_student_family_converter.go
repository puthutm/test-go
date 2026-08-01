package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstStudentFamilyModelToResponse(model model.MstStudentFamily) *dto.MstStudentFamilyResponse {
	return &dto.MstStudentFamilyResponse{
		ID:                   model.ID,
		StudentID:            model.StudentID,
		Name:                 model.Name,
		NIK:                  model.NIK,
		EducationalLevelID:   model.EducationalLevelID,
		EducationalLevelName: model.EducationalLevelName,
		Type:                 model.Type,
		Phone:                model.Phone,
		Phone2:               model.Phone2,
		Email:                model.Email,
		Kinship:              model.Kinship,
		StatusKinship:        model.StatusKinship,
		LifeStatus:           model.LifeStatus,
		Address:              model.Address,
		BirthPlaceID:         model.BirthPlaceID,
		BirthPlaceName:       model.BirthPlaceName,
		BirthDate:            model.BirthDate,
		JobID:                model.JobID,
		JobName:              model.JobName,
		Income:               model.Income,
		CreatedAt:            model.CreatedAt,
		UpdatedAt:            model.UpdatedAt,
	}
}

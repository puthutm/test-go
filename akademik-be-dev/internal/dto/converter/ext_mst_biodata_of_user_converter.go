package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstBiodataOfUserToResponse(model model.MstBiodataOfUser) *dto.MstBiodataOfUserResponse {
	return &dto.MstBiodataOfUserResponse{
		ID:               model.ID,
		NipNim:           model.NipNim,
		FrontTitle:       model.FrontTitle,
		BackTitle:        model.BackTitle,
		BirthPlace:       model.BirthPlace,
		BirthDate:        model.BirthDate,
		Gender:           model.Gender,
		MarriageStatusID: model.MarriageStatusID,
		ReligionID:       model.ReligionID,
		EthnicID:         model.EthnicID,
		Weight:           model.Weight,
		Height:           model.Height,
		BloodTypeID:      model.BloodTypeID,
		CountryID:        model.CountryID,
		Status:           model.Status,
		UserID:           model.UserID,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		DeletedAt:        model.DeletedAt,
	}
}

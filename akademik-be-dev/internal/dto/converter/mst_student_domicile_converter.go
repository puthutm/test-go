package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstStudentDomicileModelToResponse(model model.MstStudentDomicile) *dto.MstStudentDomicileResponse {
	return &dto.MstStudentDomicileResponse{
		ID:           model.ID,
		StudentID:    model.StudentID,
		CountryID:    model.CountryID,
		CountryName:  model.CountryName,
		ProvinceID:   model.ProvinceID,
		ProvinceName: model.ProvinceName,
		CityID:       model.CityID,
		CityName:     model.CityName,
		DistrictID:   model.DistrictID,
		DistrictName: model.DistrictName,
		VillageID:    model.VillageID,
		VillageName:  model.VillageName,
		RT:           model.RT,
		RW:           model.RW,
		Address:      model.Address,
		PostalCode:   model.PostalCode,
		Distance:     model.Distance,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}

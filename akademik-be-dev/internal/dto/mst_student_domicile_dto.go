package dto

import (
	"github.com/google/uuid"
)

/* Request */
type MstStudentDomicileRequest struct {
	ID         string  `json:"-"`
	StudentID  string  `json:"-"`
	CountryID  *string `json:"country_id" validate:"omitempty,uuid"`
	ProvinceID *string `json:"province_id" validate:"omitempty,uuid"`
	CityID     *string `json:"city_id" validate:"omitempty,uuid"`
	DistrictID *string `json:"district_id" validate:"omitempty,uuid"`
	VillageID  *string `json:"village_id" validate:"omitempty,uuid"`
	RT         *string `json:"rt" validate:"omitempty,stringMax=5"`
	RW         *string `json:"rw" validate:"omitempty,stringMax=5"`
	Address    *string `json:"address" validate:"omitempty"`
	PostalCode *string `json:"postal_code" validate:"omitempty,stringMax=20"`
	Distance   *string `json:"distance" validate:"omitempty,numeric"`
}

/* Response */
type MstStudentDomicileResponse struct {
	ID           uuid.UUID  `json:"id"`
	StudentID    uuid.UUID  `json:"student_id"`
	CountryID    *uuid.UUID `json:"country_id"`
	CountryName  *string    `json:"country_name"`
	ProvinceID   *uuid.UUID `json:"province_id"`
	ProvinceName *string    `json:"province_name"`
	CityID       *uuid.UUID `json:"city_id"`
	CityName     *string    `json:"city_name"`
	DistrictID   *uuid.UUID `json:"district_id"`
	DistrictName *string    `json:"district_name"`
	VillageID    *uuid.UUID `json:"village_id"`
	VillageName  *string    `json:"village_name"`
	RT           *string    `json:"rt"`
	RW           *string    `json:"rw"`
	Address      *string    `json:"address"`
	PostalCode   *string    `json:"postal_code"`
	Distance     *string    `json:"distance"`
	CreatedAt    int64      `json:"created_at"`
	UpdatedAt    *int64     `json:"updated_at"`
}

package model

import "github.com/google/uuid"

type MstStudentDomicile struct {
	ID         uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	StudentID  uuid.UUID  `gorm:"type:char(36);column:student_id;not null"`
	CountryID  *uuid.UUID `gorm:"type:char(36);column:country_id"`
	ProvinceID *uuid.UUID `gorm:"type:char(36);column:province_id"`
	CityID     *uuid.UUID `gorm:"type:char(36);column:city_id"`
	DistrictID *uuid.UUID `gorm:"type:char(36);column:district_id"`
	VillageID  *uuid.UUID `gorm:"type:char(36);column:village_id"`
	RT         *string    `gorm:"type:varchar(5);column:rt"`
	RW         *string    `gorm:"type:varchar(5);column:rw"`
	Address    *string    `gorm:"type:varchar(max);column:address"`
	PostalCode *string    `gorm:"type:varchar(20);column:postal_code"`
	Distance   *string    `gorm:"type:varchar(20);column:distance"`
	CreatedAt  int64      `gorm:"type:bigint;column:created_at"`
	UpdatedAt  *int64     `gorm:"type:bigint;column:updated_at"`

	// Additional Fields
	CountryName  *string `gorm:"column:country_name"`
	ProvinceName *string `gorm:"column:province_name"`
	CityName     *string `gorm:"column:city_name"`
	DistrictName *string `gorm:"column:district_name"`
	VillageName  *string `gorm:"column:village_name"`
}

func (MstStudentDomicile) TableName() string {
	return "mst_student_domiciles"
}

package model

import (
	"time"

	"github.com/google/uuid"
)

type MstStudentBio struct {
	ID                       uuid.UUID  `gorm:"type:char(36);column:id;primaryKey"`
	NIK                      string     `gorm:"type:varchar(20);column:nik;not null"`
	Name                     string     `gorm:"type:nvarchar(225);column:name;not null"`
	EthnicID                 *uuid.UUID `gorm:"type:char(36);column:ethnic_id"`
	SchoolID                 *uuid.UUID `gorm:"type:char(36);column:school_id"`
	YearOfGraduation         *string    `gorm:"type:varchar(4);column:year_of_graduation"`
	NPSN                     *string    `gorm:"type:varchar(225);column:npsn"`
	ProvinceIDOfSchoolOrigin *uuid.UUID `gorm:"type:char(36);column:province_id_of_school_origin"`
	SchoolMajor              *string    `gorm:"type:nvarchar(50);column:school_major"`
	CityIDOfSchoolOrigin     *uuid.UUID `gorm:"type:char(36);column:city_id_of_school_origin"`
	BirthPlaceID             *uuid.UUID `gorm:"type:char(36);column:birth_place_id"`
	ReligionID               *uuid.UUID `gorm:"type:char(36);column:religion_id"`
	BirthDate                *time.Time `gorm:"type:date;column:birth_date"`
	Height                   *float64   `gorm:"type:decimal(5,2);column:height"`
	Gender                   *string    `gorm:"type:nvarchar(50);column:gender"`
	Weight                   *float64   `gorm:"type:decimal(5,2);column:weight"`
	Phone                    *string    `gorm:"type:varchar(30);column:phone"`
	PrivateEmail             *string    `gorm:"type:varchar(225);column:private_email"`
	CollegeEmail             *string    `gorm:"type:varchar(225);column:college_email"`
	NoPassport               *string    `gorm:"type:varchar(50);column:no_passport"`
	NoKK                     *string    `gorm:"type:varchar(50);column:no_kk"`
	CitizenshipID            *uuid.UUID `gorm:"type:char(36);column:citizenship_id"`
	JobID                    *uuid.UUID `gorm:"type:char(36);column:job_id"`
	Institution              *string    `gorm:"type:nvarchar(225);column:institution"`
	AlmamaterSizeID          *uuid.UUID `gorm:"type:char(36);column:almamater_size_id"`
	PersonalIncome           *float64   `gorm:"type:decimal(18,2);column:personal_income"`
	DomicileAddress          *string    `gorm:"type:varchar(max);column:domicile_address"`
	PhotoProfilePath         *string    `gorm:"type:varchar(max);column:photo_profile_path"`
	BloodTypeID              *uuid.UUID `gorm:"type:char(36);column:blood_type_id"`
	BackDegree               *string    `gorm:"type:varchar(100);column:back_degree"`
	TransportationID         *uuid.UUID `gorm:"type:char(36);column:transportation_id"`
	SintaID                  *string    `gorm:"type:varchar(225);column:sinta_id"`
	ScopusID                 *string    `gorm:"type:varchar(225);column:scopus_id"`
	GoogleScholar            *string    `gorm:"type:varchar(max);column:google_scholar"`
	SignaturePathFile        *string    `gorm:"type:varchar(max);column:signature_path_file"`
	NPWP                     *string    `gorm:"type:varchar(50);column:npwp"`
	NPWPFilepath             *string    `gorm:"type:varchar(max);column:npwp_filepath"`
	BPJSHealthcare           *string    `gorm:"type:varchar(50);column:bpjs_healthcare"`
	BPJSHealthcareFilepath   *string    `gorm:"type:varchar(max);column:bpjs_healthcare_filepath"`
	BPJSEmployment           *string    `gorm:"type:varchar(50);column:bpjs_employment"`
	BPJSEmploymentFilepath   *string    `gorm:"type:varchar(max);column:bpjs_employment_filepath"`
	BankID                   *uuid.UUID `gorm:"type:char(36);column:bank_id"`
	AccountNumber            *string    `gorm:"type:varchar(225);column:account_number"`
	AccountName              *string    `gorm:"type:varchar(225);column:account_name"`
	AccountFilepath          *string    `gorm:"type:varchar(max);column:account_filepath"`
	UserID                   *uuid.UUID `gorm:"type:char(36);column:user_id"`
	StatusID                 *uuid.UUID `gorm:"type:char(36);column:status_id"`
	CreatedAt                int64      `gorm:"type:bigint;column:created_at"`
	UpdatedAt                int64      `gorm:"type:bigint;column:updated_at"`
	DeletedAt                *int64     `gorm:"type:bigint;column:deleted_at"`

	// Additional Fields
	StatusName         string     `gorm:"type:nvarchar(225)"`
	ReligionName       string     `gorm:"type:nvarchar(225)"`
	EthnicName         string     `gorm:"type:nvarchar(225)"`
	BloodTypeName      string     `gorm:"type:nvarchar(225)"`
	LastNIM            string     `gorm:"column:last_nim;type:varchar(50)"`
	BirthPlaceName     string     `gorm:"column:birth_place_name"`
	NISN               string     `gorm:"column:nisn"`
	TransportationName string     `gorm:"type:nvarchar(225)"`
	CountryName        string     `gorm:"type:nvarchar(225)"`
	AlmamaterSizeName  string     `gorm:"type:nvarchar(225)"`
	JobName            string     `gorm:"type:nvarchar(225)"`
	StudyProgramID     *uuid.UUID `gorm:"type:char(36)"`
	StudyProgramName   string     `gorm:"type:nvarchar(225)"`
	BankName           string     `gorm:"column:bank_name"`
}

func (MstStudentBio) TableName() string {
	return "mst_student_bios"
}

package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type MstStudentBioRequest struct {
	ID uuid.UUID `json:"-"`
}

// TODO: update only user
type MstStudentBioUpdateOnlyUser struct {
	ID           uuid.UUID  `json:"-" validate:""`
	Name         string     `json:"name" validate:"required,min=1,stringMax=225"`
	BackDegree   *string    `json:"back_degree" validate:"omitempty,min=1,stringMax=100"`
	NIK          string     `json:"nik" validate:"required,len=16,nik"`
	NoKK         *string    `json:"no_kk" validate:"required,len=16"`
	BirthPlaceID *uuid.UUID `json:"birth_place_id" validate:"uuid"`
	BirthDate    *string    `json:"birth_date" validate:"required,datetime=2006-01-02"`
	Gender       *string    `json:"gender" validate:"required,min=1,stringMax=50"`
	StatusID     *uuid.UUID `json:"status_id" validate:"required,uuid"`
	ReligionID   *uuid.UUID `json:"religion_id" validate:"required,uuid"`
	EthnicID     *uuid.UUID `json:"ethnic_id" validate:"required,uuid"`
	Height       *float64   `json:"height" validate:"omitempty,gte=0"`
	Weight       *float64   `json:"weight" validate:"omitempty,gte=0"`
	BloodTypeID  *uuid.UUID `json:"blood_type_id" validate:"omitempty,uuid"`

	UserID uuid.UUID `json:"-"`
}
type MstStudentBioUpdateCompletenessOnlyUser struct {
	ID                uuid.UUID             `form:"id" validate:""`
	NoPassport        *string               `form:"no_passport" validate:"omitempty,stringMax=50"`
	GoogleScholar     *string               `form:"google_scholar" validate:"omitempty,stringMax=4000"`
	SintaID           *string               `form:"sinta_id" validate:"omitempty,stringMax=225"`
	ScopusID          *string               `form:"scopus_id" validate:"omitempty,stringMax=225"`
	SignaturePathFile *multipart.FileHeader `form:"signature_path_file"`

	UserID uuid.UUID `json:"-"`
}
type MstStudentBioUpdateInformationOnlyUser struct {
	ID               uuid.UUID  `json:"id" validate:""`
	PrivateEmail     *string    `json:"private_email" validate:"required,email,stringMax=225"`
	Phone            *string    `json:"phone" validate:"required,stringMax=30"`
	TransportationID *uuid.UUID `json:"transportation_id" validate:"required,uuid"`
	CitizenshipID    *uuid.UUID `json:"citizenship_id" validate:"required,uuid"`
	AlmamaterSizeID  *uuid.UUID `json:"almamater_size_id" validate:"required,uuid"`
	JobID            *uuid.UUID `json:"job_id" validate:"required,uuid"`

	UserID uuid.UUID `json:"-"`
}
type MstStudentBioUpdateDocumentOnlyUser struct {
	ID                     uuid.UUID             `form:"id" validate:""`
	NPWP                   *string               `form:"npwp" validate:"omitempty,stringMax=50"`
	NPWPFilePath           *multipart.FileHeader `form:"npwp_filepath"`
	BPJSHealthcare         *string               `form:"bpjs_healthcare" validate:"omitempty,stringMax=50"`
	BPJSHealthcareFilePath *multipart.FileHeader `form:"bpjs_healthcare_filepath" validate:"omitempty,stringMax=255"`
	BPJSEmployment         *string               `form:"bpjs_employment" validate:"omitempty,stringMax=50"`
	BPJSEmploymentFilePath *multipart.FileHeader `form:"bpjs_employment_filepath" validate:"omitempty,stringMax=255"`

	UserID uuid.UUID `json:"-"`
}
type MstStudentBioUpdateBankAccountOnlyUser struct {
	ID              uuid.UUID             `form:"id" validate:""`
	BankID          *uuid.UUID            `form:"bank_id" validate:"omitempty,uuid"`
	AccountNumber   *string               `form:"account_number" validate:"omitempty,stringMax=225"`
	AccountName     *string               `form:"account_name" validate:"omitempty,stringMax=225"`
	AccountFilePath *multipart.FileHeader `form:"account_filepath"`

	UserID uuid.UUID `json:"-"`
}

// TODO: update all
type MstStudentBioUpdate struct {
	MstStudentBioUpdateOnlyUser
}
type MstStudentBioUpdateCompleteness struct {
	MstStudentBioUpdateCompletenessOnlyUser
}
type MstStudentBioUpdateInformation struct {
	MstStudentBioUpdateInformationOnlyUser
}
type MstStudentUBiopdateDocument struct {
	MstStudentBioUpdateDocumentOnlyUser
}
type MstStudentBioUpdateBankAccount struct {
	MstStudentBioUpdateBankAccountOnlyUser
}

// TODO: response
type MstStudentBioResponse struct {
	ID                       uuid.UUID  `json:"id"`
	NIK                      string     `json:"nik"`
	Name                     string     `json:"name"`
	EthnicID                 *uuid.UUID `json:"ethnic_id"`
	SchoolID                 *uuid.UUID `json:"school_id"`
	YearOfGraduation         *string    `json:"year_of_graduation"`
	NPSN                     *string    `json:"npsn"`
	ProvinceIDOfSchoolOrigin *uuid.UUID `json:"province_id_of_school_origin"`
	SchoolMajor              *string    `json:"school_major"`
	CityIDOfSchoolOrigin     *uuid.UUID `json:"city_id_of_school_origin"`
	BirthPlaceID             *uuid.UUID `json:"birth_place_id"`
	ReligionID               *uuid.UUID `json:"religion_id"`
	BirthDate                *string    `json:"birth_date"`
	Height                   *float64   `json:"height"`
	Gender                   *string    `json:"gender"`
	Weight                   *float64   `json:"weight"`
	Phone                    *string    `json:"phone"`
	PrivateEmail             *string    `json:"private_email"`
	CollegeEmail             *string    `json:"college_email"`
	NoPassport               *string    `json:"no_passport"`
	NoKK                     *string    `json:"no_kk"`
	CitizenshipID            *uuid.UUID `json:"citizenship_id"`
	JobID                    *uuid.UUID `json:"job_id"`
	Institution              *string    `json:"institution"`
	AlmamaterSizeID          *uuid.UUID `json:"almamater_size_id"`
	PersonalIncome           *float64   `json:"personal_income"`
	DomicileAddress          *string    `json:"domicile_address"`
	PhotoProfilePath         *string    `json:"photo_profile_path"`
	BloodTypeID              *uuid.UUID `json:"blood_type_id"`
	BackDegree               *string    `json:"back_degree"`
	TransportationID         *uuid.UUID `json:"transportation_id"`
	SintaID                  *string    `json:"sinta_id"`
	ScopusID                 *string    `json:"scopus_id"`
	GoogleScholar            *string    `json:"google_scholar"`
	SignaturePathFile        *string    `json:"signature_path_file"`
	NPWP                     *string    `json:"npwp"`
	NPWPFilepath             *string    `json:"npwp_filepath"`
	BPJSHealthcare           *string    `json:"bpjs_healthcare"`
	BPJSHealthcareFilepath   *string    `json:"bpjs_healthcare_filepath"`
	BPJSEmployment           *string    `json:"bpjs_employment"`
	BPJSEmploymentFilepath   *string    `json:"bpjs_employment_filepath"`
	BankID                   *uuid.UUID `json:"bank_id"`
	AccountNumber            *string    `json:"account_number"`
	AccountName              *string    `json:"account_name"`
	AccountFilepath          *string    `json:"account_filepath"`
	UserID                   *uuid.UUID `json:"user_id"`
	StatusID                 *uuid.UUID `json:"status_id"`
	CreatedAt                int64      `json:"created_at"`
	UpdatedAt                int64      `json:"updated_at"`
	DeletedAt                *int64     `json:"deleted_at"`

	// field tambahan
	StatusName    string `json:"status_name,omitempty"`
	ReligionName  string `json:"religion_name,omitempty"`
	EthnicName    string `json:"ethnic_name,omitempty"`
	BloodTypeName string `json:"blood_type_name,omitempty"`
	LastNIM       string `json:"last_nim,omitempty"`

	// field tambahan completeneses
	NISN string `json:"nisn,omitempty"`

	// field tambahan infomation
	TransportationName string     `json:"transportation_name,omitempty"`
	CountryName        string     `json:"country_name,omitempty"`
	AlmamaterSizeName  string     `json:"almamater_size_name,omitempty"`
	JobName            string     `json:"job_name,omitempty"`
	StudyProgramID     *uuid.UUID `json:"study_program_id,omitempty"`
	StudyProgramName   string     `json:"study_program_name,omitempty"`
}

type MstStudentBioGeneralResponse struct {
	ID           uuid.UUID  `json:"id"`
	NIK          string     `json:"nik"`
	Name         string     `json:"name"`
	EthnicID     *uuid.UUID `json:"ethnic_id"`
	BirthPlaceID *uuid.UUID `json:"birth_place_id"`
	BirthDate    *string    `json:"birth_date"`
	Gender       *string    `json:"gender"`
	Height       *float64   `json:"height"`
	Weight       *float64   `json:"weight"`
	NoKK         *string    `json:"no_kk"`
	BackDegree   *string    `json:"back_degree"`
	StatusID     *uuid.UUID `json:"status_id"`
	BloodTypeID  *uuid.UUID `json:"blood_type_id"`
	ReligionID   *uuid.UUID `json:"religion_id"`
	CreatedAt    int64      `json:"created_at"`
	UpdatedAt    int64      `json:"updated_at"`
	DeletedAt    *int64     `json:"deleted_at"`

	JobName        string `json:"job_name"`
	StatusName     string `json:"status_name"`
	ReligionName   string `json:"religion_name"`
	EthnicName     string `json:"ethnic_name"`
	BloodTypeName  string `json:"blood_type_name"`
	LastNIM        string `json:"last_nim"`
	BirthPlaceName string `json:"birth_place_name"`
}

type MstStudentBioInfomationResponse struct {
	ID               uuid.UUID  `json:"id"`
	CollegeEmail     *string    `json:"college_email"`
	PrivateEmail     *string    `json:"private_email"`
	Phone            *string    `json:"phone"`
	TransportationID *uuid.UUID `json:"transportation_id"`
	CitizenshipID    *uuid.UUID `json:"citizenship_id"`
	AlmamaterSizeID  *uuid.UUID `json:"almamater_size_id"`
	JobID            *uuid.UUID `json:"job_id"`
	StudyProgramID   *uuid.UUID `json:"study_program_id"`

	StudyProgramName   string `json:"study_program_name"`
	TransportationName string `json:"transportation_name"`
	CountryName        string `json:"citizenship_name"`
	AlmamaterSizeName  string `json:"almamater_size_name"`
	JobName            string `json:"job_name"`
}

type MstStudentBioCompletenesResponse struct {
	ID                uuid.UUID `json:"id"`
	NoPassport        *string   `json:"no_passport"`
	GoogleScholar     *string   `json:"google_scholar"`
	SintaID           *string   `json:"sinta_id"`
	ScopusID          *string   `json:"scopus_id"`
	SignaturePathFile *string   `json:"signature_path_file"`
	NISN              string    `json:"nisn"`
}

type MstStudentBioDocumentResponse struct {
	ID                     uuid.UUID `json:"id"`
	Npwp                   *string   `json:"npwp"`
	NPWPFilepath           *string   `json:"npwp_filepath"`
	BPJSHealthcare         *string   `json:"bpjs_healthcare"`
	BPJSHealthcareFilepath *string   `json:"bpjs_healthcare_filepath"`
	BPJSEmployment         *string   `json:"bpjs_employment"`
	BPJSEmploymentFilepath *string   `json:"bpjs_employment_filepath"`
}

type MstStudentBioBankAccountResponse struct {
	ID              uuid.UUID  `json:"id"`
	BankID          *uuid.UUID `json:"bank_id"`
	AccountNumber   *string    `json:"account_number"`
	AccountName     *string    `json:"account_name"`
	AccountFilepath *string    `json:"account_filepath"`
	BankName        string     `json:"bank_name"`
}

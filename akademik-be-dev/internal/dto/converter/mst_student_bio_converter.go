package converter

import (
	"time"

	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

// TODO: request only user to model
func ConvertMstStudentBioUpdateOnlyUserToModelPointer(req dto.MstStudentBioUpdateOnlyUser, data *model.MstStudentBio) {
	var datetime *time.Time
	if req.BirthDate == nil {
		datetime = nil
	} else {
		v := req.BirthDate
		datetime, _ = utils.StringToDatePointer(v)
	}
	data.ID = req.ID
	data.NIK = req.NIK
	data.Name = req.Name
	data.BackDegree = req.BackDegree
	data.NoKK = req.NoKK
	data.BirthPlaceID = req.BirthPlaceID
	data.BirthDate = datetime
	data.Gender = req.Gender
	data.StatusID = req.StatusID
	data.ReligionID = req.ReligionID
	data.EthnicID = req.EthnicID
	data.Height = req.Height
	data.Weight = req.Weight
	data.BloodTypeID = req.BloodTypeID
}

func ConvertMstStudentBioUpdateCompletenessOnlyUserToModelPointer(req dto.MstStudentBioUpdateCompletenessOnlyUser, data *model.MstStudentBio) {
	data.ID = req.ID
	data.NoPassport = req.NoPassport
	data.GoogleScholar = req.GoogleScholar
	data.SintaID = req.SintaID
	data.ScopusID = req.ScopusID
}

func ConvertMstStudentBioUpdateInformationOnlyUserToModelPointer(req dto.MstStudentBioUpdateInformationOnlyUser, data *model.MstStudentBio) {
	data.ID = req.ID
	data.PrivateEmail = req.PrivateEmail
	data.Phone = req.Phone
	data.TransportationID = req.TransportationID
	data.CitizenshipID = req.CitizenshipID
	data.AlmamaterSizeID = req.AlmamaterSizeID
	data.JobID = req.JobID
}

func ConvertMstStudentBioUpdateDocumentOnlyUserToModelPointer(req dto.MstStudentBioUpdateDocumentOnlyUser, data *model.MstStudentBio) {
	data.ID = req.ID
	data.NPWP = req.NPWP
	// NPWPFilepath:           req.NPWPFilePath,
	data.BPJSHealthcare = req.BPJSHealthcare
	// BPJSHealthcareFilepath: req.BPJSHealthcareFilePath,
	data.BPJSEmployment = req.BPJSEmployment
	// BPJSEmploymentFilepath: req.BPJSEmploymentFilePath,
}

func ConvertMstStudentBioUpdateBankAccountOnlyUserToModelPointer(req dto.MstStudentBioUpdateBankAccountOnlyUser, data *model.MstStudentBio) {
	data.ID = req.ID
	data.BankID = req.BankID
	data.AccountNumber = req.AccountNumber
	data.AccountName = req.AccountName
	// AccountFilepath: req.AccountFilePath,
}

// TODO: not only user
func ConvertMstStudentBioUpdateToModelPointer(req dto.MstStudentBioUpdate) *model.MstStudentBio {
	var datetime *time.Time
	if req.BirthDate == nil {
		datetime = nil
	} else {
		v := req.BirthDate
		datetime, _ = utils.StringToDatePointer(v)
	}
	return &model.MstStudentBio{
		ID:           req.ID,
		NIK:          req.NIK,
		Name:         req.Name,
		BackDegree:   req.BackDegree,
		NoKK:         req.NoKK,
		BirthPlaceID: req.BirthPlaceID,
		BirthDate:    datetime,
		Gender:       req.Gender,
		StatusID:     req.StatusID,
		ReligionID:   req.ReligionID,
		EthnicID:     req.EthnicID,
		Height:       req.Height,
		Weight:       req.Weight,
		BloodTypeID:  req.BloodTypeID,
	}
}

func ConvertMstStudentBioUpdateCompletenessToModelPointer(req dto.MstStudentBioUpdateCompleteness) *model.MstStudentBio {
	return &model.MstStudentBio{
		ID:            req.ID,
		NoPassport:    req.NoPassport,
		GoogleScholar: req.GoogleScholar,
		SintaID:       req.SintaID,
		ScopusID:      req.ScopusID,
	}
}

func ConvertMstStudentBioUpdateInformationToModelPointer(req dto.MstStudentBioUpdateInformation) *model.MstStudentBio {
	return &model.MstStudentBio{
		ID:               req.ID,
		PrivateEmail:     req.PrivateEmail,
		Phone:            req.Phone,
		TransportationID: req.TransportationID,
		CitizenshipID:    req.CitizenshipID,
		AlmamaterSizeID:  req.AlmamaterSizeID,
		JobID:            req.JobID,
	}
}

func ConvertMstStudentBioUpdateDocumentToModelPointer(req dto.MstStudentUBiopdateDocument) *model.MstStudentBio {
	return &model.MstStudentBio{
		ID:   req.ID,
		NPWP: req.NPWP,
		// NPWPFilepath:           req.NPWPFilePath,
		BPJSHealthcare: req.BPJSHealthcare,
		// BPJSHealthcareFilepath: req.BPJSHealthcareFilePath,
		BPJSEmployment: req.BPJSEmployment,
		// BPJSEmploymentFilepath: req.BPJSEmploymentFilePath,
	}
}

func ConvertMstStudentBioUpdateBankAccountToModelPointer(req dto.MstStudentBioUpdateBankAccount) *model.MstStudentBio {
	return &model.MstStudentBio{
		ID:            req.ID,
		BankID:        req.BankID,
		AccountNumber: req.AccountNumber,
		AccountName:   req.AccountName,
		// AccountFilepath: req.AccountFilePath,
	}
}

// TODO: Model to Response
func ConvertMstStudentBioToResponse(model *model.MstStudentBio) *dto.MstStudentBioResponse {
	var dateString *string
	if model.BirthDate == nil {
		dateString = nil
	} else {
		v := model.BirthDate
		dateString = utils.DateToStringPointer(v)
	}
	return &dto.MstStudentBioResponse{
		ID:                       model.ID,
		NIK:                      model.NIK,
		Name:                     model.Name,
		EthnicID:                 model.EthnicID,
		SchoolID:                 model.SchoolID,
		YearOfGraduation:         model.YearOfGraduation,
		NPSN:                     model.NPSN,
		ProvinceIDOfSchoolOrigin: model.ProvinceIDOfSchoolOrigin,
		SchoolMajor:              model.SchoolMajor,
		CityIDOfSchoolOrigin:     model.CityIDOfSchoolOrigin,
		BirthPlaceID:             model.BirthPlaceID,
		ReligionID:               model.ReligionID,
		BirthDate:                dateString,
		Height:                   model.Height,
		Gender:                   model.Gender,
		Weight:                   model.Weight,
		Phone:                    model.Phone,
		PrivateEmail:             model.PrivateEmail,
		CollegeEmail:             model.CollegeEmail,
		NoPassport:               model.NoPassport,
		NoKK:                     model.NoKK,
		CitizenshipID:            model.CitizenshipID,
		JobID:                    model.JobID,
		Institution:              model.Institution,
		AlmamaterSizeID:          model.AlmamaterSizeID,
		PersonalIncome:           model.PersonalIncome,
		DomicileAddress:          model.DomicileAddress,
		PhotoProfilePath:         model.PhotoProfilePath,
		BloodTypeID:              model.BloodTypeID,
		BackDegree:               model.BackDegree,
		TransportationID:         model.TransportationID,
		SintaID:                  model.SintaID,
		ScopusID:                 model.ScopusID,
		GoogleScholar:            model.GoogleScholar,
		SignaturePathFile:        model.SignaturePathFile,
		NPWP:                     model.NPWP,
		NPWPFilepath:             model.NPWPFilepath,
		BPJSHealthcare:           model.BPJSHealthcare,
		BPJSHealthcareFilepath:   model.BPJSHealthcareFilepath,
		BPJSEmployment:           model.BPJSEmployment,
		BPJSEmploymentFilepath:   model.BPJSEmploymentFilepath,
		BankID:                   model.BankID,
		AccountNumber:            model.AccountNumber,
		AccountName:              model.AccountName,
		AccountFilepath:          model.AccountFilepath,
		UserID:                   model.UserID,
		StatusID:                 model.StatusID,
		CreatedAt:                model.CreatedAt,
		UpdatedAt:                model.UpdatedAt,
		DeletedAt:                model.DeletedAt,

		// relasi
		StatusName:    model.StatusName,
		EthnicName:    model.EthnicName,
		ReligionName:  model.ReligionName,
		BloodTypeName: model.BloodTypeName,
		LastNIM:       model.LastNIM,

		//
		NISN:               model.NISN,
		TransportationName: model.TransportationName,
		CountryName:        model.CountryName,
		AlmamaterSizeName:  model.AlmamaterSizeName,
		JobName:            model.JobName,
		StudyProgramID:     model.StudyProgramID,
		StudyProgramName:   model.StudyProgramName,
	}
}

func ConvertMstStudentBioGeneralToResponse(model *model.MstStudentBio) *dto.MstStudentBioGeneralResponse {
	var dateString *string
	if model.BirthDate == nil {
		dateString = nil
	} else {
		v := model.BirthDate
		dateString = utils.DateToStringPointer(v)
	}
	return &dto.MstStudentBioGeneralResponse{
		ID:           model.ID,
		NIK:          model.NIK,
		Name:         model.Name,
		EthnicID:     model.EthnicID,
		BirthPlaceID: model.BirthPlaceID,
		ReligionID:   model.ReligionID,
		BirthDate:    dateString,
		Height:       model.Height,
		Gender:       model.Gender,
		Weight:       model.Weight,
		NoKK:         model.NoKK,
		BackDegree:   model.BackDegree,
		BloodTypeID:  model.BloodTypeID,
		StatusID:     model.StatusID,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    model.DeletedAt,

		// relasi
		JobName:        model.JobName,
		StatusName:     model.StatusName,
		EthnicName:     model.EthnicName,
		ReligionName:   model.ReligionName,
		BloodTypeName:  model.BloodTypeName,
		LastNIM:        model.LastNIM,
		BirthPlaceName: model.BirthPlaceName,
	}
}

func ConvertMstStudentBioInformationToResponse(model *model.MstStudentBio) *dto.MstStudentBioInfomationResponse {
	return &dto.MstStudentBioInfomationResponse{
		ID:               model.ID,
		Phone:            model.Phone,
		PrivateEmail:     model.PrivateEmail,
		CollegeEmail:     model.CollegeEmail,
		CitizenshipID:    model.CitizenshipID,
		AlmamaterSizeID:  model.AlmamaterSizeID,
		TransportationID: model.TransportationID,
		JobID:            model.JobID,

		// relasi
		TransportationName: model.TransportationName,
		CountryName:        model.CountryName,
		AlmamaterSizeName:  model.AlmamaterSizeName,
		JobName:            model.JobName,
		StudyProgramID:     model.StudyProgramID,
		StudyProgramName:   model.StudyProgramName,
	}
}

func ConvertMstStudentBioCompletenesToResponse(model *model.MstStudentBio) *dto.MstStudentBioCompletenesResponse {
	return &dto.MstStudentBioCompletenesResponse{
		ID:                model.ID,
		NoPassport:        model.NoPassport,
		SintaID:           model.SintaID,
		ScopusID:          model.ScopusID,
		GoogleScholar:     model.GoogleScholar,
		SignaturePathFile: model.SignaturePathFile,

		// relasi
		NISN: model.NISN,
	}
}

func ConvertMstStudentBioToDocumentResponse(student *model.MstStudentBio) *dto.MstStudentBioDocumentResponse {
	return &dto.MstStudentBioDocumentResponse{
		ID:                     student.ID,
		Npwp:                   student.NPWP,
		NPWPFilepath:           student.NPWPFilepath,
		BPJSHealthcare:         student.BPJSHealthcare,
		BPJSHealthcareFilepath: student.BPJSHealthcareFilepath,
		BPJSEmployment:         student.BPJSEmployment,
		BPJSEmploymentFilepath: student.BPJSEmploymentFilepath,
	}
}

func ConvertMstStudentBioToBankAccountResponse(student *model.MstStudentBio) *dto.MstStudentBioBankAccountResponse {
	return &dto.MstStudentBioBankAccountResponse{
		ID:              student.ID,
		BankID:          student.BankID,
		AccountNumber:   student.AccountNumber,
		AccountName:     student.AccountName,
		AccountFilepath: student.AccountFilepath,
		BankName:        student.BankName,
	}
}

package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func MstStudentEducationModelToResponse(model model.MstStudentEducation) *dto.MstStudentEducationResponse {
	return &dto.MstStudentEducationResponse{
		ID:                  model.ID,
		StudentID:           model.StudentID,
		SchoolID:            model.SchoolID,
		EducationalLevelID:  model.EducationalLevelID,
		InstitutionName:     model.InstitutionName,
		SchoolMajor:         model.SchoolMajor,
		YearOfGraduation:    model.YearOfGraduation,
		NISN:                model.NISN,
		ProvinceID:          model.ProvinceID,
		CityID:              model.CityID,
		NationalExamScore:   model.NationalExamScore,
		CertificateNumber:   model.CertificateNumber,
		CertificateFilepath: model.CertificateFilepath,
		CertificateFileURL:  model.CertificateFileURL,
		TranscriptsFilepath: model.TranscriptsFilepath,
		TranscriptsFileURL:  model.TranscriptsFileURL,
		CreatedAt:           model.CreatedAt,
		UpdatedAt:           model.UpdatedAt,
	}
}

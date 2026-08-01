package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

// ConvertMstClassLecturerRequestToModelPointer converts request DTO to model pointer
func ConvertMstClassLecturerRequestToModelPointer(request dto.MstClassLecturerRequest, model *model.MstClassLecturer) {
	model.ID = utils.GenerateUUID()
	model.ClassID = request.ClassID
	model.LecturerID = request.LecturerID

	model.SubtituteLecturerID = request.SubtituteLecturerID
}

func ConvertMstClassLecturerUpdateToModelPointer(request dto.MstClassLecturerUpdate, model *model.MstClassLecturer) {
	model.ID = request.ID
	model.ClassID = request.ClassID
	model.LecturerID = request.LecturerID

	model.SubtituteLecturerID = request.SubtituteLecturerID
}

// ConvertModelToMstClassLecturerResponsePointer converts model to response DTO pointer
func ConvertModelToMstClassLecturerResponsePointer(model *model.MstClassLecturer) *dto.MstClassLecturerResponse {
	return &dto.MstClassLecturerResponse{
		ID:                    model.ID,
		ClassID:               model.ClassID,
		LecturerID:            model.LecturerID,
		SubtituteLecturerID:   model.SubtituteLecturerID,
		LecturerName:          model.LecturerName,
		SubtituteLecturerName: model.SubtituteLecturerName,
	}
}

// ConvertModelToMstClassLecturerResponse converts model to response DTO
func ConvertModelToMstClassLecturerResponse(model *model.MstClassLecturer) dto.MstClassLecturerResponse {
	return dto.MstClassLecturerResponse{
		ID:                    model.ID,
		ClassID:               model.ClassID,
		LecturerID:            model.LecturerID,
		SubtituteLecturerID:   model.SubtituteLecturerID,
		LecturerName:          model.LecturerName,
		SubtituteLecturerName: model.SubtituteLecturerName,
	}
}

func MstLecturerModelByProgramHeadToResponse(model model.MstLecturer) *dto.MstLecturerByProgramHeadResponse {
	return &dto.MstLecturerByProgramHeadResponse{
		ID:               model.ID,
		NIP:              model.NIP,
		Name:             model.Name,
		NIDN:             model.NIDN,
		Gender:           model.Gender,
		Phone:            model.Phone,
		Email:            model.Email,
		Status:           model.Status,
		StudyProgramName: model.StudyProgramName,
	}
}

func ConvertMstClassLecturerGetByAcademicPeriodAndSubjectAndUserForLecturerToResponse(
	data *model.MstClassLecturerGetByAcademicPeriodAndSubjectAndUserForLecturer,
) dto.MstClassLecturerGetByAcademicPeriodAndSubjectAndUserForLecturerResponse {
	return dto.MstClassLecturerGetByAcademicPeriodAndSubjectAndUserForLecturerResponse{
		ClassName:      data.ClassName,
		ClassID:        data.ClassID,
		ClassCode:      data.ClassCode,
		StudyProgramID: data.StudyProgramID,
	}
}

func ConvertMstClassGetResultSubjectAndClassCountToResponse(data model.MstClassGetResultSubjectAndClassCount) dto.MstClassGetResultSubjectAndClassCountResponse {
	return dto.MstClassGetResultSubjectAndClassCountResponse{
		AcademicPeriodeID:   data.AcademicPeriodeID,
		AcademicPeriodeName: data.AcademicPeriodeName,
		SubjectID:           data.SubjectID,
		SubjectNameID:       data.SubjectNameID,
		SubjectNameEN:       data.SubjectNameEN,
		StudyProgramID:      data.StudyProgramID,
		LecturerID:          data.LecturerID,
		ClassCount:          data.ClassCount,
	}
}

func ConvertMstClassGetResultSubjectAndClassCountArrayToResponse(data []model.MstClassGetResultSubjectAndClassCount) []dto.MstClassGetResultSubjectAndClassCountResponse {
	var result []dto.MstClassGetResultSubjectAndClassCountResponse
	for _, item := range data {
		result = append(result, ConvertMstClassGetResultSubjectAndClassCountToResponse(item))
	}
	return result
}

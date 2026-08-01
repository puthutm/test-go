package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func TrxFinalProjectProposalModelToResponse(model model.TrxFinalProjectProposal) *dto.TrxFinalProjectProposalResponse {
	return &dto.TrxFinalProjectProposalResponse{
		ID:                     model.ID,
		StudentID:              model.StudentID,
		TitleID:                model.TitleID,
		TitleEN:                model.TitleEN,
		Topic:                  model.Topic,
		AcademicPeriodeID:      model.AcademicPeriodeID,
		StudyProgramID:         model.StudyProgramID,
		Abstract:               model.Abstract,
		FilePath:               model.FilePath,
		Status:                 model.Status,
		Date:                   model.Date,
		ConfirmationStatusDate: model.ConfirmationStatusDate,
		ConfirmationBy:         model.ConfirmationBy,
		Feedback:               model.Feedback,
	}
}

func TrxFinalProjectProposalModelPointerToResponse(model *model.TrxFinalProjectProposal) *dto.TrxFinalProjectProposalResponse {
	return &dto.TrxFinalProjectProposalResponse{
		ID:                     model.ID,
		StudentID:              model.StudentID,
		TitleID:                model.TitleID,
		TitleEN:                model.TitleEN,
		Topic:                  model.Topic,
		AcademicPeriodeID:      model.AcademicPeriodeID,
		StudyProgramID:         model.StudyProgramID,
		Abstract:               model.Abstract,
		FilePath:               model.FilePath,
		Status:                 model.Status,
		Date:                   model.Date,
		ConfirmationStatusDate: model.ConfirmationStatusDate,
		ConfirmationBy:         model.ConfirmationBy,
		Feedback:               model.Feedback,
	}
}

func TrxFinalProjectProposalModelPointerToResponseWithMentorLecturerResponse(
	model *model.TrxFinalProjectProposal, modelMentor []model.TrxFinalProjectProposalMentorLecturer,
) *dto.TrxFinalProjectProposalResponseForAdmin {
	dtoMentors := make([]dto.TrxFinalProjectProposalMentorLecturerResponse, 0)

	for _, v := range modelMentor {
		dtoMentors = append(dtoMentors, TrxFinalProjectProposalMentorLecturerModelToResponseNoPointer(v))
	}

	return &dto.TrxFinalProjectProposalResponseForAdmin{
		TrxFinalProjectProposalResponse: dto.TrxFinalProjectProposalResponse{
			ID:                     model.ID,
			StudentID:              model.StudentID,
			TitleID:                model.TitleID,
			TitleEN:                model.TitleEN,
			Topic:                  model.Topic,
			AcademicPeriodeID:      model.AcademicPeriodeID,
			StudyProgramID:         model.StudyProgramID,
			Abstract:               model.Abstract,
			FilePath:               model.FilePath,
			Status:                 model.Status,
			Date:                   model.Date,
			ConfirmationStatusDate: model.ConfirmationStatusDate,
			ConfirmationBy:         model.ConfirmationBy,
			Feedback:               model.Feedback,
		},
		MentorLecturers: dtoMentors,
	}
}

func TrxFinalProjectProposalModelByUserIDToResponse(model model.TrxFinalProjectProposal) *dto.TrxFinalProjectProposalByUserIDResponse {
	return &dto.TrxFinalProjectProposalByUserIDResponse{
		ID:      model.ID,
		TitleID: model.TitleID,
		TitleEN: model.TitleEN,
		Status:  model.Status,
		Date:    model.Date,
	}
}

func TrxFinalProjectProposalModelProgramHeadAllToResponse(model model.TrxFinalProjectProposal) *dto.TrxFinalProjectProposalsResponseForAdmin {
	res := new(dto.TrxFinalProjectProposalsResponseForAdmin)

	res.TrxFinalProjectProposalByUserIDResponse = *TrxFinalProjectProposalModelByUserIDToResponse(model)
	res.Topic = model.Topic
	res.StudentName = model.StudentName
	res.StudyProgramName = model.StudyProgramName
	res.LecturerName = model.LecturerName

	return res
}

func TrxFinalProjectProposalModelByStudenIDandStudyProgramIDToResponse(model model.TrxFinalProjectProposal) *dto.TrxFinalProjectProposalByStudenIDandStudyProgramIDResponse {
	return &dto.TrxFinalProjectProposalByStudenIDandStudyProgramIDResponse{
		StudentName:          model.StudentName,
		StudentNIM:           model.StudentNIM,
		StudyProgramName:     model.StudyProgramName,
		AcademicPeriodeName:  model.AcademicPeriodeName,
		RegistrationPathName: model.RegistrationPathName,
	}
}

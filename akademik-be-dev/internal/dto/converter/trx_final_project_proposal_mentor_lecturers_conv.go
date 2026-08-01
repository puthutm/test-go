package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func TrxFinalProjectProposalMentorLecturerModelToResponse(model model.TrxFinalProjectProposalMentorLecturer) *dto.TrxFinalProjectProposalMentorLecturerResponse {
	return &dto.TrxFinalProjectProposalMentorLecturerResponse{
		ID:                     model.ID,
		FinalProjectProposalID: model.FinalProjectProposalID,
		LecturerID:             model.LecturerID,
		AssignDate:             model.AssignDate,
		AssignBy:               model.AssignBy,
		LecturerName:           model.LecturerName,
	}
}

func TrxFinalProjectProposalMentorLecturerModelToResponseNoPointer(model model.TrxFinalProjectProposalMentorLecturer) dto.TrxFinalProjectProposalMentorLecturerResponse {
	return dto.TrxFinalProjectProposalMentorLecturerResponse{
		ID:                     model.ID,
		FinalProjectProposalID: model.FinalProjectProposalID,
		LecturerID:             model.LecturerID,
		AssignDate:             model.AssignDate,
		AssignBy:               model.AssignBy,
		LecturerName:           model.LecturerName,
	}
}

func TrxFinalProjectProposalMentorLecturerModelPointerToResponse(model *model.TrxFinalProjectProposalMentorLecturer) *dto.TrxFinalProjectProposalMentorLecturerResponse {
	if model == nil {
		return nil
	}

	return &dto.TrxFinalProjectProposalMentorLecturerResponse{
		ID:                     model.ID,
		FinalProjectProposalID: model.FinalProjectProposalID,
		LecturerID:             model.LecturerID,
		AssignDate:             model.AssignDate,
		AssignBy:               model.AssignBy,
		LecturerName:           model.LecturerName,
	}
}

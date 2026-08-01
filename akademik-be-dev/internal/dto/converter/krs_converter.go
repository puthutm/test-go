package converter

import (
	"time"

	"github.com/google/uuid"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

// TrxKrsLecturerStudentModelToResponse - Convert model to DTO response
func TrxKrsLecturerStudentModelToResponse(m model.TrxKrsLecturerStudent) dto.TrxKrsLecturerStudentResponse {
	return dto.TrxKrsLecturerStudentResponse{
		KrsHeaderID:         m.KrsHeaderID,
		AcademicPeriodeID:   m.AcademicPeriodeID,
		AcademicPeriodeName: m.AcademicPeriodeName,
		StudentName:         m.StudentName,
		StudentNIM:          m.StudentNIM,
		TotalSKS:            m.TotalSKS,
		StudentID:           m.StudentID,
		CreatedAt:           m.CreatedAt,
	}
}

// TrxKrsLecturerStudentDetailModelToResponse - Convert models to DTO response for detail
func TrxKrsLecturerStudentDetailModelToResponse(
	detail model.TrxKrsLecturerStudentDetail,
	totalSKS model.TrxKrsLecturerStudentTotalSKS,
	items []model.TrxKrsLecturerStudentItem,
) dto.TrxKrsLecturerStudentDetailResponse {
	itemResponses := make([]dto.TrxKrsLecturerItemResponse, 0, len(items))
	for _, item := range items {
		itemResponses = append(itemResponses, dto.TrxKrsLecturerItemResponse{
			KrsItemID:     item.KrsItemID,
			SubjectNameID: item.SubjectNameID,
			SubjectNameEN: item.SubjectNameEN,
			ClassCode:     item.ClassCode,
			ClassName:     item.ClassName,
			TotalSKS:      item.TotalSKS,
		})
	}

	return dto.TrxKrsLecturerStudentDetailResponse{
		StudentID:           detail.StudentID,
		StudentNIM:          detail.StudentNIM,
		StudentName:         detail.StudentName,
		StudyProgramName:    detail.StudyProgramName,
		AcademicPeriodeName: detail.AcademicPeriodeName,
		TotalSKSTaken:       totalSKS.TotalSKSTaken,
		KRSItems:            itemResponses,
	}
}

// TrxKrsLecturerStudentItemUpdateRequestToModel - Convert request to model for update
func TrxKrsLecturerStudentItemUpdateRequestToModel(
	req dto.TrxKrsLecturerStudentItemUpdateStatusRequest,
	updatedBy uuid.UUID,
) model.TrxKrsLecturerStudentItemUpdate {
	return model.TrxKrsLecturerStudentItemUpdate{
		KrsItemID:    uuid.MustParse(req.KrsItemID),
		ItemStatus:   req.ItemStatus,
		RejectReason: req.RejectReason,
		UpdatedAt:    time.Now().UnixMilli(),
		UpdatedBy:    updatedBy,
	}
}

// TrxKrsLecturerStudentItemUpdateModelToResponse - Convert model to response DTO
func TrxKrsLecturerStudentItemUpdateModelToResponse(
	m model.TrxKrsLecturerStudentItemUpdate,
) dto.TrxKrsLecturerStudentItemUpdateResponse {
	return dto.TrxKrsLecturerStudentItemUpdateResponse{
		KrsItemID:    m.KrsItemID,
		ItemStatus:   m.ItemStatus,
		RejectReason: m.RejectReason,
		UpdatedAt:    m.UpdatedAt,
		UpdatedBy:    m.UpdatedBy,
	}
}

// TrxKrsAcademicPeriodModelToResponse - Convert model to DTO response
func TrxKrsAcademicPeriodModelToResponse(m model.TrxKrsAcademicPeriod) dto.TrxKrsAcademicPeriodResponse {
	var startDate, endDate int64
	if m.StartDateOfCollege != nil {
		startDate = m.StartDateOfCollege.UnixMilli()
	}
	if m.EndDateOfCollege != nil {
		endDate = m.EndDateOfCollege.UnixMilli()
	}

	return dto.TrxKrsAcademicPeriodResponse{
		ID:                 m.ID,
		Code:               m.Code,
		Fullname:           m.Fullname,
		Shortname:          m.Shortname,
		IsActive:           m.IsActive,
		StartDateOfCollege: &startDate,
		EndDateOfCollege:   &endDate,
	}
}

// TrxKrsClassForPickModelToResponse - Convert model to DTO response
func TrxKrsClassForPickModelToResponse(m model.TrxKrsClassForPick) dto.TrxKrsClassForPickResponse {
	return dto.TrxKrsClassForPickResponse{
		ClassID:     m.ClassID,
		SubjectID:   m.SubjectID,
		SubjectCode: m.SubjectCode,
		SubjectName: m.SubjectName,
		Schedule:    m.Schedule,
		ClassCode:   m.ClassCode,
		ClassName:   m.ClassName,
		SKS:         m.SKS,
		Capacity:    m.Capacity,
		UsedQuota:   m.UsedQuota,
		QuotaText:   m.QuotaText,
		ButtonState: m.ButtonState,
	}
}

// TrxKrsProgramHeadClassModelToResponse - Convert model to DTO response
func TrxKrsProgramHeadClassModelToResponse(m model.TrxKrsProgramHeadClass) dto.TrxKrsProgramHeadClassResponse {
	return dto.TrxKrsProgramHeadClassResponse{
		AcademicPeriodeName: m.AcademicPeriodeName,
		ClassCode:           m.ClassCode,
		ClassName:           m.ClassName,
		LecturerName:        m.LecturerName,
		Schedule:            m.Schedule,
		ClassQuota:          m.ClassQuota,
		Filled:              m.Filled,
		Remaining:           m.Remaining,
		ClassID:             m.ClassID,
		CreatedAt:           m.CreatedAt,
	}
}

// TrxKrsSavedItemModelToResponse - Convert model to DTO response
func TrxKrsSavedItemModelToResponse(m model.TrxKrsSavedItem) dto.TrxKrsSavedItemResponse {
	return dto.TrxKrsSavedItemResponse{
		KrsItemID:     m.KrsItemID,
		KrsID:         m.KrsID,
		ClassID:       m.ClassID,
		SubjectID:     m.SubjectID,
		SubjectCode:   m.SubjectCode,
		SubjectName:   m.SubjectName,
		Schedule:      m.Schedule,
		LecturerNames: m.LecturerNames,
		ClassName:     m.ClassName,
		SKS:           m.SKS,
		ItemStatus:    m.ItemStatus,
	}
}

// TrxKrsTakeClassModelToResponse - Convert model to response DTO
func TrxKrsTakeClassModelToResponse(m model.TrxKrsTakeClassResult) dto.TrxKrsTakeClassResponse {
	return dto.TrxKrsTakeClassResponse{
		KrsID:                 m.KrsID,
		KrsItemID:             m.KrsItemID,
		AcademicPeriodeID:     m.AcademicPeriodeID,
		PrevAcademicPeriodeID: m.PrevAcademicPeriodeID,
		PrevIPS:               m.PrevIPS,
		MaxSKSAllowed:         m.MaxSKSAllowed,
		TotalSKSSelected:      m.TotalSKSSelected,
		RemainingSKS:          m.RemainingSKS,
		ClassCapacity:         m.ClassCapacity,
		UsedQuotaAfter:        m.UsedQuotaAfter,
	}
}

// TrxKrsMaxSksInfoModelToResponse - Convert model to DTO response
func TrxKrsMaxSksInfoModelToResponse(m model.TrxKrsMaxSksInfo) dto.TrxKrsMaxSksInfoResponse {
	return dto.TrxKrsMaxSksInfoResponse{
		StudentID:               m.StudentID,
		AcademicPeriodeID:       m.AcademicPeriodeID,
		AcademicPeriodeIDBefore: m.AcademicPeriodeIDBefore,
		IpsBefore:               m.IpsBefore,
		SksLimitID:              m.SksLimitID,
		IpsMin:                  m.IpsMin,
		IpsMax:                  m.IpsMax,
		MaxSks:                  m.MaxSks,
	}
}
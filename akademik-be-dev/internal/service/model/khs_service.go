package servicemodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type KhsService interface {
	GetKHSByUserID(ctx context.Context) (*dto.KhsDataResponse, error)
}

type khsService struct {
	log     *logrus.Logger
	db      *gorm.DB
	khsRepo repositorymodel.KhsRepository
}

func NewKhsService(log *logrus.Logger, db *gorm.DB, khsRepo repositorymodel.KhsRepository) KhsService {
	return &khsService{
		log:     log,
		db:      db,
		khsRepo: khsRepo,
	}
}

func (s *khsService) GetKHSByUserID(ctx context.Context) (*dto.KhsDataResponse, error) {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	tx := s.db.WithContext(ctx)
	subjects, semesters, notPassedRaw, err := s.khsRepo.GetKHSByUserID(tx, ctx, user.ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), user.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	// Map raw subjects by Semester ID for easy grouping
	subjectMap := make(map[string][]dto.KhsSubjectResponse)
	for _, s := range subjects {
		subjectMap[s.AcademicPeriodeId] = append(subjectMap[s.AcademicPeriodeId], dto.KhsSubjectResponse{
			SubjectCode:  s.SubjectCode,
			SubjectName:  s.SubjectName,
			TotalSks:     s.TotalSks,
			FinalScore:   s.FinalScore,
			GradeCode:    s.GradeCode,
			QualityValue: s.QualityValue,
			Weight:       s.Bobot,
			IsPassed:     s.IsPassed == 1,
		})
	}

	// Build Semester Responses
	semesterResponses := make([]dto.KhsSemesterResponse, 0, len(semesters))
	for _, sm := range semesters {
		semesterResponses = append(semesterResponses, dto.KhsSemesterResponse{
			AcademicPeriodeId:   sm.AcademicPeriodeId,
			AcademicPeriodeName: sm.AcademicPeriodeName,
			TotalWeight:         sm.TotalBobot,
			TotalSks:            sm.TotalSks,
			Ips:                 sm.Ips,
			Subjects:            subjectMap[sm.AcademicPeriodeId],
		})
	}

	// Build Not Passed Responses
	notPassedResponses := make([]dto.KhsNotPassedResponse, 0, len(notPassedRaw))
	for _, np := range notPassedRaw {
		notPassedResponses = append(notPassedResponses, dto.KhsNotPassedResponse{
			AcademicPeriodeId:   np.AcademicPeriodeId,
			AcademicPeriodeName: np.AcademicPeriodeName,
			SubjectCode:         np.SubjectCode,
			SubjectName:         np.SubjectName,
			FinalScore:          np.FinalScore,
			GradeCode:           np.GradeCode,
		})
	}

	return &dto.KhsDataResponse{
		Semesters: semesterResponses,
		NotPassed: notPassedResponses,
	}, nil
}

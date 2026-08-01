package servicemodel

import (
	"context"
	"encoding/csv"
	"io"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	internalutils "unsia.ac.id/akademic_be/internal/utils"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

const defaultPassword = "12345678"

type PortalStudentService interface {
	GetAllBatches(ctx context.Context) ([]dto.PortalStudentBatchResponse, error)
	CreateStudent(ctx context.Context, req dto.PortalStudentCreateRequest) error
	CreateStudentBulk(ctx context.Context, file io.Reader) (*dto.PortalStudentBulkResponse, error)
	GetStudentListWithCount(ctx context.Context, pg pageable.PageableRequestClass) (*pageable.PageableResponse[dto.PortalStudentListResponse], error)
}

type portalStudentService struct {
	log               *logrus.Logger
	db                *gorm.DB
	portalStudentRepo repositorymodel.PortalStudentRepository
}

func NewPortalStudentService(log *logrus.Logger, db *gorm.DB, portalStudentRepo repositorymodel.PortalStudentRepository) PortalStudentService {
	return &portalStudentService{
		log:               log,
		db:                db,
		portalStudentRepo: portalStudentRepo,
	}
}

func (s *portalStudentService) GetAllBatches(ctx context.Context) ([]dto.PortalStudentBatchResponse, error) {
	tx := s.db.WithContext(ctx)

	batches, err := s.portalStudentRepo.GetAllBatches(tx)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.PortalStudentBatchResponse, 0, len(batches))
	for _, b := range batches {
		res = append(res, dto.PortalStudentBatchResponse{
			BatchDetailID: b.BatchDetailID,
			BatchName:     b.BatchName,
		})
	}

	return res, nil
}

func (s *portalStudentService) CreateStudent(ctx context.Context, req dto.PortalStudentCreateRequest) error {
	tx := s.db.WithContext(ctx)

	req.ID = utils.GenerateUUID().String()

	password := req.Password
	if password == "" {
		password = defaultPassword
	}

	hashedPassword, err := internalutils.HashPassword(password)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)
		return err
	}
	req.Password = hashedPassword

	err = s.portalStudentRepo.CreateStudent(tx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

func (s *portalStudentService) CreateStudentBulk(ctx context.Context, file io.Reader) (*dto.PortalStudentBulkResponse, error) {
	reader := csv.NewReader(file)

	// Read header
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Normalize header names
	colIndex := make(map[string]int)
	for i, h := range header {
		colIndex[strings.TrimSpace(strings.ToLower(h))] = i
	}

	// Validate required columns
	requiredCols := []string{"batch_detail_id", "nik", "name", "email", "phone"}
	for _, col := range requiredCols {
		if _, ok := colIndex[col]; !ok {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Missing required column: "+col)
		}
	}

	_, hasPassword := colIndex["password"]

	result := &dto.PortalStudentBulkResponse{
		FailedRows: make([]dto.PortalStudentBulkFailedRow, 0),
	}
	rowNum := 1 // header is row 0

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		rowNum++
		result.TotalData++

		req := dto.PortalStudentCreateRequest{
			BatchDetailID: strings.TrimSpace(row[colIndex["batch_detail_id"]]),
			NIK:           strings.TrimSpace(row[colIndex["nik"]]),
			Name:          strings.TrimSpace(row[colIndex["name"]]),
			Email:         strings.TrimSpace(row[colIndex["email"]]),
			Phone:         strings.TrimSpace(row[colIndex["phone"]]),
		}

		if hasPassword {
			req.Password = strings.TrimSpace(row[colIndex["password"]])
		}

		req.ID = utils.GenerateUUID().String()

		password := req.Password
		if password == "" {
			password = defaultPassword
		}

		hashedPassword, hashErr := internalutils.HashPassword(password)
		if hashErr != nil {
			result.Failed++
			result.FailedRows = append(result.FailedRows, dto.PortalStudentBulkFailedRow{
				Row:   rowNum,
				Name:  req.Name,
				Email: req.Email,
				Error: "Failed to hash password",
			})
			continue
		}
		req.Password = hashedPassword

		tx := s.db.WithContext(ctx)
		if repoErr := s.portalStudentRepo.CreateStudent(tx, req); repoErr != nil {
			_, errMsg := utils.ErrorSpToMessageError(repoErr)
			result.Failed++
			result.FailedRows = append(result.FailedRows, dto.PortalStudentBulkFailedRow{
				Row:   rowNum,
				Name:  req.Name,
				Email: req.Email,
				Error: errMsg,
			})
			s.log.WithError(repoErr).Warn(utils.CreateMsgDebuging(repoErr.Error(), req.Email, utils.ErrorLocation()))
			continue
		}

		result.Success++
	}

	return result, nil
}

func (s *portalStudentService) GetStudentListWithCount(ctx context.Context, pg pageable.PageableRequestClass) (*pageable.PageableResponse[dto.PortalStudentListResponse], error) {
	tx := s.db.WithContext(ctx)

	data, totalData, err := s.portalStudentRepo.GetStudentListWithCount(tx, pg)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := make([]dto.PortalStudentListResponse, 0, len(data))
	for _, item := range data {
		res = append(res, dto.PortalStudentListResponse{
			StudentID:                      item.StudentID,
			StudentStudyProgramID:          item.StudentStudyProgramID,
			Nim:                            item.Nim,
			StudentName:                    item.StudentName,
			StudentStatus:                  item.StudentStatus,
			StudyProgramID:                 item.StudyProgramID,
			StudyProgramCode:               item.StudyProgramCode,
			StudyProgramName:               item.StudyProgramName,
			LecturerPABiodataID:            item.LecturerPABiodataID,
			LecturerPAUserID:               item.LecturerPAUserID,
			LecturerPAName:                 item.LecturerPAName,
			EntryAcademicPeriodID:          item.EntryAcademicPeriodID,
			EntryAcademicPeriodName:        item.EntryAcademicPeriodName,
			CurrentAcademicPeriodID:        item.CurrentAcademicPeriodID,
			CurrentAcademicPeriodName:      item.CurrentAcademicPeriodName,
			CurrentAcademicPeriodShortname: item.CurrentAcademicPeriodShortname,
			EntryPeriodRank:                item.EntryPeriodRank,
			CurrentPeriodRank:              item.CurrentPeriodRank,
			CurrentSemester:                item.CurrentSemester,
			SemesterLabel:                  item.SemesterLabel,
			TotalSksTaken:                  item.TotalSksTaken,
			Ipk:                            item.Ipk,
			CreatedAt:                      item.CreatedAt,
		})
	}

	return &pageable.PageableResponse[dto.PortalStudentListResponse]{
		Data: res,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: utils.TotalPage(totalData, pg.GetDefaultLimit()),
			Page:      pg.GetDefaultPage(),
			Size:      pg.GetDefaultLimit(),
		},
	}, nil
}

package servicemodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/delivery/http/middleware"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstClassScheduleTaskService struct {
	log                            *logrus.Logger
	db                             *gorm.DB
	cache                          cached.CacheRepository
	mstClassScheduleTaskRepository *repositorymodel.MstClassScheduleTaskRepository
}

func NewMstClassScheduleTaskService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstClassScheduleTaskRepository *repositorymodel.MstClassScheduleTaskRepository,
) *MstClassScheduleTaskService {
	return &MstClassScheduleTaskService{
		log:                            log,
		db:                             db,
		cache:                          cache,
		mstClassScheduleTaskRepository: mstClassScheduleTaskRepository,
	}
}

/* Create */
func (s *MstClassScheduleTaskService) Create(ctx context.Context, req dto.MstClassScheduleTaskRequest) (*dto.MstClassScheduleTaskResponse, error) {
	tx := s.db.WithContext(ctx)

	err := s.mstClassScheduleTaskRepository.Create(tx, req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	classScheduleTask, err := s.GetByID(ctx, req.ID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	return classScheduleTask, nil
}

/* Read */
func (s *MstClassScheduleTaskService) GetAll(ctx context.Context, ClassID string) (T []*dto.MstClassScheduleTaskResponse, err error) {
	tx := s.db.WithContext(ctx)

	user := middleware.GetUserClaimsCtx(ctx)

	classScheduleTasks, err := s.mstClassScheduleTaskRepository.GetAll(tx, ClassID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all class participant service",
			"user-id": user.ID,
		}).Error(err)
		err = utils.ErrorSpToErrorFiber(err)
		return nil, err
	}

	results := make([]*dto.MstClassScheduleTaskResponse, 0)

	for _, classScheduleTask := range classScheduleTasks {
		results = append(results, converter.MstClassScheduleTaskToResponse(classScheduleTask))
	}

	return results, nil
}

func (s *MstClassScheduleTaskService) GetByID(ctx context.Context, ID string) (*dto.MstClassScheduleTaskResponse, error) {
	tx := s.db.WithContext(ctx)

	classScheduleTask := new(model.MstClassScheduleTask)

	err := s.mstClassScheduleTaskRepository.GetByID(tx, ID, classScheduleTask)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), ID, utils.ErrorLocation())
		s.log.WithError(err).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	result := converter.MstClassScheduleTaskToResponse(*classScheduleTask)

	return result, nil
}

/* Update */
/* Delete */

package repositorymodel

import (
	"github.com/sirupsen/logrus"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
)

type MstClassScheduleTaskCollectRepository struct {
	repository.Repository[model.MstClassScheduleTaskCollect]
	log *logrus.Logger
}

func NewMstClassScheduleTaskCollectRepository(log *logrus.Logger) *MstClassScheduleTaskCollectRepository {
	return &MstClassScheduleTaskCollectRepository{
		log: log,
	}
}

/* Create */
/* Read */
/* Update */
/* Delete */

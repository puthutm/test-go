package repositorymodel

import (
	"github.com/sirupsen/logrus"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
)

type MstClassScheduleTeachingMaterialRepository struct {
	repository.Repository[model.MstClassScheduleTeachingMaterial]
	log             *logrus.Logger
	cacheRepository cached.CacheRepository
}

func NewMstClassScheduleTeachingMaterialRepository(log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) *MstClassScheduleTeachingMaterialRepository {
	return &MstClassScheduleTeachingMaterialRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

/* Create */
/* Read */
/* Update */
/* Delete */

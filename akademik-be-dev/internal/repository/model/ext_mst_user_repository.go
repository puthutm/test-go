package repositorymodel

import (
	"github.com/sirupsen/logrus"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
)

type MstUserRepository struct {
	repository.Repository[model.MstUser]
	log             *logrus.Logger
	cacheRepository cached.CacheRepository
}

func NewMstUserRepository(log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) *MstUserRepository {
	return &MstUserRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

/* Create */
/* Read */
/* Update */
/* Delete */

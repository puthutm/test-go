package repositorymodel

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
)

type MstBiodataOfUserRepository struct {
	repository.Repository[model.MstBiodataOfUser]
	log             *logrus.Logger
	cacheRepository cached.CacheRepository
}

func NewMstBiodataOfUserRepository(log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) *MstBiodataOfUserRepository {
	return &MstBiodataOfUserRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

/* Create */
/* Read */
func (r *MstBiodataOfUserRepository) GetByUserID(db *gorm.DB, UserID string, data *model.MstBiodataOfUser) error {
	return db.Where("user_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", UserID).First(data).Error
}

/* Update */
/* Delete */

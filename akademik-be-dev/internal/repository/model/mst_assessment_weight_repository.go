package repositorymodel

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
)

type MstAssessmentWeightRepository struct {
	repository.Repository[model.MstAssessmentWeight]
	log             *logrus.Logger
	cacheRepository cached.CacheRepository
}

func NewMstAssessmentWeightRepository(
	log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) *MstAssessmentWeightRepository {
	return &MstAssessmentWeightRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

/* Create */
func (r *MstAssessmentWeightRepository) Create(db *gorm.DB, req *dto.MstAssessmentWeightRequest) error {
	record := map[string]interface{}{
		"id":                           req.ID,
		"attitude_behavior_percentage": req.AttitudeBehaviorPercentage,
		"task_percentage":              req.TaskPercentage,
		"uts_percentage":               req.UTSPercentage,
		"uas_percentage":               req.UASPercentage,
	}
	return db.Table("mst_assessment_weights").Save(record).Error
}

/* Read */
func (r *MstAssessmentWeightRepository) GetFirst(db *gorm.DB) (*model.MstAssessmentWeight, error) {
	var result model.MstAssessmentWeight
	err := db.First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

/* Update */
/* Delete */

package servicemodel

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
)

type MstAssessmentWeightService struct {
	log                           *logrus.Logger
	db                            *gorm.DB
	cache                         cached.CacheRepository
	mstAssessmentWeightRepository *repositorymodel.MstAssessmentWeightRepository
}

func NewMstAssessmentWeightService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstAssessmentWeightRepository *repositorymodel.MstAssessmentWeightRepository,
) *MstAssessmentWeightService {
	return &MstAssessmentWeightService{
		log:                           log,
		db:                            db,
		cache:                         cache,
		mstAssessmentWeightRepository: mstAssessmentWeightRepository,
	}
}

/* Create */
func (s *MstAssessmentWeightService) Create(req *dto.MstAssessmentWeightRequest) error {
	return s.mstAssessmentWeightRepository.Create(s.db, req)
}

/* Read */
func (s *MstAssessmentWeightService) GetFirst() (*dto.MstAssessmentWeightResponse, error) {
	m, err := s.mstAssessmentWeightRepository.GetFirst(s.db)
	if err != nil {
		return nil, err
	}
	return converter.MstAssessmentWeightModelToResponse(m), nil
}

/* Update */
/* Delete */

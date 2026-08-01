package servicemodel

import (
	"github.com/sirupsen/logrus"
	restapisdm "gitlab.unsia.ac.id/icems/icems-tools/gateway/rest-api/sdm"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
)

type MstStudyProgramCurriculumSubjectPrerequisiteService struct {
	log                                                    *logrus.Logger
	db                                                     *gorm.DB
	cache                                                  cached.CacheRepository
	mstStudyProgramCurriculumRepository                    *repositorymodel.MstStudyProgramCurriculumRepository
	mstClassRepository                                     *repositorymodel.MstClassRepository
	mstSubjectRepository                                   *repositorymodel.MstSubjectRepository
	mstStudyProgramCurriculumSubjectPrerequisiteRepository *repositorymodel.MstStudyProgramCurriculumSubjectPrerequisiteRepository

	// external
	generalInformationGatewayRest *restapisdm.GeneralInformationRest
}

func NewMstStudyProgramCurriculumSubjectPrerequisiteService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstStudyProgramCurriculumSubjectPrerequisiteRepository *repositorymodel.MstStudyProgramCurriculumSubjectPrerequisiteRepository,

	// external
	generalInformationGatewayRest *restapisdm.GeneralInformationRest,
) *MstStudyProgramCurriculumSubjectPrerequisiteService {
	return &MstStudyProgramCurriculumSubjectPrerequisiteService{
		log:   log,
		db:    db,
		cache: cache,
		mstStudyProgramCurriculumSubjectPrerequisiteRepository: mstStudyProgramCurriculumSubjectPrerequisiteRepository,

		// external
		generalInformationGatewayRest: generalInformationGatewayRest,
	}
}

// func (s *MstStudyProgramCurriculumSubjectPrerequisiteService) GetForSearchBy(
// 	ctx context.Context,
// 	req dto.GetStudyProgramCurriculumRequest,
// ) ([]dto.MstStudyProgramCurriculumSubjectPrerequisiteResponse, error) {
// 	tx := s.db.WithContext(ctx)
// 	subjectPrerequisiteData, err := s.mstStudyProgramCurriculumSubjectPrerequisiteRepository.GetByStudyProgramCurriculumID(tx, )
// 	if err != nil {
// 		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
// 		s.log.WithError(err).Error(createMsg)
//
// 	}
// 	subjectPrerequisiteDto := make([]dto.MstStudyProgramCurriculumSubjectPrerequisiteResponse, 0, len(subjectPrerequisiteData))
// 	for _, v := range subjectPrerequisiteData {
// 		subjectPrerequisiteDto = append(subjectPrerequisiteDto, converter.MstStudyProgramCurriculumSubjectPrerequisiteModelToResponse(v))
// 	}
// 	return nil, nil
// }

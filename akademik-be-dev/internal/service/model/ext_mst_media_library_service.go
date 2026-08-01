package servicemodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"unsia.ac.id/akademic_be/pkg/storage-library-be/folders"
	medialibrary "unsia.ac.id/akademic_be/pkg/storage-library-be/media_library"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/internal/service/command"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstMediaLibraryService struct {
	log                       *logrus.Logger
	db                        *gorm.DB
	cache                     cached.CacheRepository
	mstMediaLibraryRepository *repositorymodel.MstMediaLibraryRepository
}

func NewMstMediaLibraryService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstMediaLibraryRepository *repositorymodel.MstMediaLibraryRepository,
) *MstMediaLibraryService {
	return &MstMediaLibraryService{
		log:                       log,
		db:                        db,
		cache:                     cache,
		mstMediaLibraryRepository: mstMediaLibraryRepository,
	}
}

// Create
func (s *MstMediaLibraryService) Create(ctx context.Context, req medialibrary.TrxMediaLibrary) error {
	tx := s.db.WithContext(ctx)
	data := new(medialibrary.TrxMediaLibrary)
	err := s.mstMediaLibraryRepository.Create(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create file in storage-folder service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstMediaLibraryService) CreateFolderRoot(ctx context.Context, req folders.MstFolder) error {
	tx := s.db.WithContext(ctx)
	data := new(folders.MstFolder)
	err := s.mstMediaLibraryRepository.CreateFolderRoot(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create folder root in storage-folder service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstMediaLibraryService) CreateFolderByParent(ctx context.Context, req folders.MstFolder) error {
	tx := s.db.WithContext(ctx)
	data := new(folders.MstFolder)
	err := s.mstMediaLibraryRepository.CreateFolderByParent(tx, data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "create folder by parent in storage-folder service",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

func (s *MstMediaLibraryService) GetFileByFolderAndSubject(
	ctx context.Context, req folders.MstFolderRequest_GetFileByFolderAndSubject,
) (T []medialibrary.TrxMediaLibraryResponse, err error) {
	tx := s.db.WithContext(ctx)

	commandReq := command.ConvertDTOtoCommandMediaLibrary_GetFileByFolderAndSubject(req)

	folderss, err := s.mstMediaLibraryRepository.GetFileByFolderAndSubject(tx, commandReq)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service": "get file by folder and subject service",
		}).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	results := make([]medialibrary.TrxMediaLibraryResponse, 0, len(folderss))
	for _, v := range folderss {
		results = append(results, medialibrary.ConvertToTrxMediaLibraryResponse(v))
	}

	return results, nil
}

func (s *MstMediaLibraryService) GetFolderByParent(
	ctx context.Context, parentID string,
) ([]folders.MstFolderResponse, error) {
	tx := s.db.WithContext(ctx)

	folderss, err := s.mstMediaLibraryRepository.GetFolderByParent(tx, parentID)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service": "get folder by parentID service",
		}).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	results := make([]folders.MstFolderResponse, 0, len(folderss))
	for _, v := range folderss {
		results = append(results, folders.ConvertToMstFolderResponse(v))
	}

	return results, nil
}

func (s *MstMediaLibraryService) GetFolderRoot(
	ctx context.Context,
) ([]folders.MstFolderResponse, error) {
	tx := s.db.WithContext(ctx)

	folderss, err := s.mstMediaLibraryRepository.GetFolderRoot(tx)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
		s.log.WithFields(logrus.Fields{
			"service": "get folder root service",
		}).Error(createMsg)

		return nil, utils.ErrorSpToErrorFiber(err)
	}

	results := make([]folders.MstFolderResponse, 0, len(folderss))
	for _, v := range folderss {
		results = append(results, folders.ConvertToMstFolderResponse(v))
	}

	return results, nil
}

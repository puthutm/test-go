package repositorymodel

import (
	"github.com/sirupsen/logrus"
	"unsia.ac.id/akademic_be/pkg/storage-library-be/folders"
	medialibrary "unsia.ac.id/akademic_be/pkg/storage-library-be/media_library"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	"unsia.ac.id/akademic_be/internal/service/command"
)

type MstMediaLibraryRepository struct {
	log             *logrus.Logger
	cacheRepository cached.CacheRepository
}

func NewMstMediaLibraryRepository(log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) *MstMediaLibraryRepository {
	return &MstMediaLibraryRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

// Create
func (r *MstMediaLibraryRepository) Create(db *gorm.DB, req *medialibrary.TrxMediaLibrary) error {
	err := db.Create(req).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "create media library",
			"user_id":    req.CreatedBy,
		}).Error(msg.ErrCreate.Error())
		return err
	}
	return nil
}

func (r *MstMediaLibraryRepository) CreateFolderRoot(db *gorm.DB, req *folders.MstFolder) error {
	err := db.Table("mst_folders").Create(req).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "create media library folder root",
			"user_id":    req.CreatedBy,
		}).Error(msg.ErrCreate.Error())
		return err
	}
	return nil
}

func (r *MstMediaLibraryRepository) CreateFolderByParent(db *gorm.DB, req *folders.MstFolder) error {
	err := db.Table("mst_folders").Create(req).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "create media library folder",
			"user_id":    req.CreatedBy,
		}).Error(msg.ErrCreate.Error())
		return err
	}
	return nil
}

func (r *MstMediaLibraryRepository) GetByID(db *gorm.DB, ID string, data *medialibrary.TrxMediaLibrary) error {
	err := db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get  media library",
			"id":         ID,
		}).Error(msg.ErrRead.Error())
		return err
	}
	return nil
}

func (r *MstMediaLibraryRepository) GetFileByFolderAndSubject(
	db *gorm.DB, req command.MstMediaLibraryRequest_GetFileByFolderAndSubject,
) (T []medialibrary.TrxMediaLibrary, err error) {
	err = db.Table("trx_media_libraries").
		Where("folder_id = ? AND subject_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", req.FolderID, req.SubjectID).
		Find(&T).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get  media library",
			"req":        req,
		}).Error(msg.ErrRead.Error())
		return nil, err
	}
	return T, nil
}

func (r *MstMediaLibraryRepository) GetFolderByParent(
	db *gorm.DB, parentID string,
) (T []folders.MstFolder, err error) {
	err = db.Table("mst_folders").
		Where("parent_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", parentID).
		Find(&T).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get  media library folder by parent id",
			"req":        parentID,
		}).Error(msg.ErrRead.Error())
		return nil, err
	}
	return T, nil
}

func (r *MstMediaLibraryRepository) GetFolderRoot(
	db *gorm.DB,
) (T []folders.MstFolder, err error) {
	err = db.Table("mst_folders").
		Where("(parent_id IS NULL OR parent_id = '') AND (deleted_at IS NULL OR deleted_at = 0)").
		Find(&T).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get  media library folder root",
		}).Error(msg.ErrRead.Error())
		return nil, err
	}
	return T, nil
}

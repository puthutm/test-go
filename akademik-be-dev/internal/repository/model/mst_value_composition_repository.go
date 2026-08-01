package repositorymodel

import (
	"time"

	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
)

type MstValueCompositionRepository interface {
	Create(db *gorm.DB, data *model.MstValueComposition) error
	UpdateByID(db *gorm.DB, data *model.MstValueComposition) error
	DeleteByID(db *gorm.DB, ID, deleteBy string) error
	RestoreByID(db *gorm.DB, ID string) error
	GetByID(db *gorm.DB, ID string, data *model.MstValueComposition) error
	GetAllWithCount(db *gorm.DB, deleted bool, pageble pageable.PageableRequestValueComposition) (T []model.MstValueComposition, count int64, err error)
	GetAcademicPeriodsWithCount(db *gorm.DB, pageble pageable.PageableRequest) (T []model.ValueCompositionGroup, count int64, err error)
	DuplicateByAcademicPeriodID(db *gorm.DB, req dto.MstValueCompositionDuplicateRequest, createdAt int64, createdBy string) error
}

type mstValueCompositionRepository struct {
	log *logrus.Logger
	repository.Repository[model.MstValueComposition]
	cacheRepository cached.CacheRepository
}

func NewMstValueCompositionRepository(
	log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) MstValueCompositionRepository {
	return &mstValueCompositionRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

// Create
func (r *mstValueCompositionRepository) Create(db *gorm.DB, data *model.MstValueComposition) error {
	data.CreatedAt = time.Now().UnixMilli()
	err := db.Create(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "create value composition",
			"id":         data.ID,
		}).Error(msg.ErrCreate.Error())
		return err
	}
	return nil
}

// Update
func (r *mstValueCompositionRepository) UpdateByID(db *gorm.DB, data *model.MstValueComposition) error {
	now := time.Now().UnixMilli()
	data.UpdatedAt = &now

	err := db.Model(&model.MstValueComposition{}).Where("id = ?", data.ID).Updates(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update value composition",
			"id":         data.ID,
		}).Error(msg.ErrUpdate.Error())
		return err
	}
	return nil
}

// Delete
func (r *mstValueCompositionRepository) DeleteByID(db *gorm.DB, ID, deleteBy string) error {
	err := db.Model(&model.MstValueComposition{}).Where("id = ?", ID).Updates(map[string]interface{}{
		"deleted_at": time.Now().UnixMilli(),
		"deleted_by": deleteBy,
	}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "delete value composition",
			"id":         ID,
		}).Error(msg.ErrDelete.Error())
		return err
	}
	return nil
}

// Restore
func (r *mstValueCompositionRepository) RestoreByID(db *gorm.DB, ID string) error {
	err := db.Model(&model.MstValueComposition{}).Where("id = ?", ID).Update("deleted_at", nil).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "restore value composition",
			"id":         ID,
		}).Error(msg.ErrRestore.Error())
		return err
	}
	return nil
}

// GetByID
func (r *mstValueCompositionRepository) GetByID(db *gorm.DB, ID string, data *model.MstValueComposition) error {
	err := db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get value composition",
			"id":         ID,
		}).Error(msg.ErrRestore.Error())
		return err
	}
	return nil
}

// GetAllWithCount
func (r *mstValueCompositionRepository) GetAllWithCount(
	db *gorm.DB,
	deleted bool,
	pageble pageable.PageableRequestValueComposition,
) (T []model.MstValueComposition, count int64, err error) {
	query := db.Model(&model.MstValueComposition{})
	if deleted {
		query = query.Where("deleted_at IS NOT NULL AND deleted_at > 0")
	} else {
		query = query.Where("deleted_at IS NULL OR deleted_at = 0")
	}

	if err := query.Count(&count).Error; err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get all value composition",
		}).Error(msg.ErrMultipleRead.Error())
		return T, count, err
	}

	page := pageble.Page
	limit := pageble.Limit
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	err = query.Offset(offset).Limit(limit).Find(&T).Error
	return T, count, err
}

func (r *mstValueCompositionRepository) GetAcademicPeriodsWithCount(
	db *gorm.DB,
	pageble pageable.PageableRequest,
) (T []model.ValueCompositionGroup, count int64, err error) {
	query := db.Table("mst_value_compositions").
		Select("academic_periode_id, COUNT(id) as total_composition").
		Where("deleted_at IS NULL OR deleted_at = 0").
		Group("academic_periode_id")

	if err := query.Count(&count).Error; err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get academic periods with count",
		}).Error(msg.ErrMultipleRead.Error())
		return T, count, err
	}

	page := pageble.GetDefaultPage()
	limit := pageble.GetDefaultLimit()
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	err = query.Offset(offset).Limit(limit).Scan(&T).Error
	return T, count, err
}

func (r *mstValueCompositionRepository) DuplicateByAcademicPeriodID(db *gorm.DB, req dto.MstValueCompositionDuplicateRequest, createdAt int64, createdBy string) error {
	var sources []model.MstValueComposition
	if err := db.Where("academic_periode_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", req.AcademicPeriodIDSource).Find(&sources).Error; err != nil {
		return err
	}

	if req.IsOverwrite != nil && *req.IsOverwrite {
		db.Where("academic_periode_id = ?", req.AcademicPeriodIDTarget).Delete(&model.MstValueComposition{})
	}

	createdUUID, _ := uuid.Parse(createdBy)
	for _, src := range sources {
		newComp := src
		newComp.ID = uuid.New()
		newComp.AcademicPeriodeID = req.AcademicPeriodIDTarget
		newComp.CreatedAt = createdAt
		newComp.CreatedBy = &createdUUID
		if err := db.Create(&newComp).Error; err != nil {
			r.log.WithFields(logrus.Fields{
				"repository": "duplicate value composition",
				"source":     req.AcademicPeriodIDSource,
				"target":     req.AcademicPeriodIDTarget,
			}).Error(err.Error())
			return err
		}
	}

	return nil
}

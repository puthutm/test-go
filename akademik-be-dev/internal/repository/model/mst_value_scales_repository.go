package repositorymodel

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
)

type MstValueScaleRepository struct {
	log *logrus.Logger
	repository.Repository[model.MstValueScale]
}

func NewMstValueScaleRepository(
	log *logrus.Logger,
) *MstValueScaleRepository {
	return &MstValueScaleRepository{
		log: log,
	}
}

// Create
func (r *MstValueScaleRepository) Create(db *gorm.DB, data *model.MstValueScale) error {
	data.CreatedAt = time.Now().UnixMilli()
	err := db.Create(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "create value scale",
			"id":         data.ID,
		}).Error(msg.ErrCreate.Error())
		return err
	}
	return nil
}

// Update
func (r *MstValueScaleRepository) Update(db *gorm.DB, data *model.MstValueScale) error {
	now := time.Now().UnixMilli()
	data.UpdatedAt = &now
	err := db.Model(&model.MstValueScale{}).Where("id = ?", data.ID).Updates(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update value scale",
			"id":         data.ID,
		}).Error(msg.ErrUpdate.Error())
		return err
	}
	return nil
}

// Delete
func (r *MstValueScaleRepository) DeleteByID(db *gorm.DB, ID, deleteBy string) error {
	err := db.Model(&model.MstValueScale{}).Where("id = ?", ID).Updates(map[string]interface{}{
		"deleted_at": time.Now().UnixMilli(),
		"deleted_by": deleteBy,
	}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "delete value scale",
			"id":         ID,
		}).Error(msg.ErrDelete.Error())
		return err
	}
	return nil
}

// Restore
func (r *MstValueScaleRepository) RestoreByID(db *gorm.DB, ID string) error {
	err := db.Model(&model.MstValueScale{}).Where("id = ?", ID).Update("deleted_at", nil).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "restore value scale",
			"id":         ID,
		}).Error(msg.ErrRestore.Error())
		return err
	}
	return nil
}

// GetByID
func (r *MstValueScaleRepository) GetByID(db *gorm.DB, ID string, data *model.MstValueScale) error {
	err := db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get value scale",
			"id":         ID,
		}).Error(msg.ErrRestore.Error())
		return err
	}
	return nil
}

// GetAllWithCount
func (r *MstValueScaleRepository) GetAllWithCount(
	db *gorm.DB,
	deleted bool,
	pageble pageable.PageableRequestValueScale,
) (T []model.MstValueScale, count int64, err error) {
	query := db.Model(&model.MstValueScale{})
	if deleted {
		query = query.Where("deleted_at IS NOT NULL AND deleted_at > 0")
	} else {
		query = query.Where("deleted_at IS NULL OR deleted_at = 0")
	}

	if err := query.Count(&count).Error; err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get all value scale",
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
